# MinIO Object Storage

MinIO provides high-performance, S3-compatible object storage for the Sustainable Classrooms platform, handling all file storage needs including videos, images, documents, and user-generated content.

## 🎯 Overview

MinIO serves as the primary object storage solution for:

- **Video Content** - Educational videos and thumbnails
- **User Profiles** - Profile pictures and attachments
- **Text Content** - Documents, images, and file attachments
- **Temporary Files** - Upload processing and staging

## 📦 Bucket Structure

The system uses organized buckets for different content types:

| Bucket Name                           | Purpose                    |
| ------------------------------------- | -------------------------- |
| `video-learning-videos-confirmed`     | Published video content    |
| `video-learning-videos-staging`       | Videos under review        |
| `video-learning-thumbnails-confirmed` | Published video thumbnails |
| `video-learning-thumbnails-staging`   | Thumbnail drafts           |
| `files-profiles`                      | User profile attachments   |
| `files-text`                          | Text content attachments   |

## 🔧 Configuration

### Environment Variables

Configure MinIO through environment files:

#### Development (.env.dev)

```bash
MINIO_ROOT_USER=minioadmin
MINIO_ROOT_PASSWORD=minioadmin
MINIO_SERVER_URL=http://localhost:9000
MINIO_BROWSER_REDIRECT_URL=http://localhost:9090
```

#### Production (.env.prod)

```bash
MINIO_ROOT_USER=your-secure-username
MINIO_ROOT_PASSWORD=your-very-secure-password
MINIO_SERVER_URL=https://your-domain.com
MINIO_BROWSER_REDIRECT_URL=https://console.your-domain.com
```
