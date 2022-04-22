# Docker commands
## Image commands
Show images
```
docker images
```
Remove images
```
docker rmi <img>
```
Run image
```
docker start <img>
```
## Containers commands
Active containers
```
docker ps
```
Show all containers
```
docker ps -a
```
Stop / start container
```
docker stop / start <container id>
```
Remove container
```
docker rm <container>
```
## Run docker-compose
```
docker-compose up
```
## Containerize app
Create image
```
docker build . --> EXECUTE Dockerfile
```
Run image
```
docker run <image>
```