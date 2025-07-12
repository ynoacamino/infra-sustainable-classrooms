#!/bin/sh
echo "starting buckets..."
mc alias set dockerminio http://minio:9000 minioadmin minioadmin
BUCKETS=(
    "video-learning-videos-confirmed"
    "video-learning-videos-staging"
    "video-learning-thumbnails-confirmed"
    "video-learning-thumbnails-staging"
)
for bucket in "${BUCKETS[@]}"; do
    mc mb dockerminio/$bucket
done
exit 0
