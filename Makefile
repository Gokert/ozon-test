up:
	chmod +x ./scripts/bash/selectDB.sh
	./scripts/bash/selectDB.sh run postgresql
	docker-compose --env-file .env up --build

redis:
	chmod +x ./scripts/bash/selectDB.sh
	./scripts/bash/selectDB.sh run redis
	docker-compose --env-file .env up --build

postgresql:
	chmod +x ./scripts/bash/selectDB.sh
	./scripts/bash/selectDB.sh run postgresql
	docker-compose --env-file .env up --build