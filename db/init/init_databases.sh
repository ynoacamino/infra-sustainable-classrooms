echo "Running SQL scripts"

echo "🛠️ Create databases"
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "infrastructure_db" -c "CREATE DATABASE auth_db;"
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "infrastructure_db" -c "CREATE DATABASE profiles_db;"
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "infrastructure_db" -c "CREATE DATABASE text_db;"
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "infrastructure_db" -c "CREATE DATABASE knowledge_db;"
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "infrastructure_db" -c "CREATE DATABASE video_learning_db;"
# psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" -c "CREATE DATABASE users_db;"
# psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" -c "CREATE DATABASE payments_db;"
# psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" -c "CREATE DATABASE notifications_db;"
# psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" -c "CREATE DATABASE analytics_db;"

echo "🛠️ Init databases"
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname=auth_db < /docker-entrypoint-initdb.d/auth_db.sql
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname=profiles_db < /docker-entrypoint-initdb.d/profiles_db.sql
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname=text_db < /docker-entrypoint-initdb.d/text_db.sql
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname=knowledge_db < /docker-entrypoint-initdb.d/knowledge_db.sql
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname=video_learning_db < /docker-entrypoint-initdb.d/video_learning_db.sql
# psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname=users_db < /docker-entrypoint-initdb.d/users_db.sql
# psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname=payments_db < /docker-entrypoint-initdb.d/payments_db.sql
# psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname=notifications_db < /docker-entrypoint-initdb.d/notifications_db.sql
# psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname=analytics_db < /docker-entrypoint-initdb.d/analytics_db.sql
