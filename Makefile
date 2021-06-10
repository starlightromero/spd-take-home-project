start:
	docker-compose up

stop:
	docker-compose down

reload:
	docker-compose down
	docker rmi spd-take-home-project_weather
	docker-compose up

rmi:
	docker rmi spd-take-home-project_weather