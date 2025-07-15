#!/bin/sh
echo "starting buckets..."
mc alias set dockerminio http://minio:9000 ${MINIO_ROOT_USER} ${MINIO_ROOT_PASSWORD}
BUCKETS=(
    "video-learning-videos-confirmed"
    "video-learning-videos-staging"
    "video-learning-thumbnails-confirmed"
    "video-learning-thumbnails-staging"
    "files-profiles"
    "files-text"
)
for bucket in "${BUCKETS[@]}"; do
    mc mb dockerminio/$bucket
    mc anonymous set download dockerminio/$bucket
done

echo "subiendo imágenes al bucket files-text..."
mc cp /init_media/files-text/i1.webp dockerminio/files-text/i1.webp
mc cp /init_media/files-text/i2.webp dockerminio/files-text/i2.webp
mc cp /init_media/files-text/i3.webp dockerminio/files-text/i3.webp
mc cp /init_media/files-text/i4.webp dockerminio/files-text/i4.webp

mc cp /init_media/files-text/i1.webp dockerminio/video-learning-thumbnails-confirmed/i1.webp
mc cp /init_media/files-text/i2.webp dockerminio/video-learning-thumbnails-confirmed/i2.webp
mc cp /init_media/files-text/i3.webp dockerminio/video-learning-thumbnails-confirmed/i3.webp
mc cp /init_media/files-text/i3.webp dockerminio/video-learning-thumbnails-confirmed/i3.webp

mc cp /init_media/video-learning-videos-confirmed/v1.mp4 dockerminio/video-learning-videos-confirmed/v1.mp4
mc cp /init_media/video-learning-videos-confirmed/v2.mp4 dockerminio/video-learning-videos-confirmed/v2.mp4
mc cp /init_media/video-learning-videos-confirmed/v3.mp4 dockerminio/video-learning-videos-confirmed/v3.mp4
mc cp /init_media/video-learning-videos-confirmed/v4.mp4 dockerminio/video-learning-videos-confirmed/v4.mp4

exit 0
