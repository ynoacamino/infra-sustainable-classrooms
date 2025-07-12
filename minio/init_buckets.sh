#!/bin/sh
echo "starting buckets..."
mc alias set dockerminio http://minio:9000 minioadmin minioadmin
BUCKETS=(
    "video_learning_videos_confirmed"
    "video_learning_videos_staging"
    "video_learning_thumbnails_confirmed"
    "video_learning_thumbnails_staging"
)
for bucket in "${BUCKETS[@]}"; do
    mc mb dockerminio/$bucket
done
exit 0
