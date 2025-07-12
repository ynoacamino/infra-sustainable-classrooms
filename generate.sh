rm -rf ./sercices/auth/gen
rm -rf ./sercices/profiles/gen
rm -rf ./sercices/text/gen
rm -rf ./sercices/knowledge/gen
rm -rf ./sercices/video_learning/gen

echo "Generating code for auth service..."
goa gen github.com/ynoacamino/infra-sustainable-classrooms/services/auth/design/api -o ./services/auth/

echo "Generating code for profiles service..."
goa gen github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/design/api -o ./services/profiles/

echo "Generating code for text service..."
goa gen github.com/ynoacamino/infra-sustainable-classrooms/services/text/design/api -o ./services/text/

echo "Generating code for knowledge service..."
goa gen github.com/ynoacamino/infra-sustainable-classrooms/services/knowledge/design/api -o ./services/knowledge/

echo "Generating code for video_learning service..."
goa gen github.com/ynoacamino/infra-sustainable-classrooms/services/video_learning/design/api -o ./services/video_learning/

echo "Generating SQL code..."
sqlc generate
