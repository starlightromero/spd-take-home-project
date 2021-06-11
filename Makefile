start:
	docker-compose up

stop:
	docker-compose down

reload:
	docker-compose down
	docker-compose up --build

purge:
	docker system prune --volumes --all -f