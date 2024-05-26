cat scripts/env/.env-base > .env
cat scripts/env/.env-"$2" >> .env

docker-compose --env-file .env up --build