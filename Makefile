ubuntu-19.04:
	docker build -f docker/ubuntu_19_04.Dockerfile -t vorta/ubuntu19_04 .
	docker cp $(docker ps -alq):/home/user/vorta/deploy/linux deploy/
