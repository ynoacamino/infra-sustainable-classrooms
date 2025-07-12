BUCKETS=(
    "video_learning_videos_confirmed"
    "video_learning_videos_staging"
    "video_learning_thumbnails_confirmed"
    "video_learning_thumbnails_staging"
)

MINIO_ROOT_USER=minioadmin
MINIO_ROOT_PASSWORD=minioadmin

echo "WARNING: you are meant to run this from the root of the project, otherwise the generated command will end somewhere random or not work at all."

cat <<EOF > ./minio/init_buckets.sh
#!/bin/sh
echo "starting buckets..."
mc alias set dockerminio http://minio:9000 ${MINIO_ROOT_USER} ${MINIO_ROOT_PASSWORD}
BUCKETS=(
EOF
for bucket in "${BUCKETS[@]}"; do
    echo "    \"$bucket\"" >> ./minio/init_buckets.sh
done
cat <<EOF >> ./minio/init_buckets.sh
)
for bucket in "\${BUCKETS[@]}"; do
    mc mb dockerminio/\$bucket
done
exit 0
EOF

chmod +x ./minio/init_buckets.sh

echo
echo "init_buckets.sh generated."