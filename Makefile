compose = docker-compose

build:
	${compose} build

start:
	${compose} up

stop:
	${compose} down

reload:
	${compose} down
	${compose} up --build

purge:
	docker system prune --volumes --all -f