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
done
exit 0
