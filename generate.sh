rm -rf ./sercices/auth/gen

echo "Generating code for auth service..."
goa gen github.com/ynoacamino/infra-sustainable-classrooms/services/auth/design/api -o ./services/auth/

echo "Generating code for profiles service..."
goa gen github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/design/api -o ./services/profiles/

echo "Generating code for text service..."
goa gen github.com/ynoacamino/infra-sustainable-classrooms/services/text/design/api -o ./services/text/

echo "Generating code for stats service..."
goa gen github.com/ynoacamino/infra-sustainable-classrooms/services/stats/design/api -o ./services/stats/

echo "Generating SQL code..."
sqlc generate
