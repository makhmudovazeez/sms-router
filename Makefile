init:
	docker-compose -f .\dockerdocker-compose down -v
	docker-compose -f .\dockerdocker-compose up -d --build

stop:
	docker-compose -f .\docker\docker-compose.yml stop

start:
	docker-compose -f .\docker\docker-compose.yml up -d