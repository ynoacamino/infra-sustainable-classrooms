# Mailserver Deployment Guide

This guide provides step-by-step instructions for deploying a production-ready mailserver using Docker Mailserver for the Sustainable Classrooms application.

This documentation is primarily sourced from [the official documentation](https://docker-mailserver.github.io/docker-mailserver/latest), but tailored to the application.

## Prerequisites

- A domain name (e.g., `example.com`)
- Firewall configured to allow mail ports (25, 143, 465, 587, 993)
- Firewall configured to allow other ports for additional configuration (80, 443)
- At least 2GB RAM and 10GB disk space

## DNS Configuration

### Required DNS Records

For a domain `example.com` with mail server hostname `mail.example.com` and server IP `11.22.33.44`:

#### 1. MX Record

```
example.com.    IN  MX  10  mail.example.com.
```

#### 2. A Record

```
mail.example.com.   IN  A   11.22.33.44
```

#### 3. PTR Record (Reverse DNS)

Configure with your hosting provider:

```
44.33.22.11.in-addr.arpa.   IN  PTR mail.example.com.
```

#### 4. SPF Record (TXT)

```
example.com.    IN  TXT "v=spf1 mx ~all"
```

#### 5. DKIM Record (TXT)

After deployment, generate DKIM key and add:

```
mail._domainkey.example.com.    IN  TXT "v=DKIM1; k=rsa; p=YOUR_DKIM_PUBLIC_KEY"
```

#### 6. DMARC Record (TXT)

```
_dmarc.example.com. IN  TXT "v=DMARC1; p=quarantine; rua=mailto:dmarc@example.com"
```

## SSL/TLS Certificate Setup

### Let's Encrypt (Recommended)

Let's Encrypt provides free SSL certificates. This is the recommended approach for production deployments.

#### Step 1: Get Certificate Using Certbot

Create directory structure:

```bash
mkdir -p docker-data/certbot/{certs,logs}
```

Get certificate (requires port 80 access from internet):

```bash
docker run --rm -it \
  -v "${PWD}/docker-data/certbot/certs/:/etc/letsencrypt/" \
  -v "${PWD}/docker-data/certbot/logs/:/var/log/letsencrypt/" \
  -p 80:80 \
  certbot/certbot certonly --standalone -d mail.example.com
```

#### Step 2: Configure Automatic Renewal

Create a systemd timer for automatic renewal:

Create `/etc/systemd/system/certbot-renewal.service`:

```ini
[Unit]
Description=Certbot Renewal
After=network-online.target
Wants=network-online.target

[Service]
Type=oneshot
ExecStart=/usr/bin/docker run --rm \
  -v /path/to/your/project/docker-data/certbot/certs/:/etc/letsencrypt/ \
  -v /path/to/your/project/docker-data/certbot/logs/:/var/log/letsencrypt/ \
  -p 80:80 \
  -p 443:443 \
  certbot/certbot renew
```

Create `/etc/systemd/system/certbot-renewal.timer`:

```ini
[Unit]
Description=Certbot Renewal Timer

[Timer]
OnCalendar=daily
RandomizedDelaySec=3600
Persistent=true

[Install]
WantedBy=timers.target
```

Enable and start the timer:

```bash
sudo systemctl enable certbot-renewal.timer
sudo systemctl start certbot-renewal.timer
```

## Mailserver Configuration

### Step 1: Update compose.yaml

Edit the `compose.yaml` file and update the hostname:

```yaml
services:
  mailserver:
    image: ghcr.io/docker-mailserver/docker-mailserver:latest
    container_name: mailserver
    hostname: mail.example.com # Replace with your actual domain
    env_file: mailserver.env
    ports:
      - "25:25" # SMTP
      - "143:143" # IMAP4
      - "465:465" # ESMTP (implicit TLS)
      - "587:587" # ESMTP (explicit TLS)
      - "993:993" # IMAP4 (implicit TLS)
    volumes:
      - ./docker-data/dms/mail-data/:/var/mail/
      - ./docker-data/dms/mail-state/:/var/mail-state/
      - ./docker-data/dms/mail-logs/:/var/log/mail/
      - ./docker-data/dms/config/:/tmp/docker-mailserver/
      - ./docker-data/certbot/certs/:/etc/letsencrypt/ # This one includes letsencrypt
      - /etc/localtime:/etc/localtime:ro
    restart: always
    stop_grace_period: 1m
    healthcheck:
      test: "ss --listening --ipv4 --tcp | grep --silent ':smtp' || exit 1"
      timeout: 3s
      retries: 5
```

### Step 2: Configure Environment Variables

Edit the `mailserver.env` file with your specific settings:

#### Essential Configuration

```bash
# Set your postmaster email
POSTMASTER_ADDRESS=postmaster@example.com

# Enable Let's Encrypt SSL
SSL_TYPE=letsencrypt

# Enable essential security features
ENABLE_OPENDKIM=1
ENABLE_OPENDMARC=1
ENABLE_POLICYD_SPF=1

# Enable spam protection
ENABLE_RSPAMD=1
ENABLE_CLAMAV=1

# Enable quotas
ENABLE_QUOTAS=1

# Set log level
LOG_LEVEL=info

# Disable Docker network access for security
PERMIT_DOCKER=none
```

#### Security Settings

```bash
# Enable spoof protection
SPOOF_PROTECTION=1

# Set TLS level to modern for better security
TLS_LEVEL=modern

# Enable Fail2Ban (requires NET_ADMIN capability)
ENABLE_FAIL2BAN=1
FAIL2BAN_BLOCKTYPE=reject

# Set message size limits (10MB default)
POSTFIX_MESSAGE_SIZE_LIMIT=10485760

# Enable virus mail deletion after 7 days
VIRUSMAILS_DELETE_DELAY=7
```

#### Optional Features

```bash
# Enable reports
PFLOGSUMM_TRIGGER=daily_cron
PFLOGSUMM_RECIPIENT=admin@example.com

# Log rotation
LOGROTATE_INTERVAL=weekly
LOGROTATE_COUNT=4

# Enable MTA-STS for better security
ENABLE_MTA_STS=1
```

## Deployment

### Step 1: Create Directory Structure

```bash
mkdir -p docker-data/{dms/{mail-data,mail-state,mail-logs,config},certbot/{certs,logs}}
```

### Step 2: Set Permissions

```bash
sudo chown -R 5000:5000 docker-data/dms/mail-data
sudo chown -R 5000:5000 docker-data/dms/mail-state
```

### Step 3: Start the Mailserver (Binding port 25 requires admin privileges or sudo)

```bash
docker compose up -d
```

### Step 4: Check Container Status

```bash
docker compose ps
docker compose logs -f mailserver
```

### Step 5: Quickly refer to [User Management](##User-Management) to create a initial user

Failure to do this under 2 minutes will cause the container to restart (for safety reasons), if the time expires just re-compose the container and try again.

```bash
docker compose down
docker compose up
```

## User Management

### Adding Email Accounts

Use the Docker Mailserver management script:

```bash
# Add a new user
docker exec -ti mailserver setup email add user@example.com password123

# List all users
docker exec -ti mailserver setup email list

# Update password
docker exec -ti mailserver setup email update user@example.com new_password

# Delete user
docker exec -ti mailserver setup email del user@example.com
```

### Adding Aliases

```bash
# Add alias
docker exec -ti mailserver setup alias add alias@example.com user@example.com

# List aliases
docker exec -ti mailserver setup alias list

# Delete alias
docker exec -ti mailserver setup alias del alias@example.com
```

### Configuring DKIM

```bash
# Generate DKIM key
docker exec -ti mailserver setup config dkim

# Get DKIM public key for DNS (Change example.com to the domain name)
docker exec -ti mailserver cat /tmp/docker-mailserver/opendkim/keys/example.com/mail.txt
```

Add the displayed DKIM record to your DNS.

## Monitoring and Maintenance

### Backup Strategy

Create regular backups of:

- Mail data: `docker-data/dms/mail-data/`
- Configuration: `docker-data/dms/config/`
- SSL certificates: `docker-data/certbot/certs/`

```bash
# Example backup script
tar -czf mailserver-backup-$(date +%Y%m%d).tar.gz docker-data/
```

### Updates

```bash
# Update mailserver image
docker compose pull
docker compose up -d
```

## Integration with Application

To integrate with your Sustainable Classrooms application, configure your application to use SMTP with these settings:

```
SMTP Host: mail.example.com
SMTP Port: 587 (STARTTLS) or 465 (SSL/TLS)
Username: app@example.com
Password: [your_password]
Security: STARTTLS or SSL/TLS
```

Create a dedicated email account for the application:

```bash
docker exec -ti mailserver setup email add app@example.com secure_app_password
```
