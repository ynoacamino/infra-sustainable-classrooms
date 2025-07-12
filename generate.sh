rm -rf ./sercices/auth/gen

echo "Generating code for auth service..."
goa gen github.com/ynoacamino/infra-sustainable-classrooms/services/auth/design/api -o ./services/auth/

echo "Generating code for profiles service..."
goa gen github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/design/api -o ./services/profiles/

echo "Generating code for knowledge service..."
goa gen github.com/ynoacamino/infra-sustainable-classrooms/services/knowledge/design/api -o ./services/knowledge/

echo "Generating SQL code..."
sqlc generate
