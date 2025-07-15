# Mailing Infrastructure

A complete email infrastructure setup for the Sustainable Classrooms platform, providing secure email delivery, SMTP services, and email management capabilities using Docker Mailserver.

## 🎯 Overview

The mailing infrastructure provides:

- **SMTP Server** - Secure email delivery for application notifications
- **IMAP/POP3** - Email access for administrative accounts
- **Anti-Spam Protection** - SpamAssassin and ClamAV integration
- **Security Features** - DKIM, SPF, DMARC, and TLS encryption
- **Domain Management** - Multi-domain email support
- **Monitoring** - Comprehensive logging and metrics

## 📋 Components

### Docker Mailserver

Production-ready mailserver with:
- **Postfix** - SMTP server for email delivery
- **Dovecot** - IMAP/POP3 server for email access
- **SpamAssassin** - Anti-spam filtering
- **ClamAV** - Antivirus scanning
- **OpenDKIM** - DKIM signing
- **Fail2ban** - Intrusion prevention

## 🚀 Quick Start

### Basic Setup

Please refer to the [minimal configuration](./minimal-mailing-setup.md) for a simple setup that should be easy to get set up and running.

### Full Setup

Please refer to the [full configuration](./mailing-setup.md) when you are sure that the minimal setup works and want to fully configure the mailing server to work without risks.