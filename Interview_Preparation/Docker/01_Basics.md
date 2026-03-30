🧱 Docker Interview Questions — Basic Level

1. What is Docker?
Docker is an open-source platform that automates the deployment of applications inside lightweight, portable containers.

2. What is a container?
A container is a runtime environment that packages an application and its dependencies together so it runs consistently across environments.

3. What is the difference between a container and a virtual machine (VM)?
 
Feature	     Container	              Virtual Machine
OS	           Shares host OS kernel	Has its own OS
Size	       Lightweight (MBs)	    Heavy (GBs)
Startup	       Seconds	                Minutes
Performance    Near-native	            Slight overhead

4. What is a Docker image?
A Docker image is a read-only template with instructions to create a container. It includes the application, dependencies, and configuration.

5. What is a Docker container?
A running instance of a Docker image. Containers are isolated environments that share the host’s kernel.

6. What is Docker Engine?
The core part of Docker — it’s the runtime that builds, runs, and manages containers.

7. What are the main components of Docker?

Docker Client

Docker Daemon

Docker Images

Docker Containers

Docker Registry

8. What is the Docker daemon?
The background service that manages Docker objects like images, containers, networks, and volumes. It listens for Docker API requests.

9. What is the Docker client?
The command-line tool (docker) that interacts with the Docker daemon using REST API.

10. What is Docker Hub?
A public registry where you can store and share Docker images (like GitHub for containers).

11. What is the difference between Docker Hub and Docker Registry?

Docker Hub is a public registry.

Docker Registry can be public or private (self-hosted).

12. What is a Dockerfile?
A text file containing instructions to build a Docker image step by step.

13. What are some common Dockerfile instructions?

FROM – base image

RUN – execute commands

COPY / ADD – copy files

WORKDIR – set working directory

CMD / ENTRYPOINT – define container startup command

EXPOSE – open port

14. What is the difference between CMD and ENTRYPOINT?

🧱 Docker Interview Questions — Basic Level

1. What is Docker?
Docker is an open-source platform that automates the deployment of applications inside lightweight, portable containers.

2. What is a container?
A container is a runtime environment that packages an application and its dependencies together so it runs consistently across environments.

3. What is the difference between a container and a virtual machine (VM)?

| Feature     | Container             | Virtual Machine |
| ----------- | --------------------- | --------------- |
| OS          | Shares host OS kernel | Has its own OS  |
| Size        | Lightweight (MBs)     | Heavy (GBs)     |
| Startup     | Seconds               | Minutes         |
| Performance | Near-native           | Slight overhead |


4. What is a Docker image?
A Docker image is a read-only template with instructions to create a container. It includes the application, dependencies, and configuration.

5. What is a Docker container?
A running instance of a Docker image. Containers are isolated environments that share the host’s kernel.

6. What is Docker Engine?
The core part of Docker — it’s the runtime that builds, runs, and manages containers.

7. What are the main components of Docker?

Docker Client

Docker Daemon

Docker Images

Docker Containers

Docker Registry

8. What is the Docker daemon?
The background service that manages Docker objects like images, containers, networks, and volumes. It listens for Docker API requests.

9. What is the Docker client?
The command-line tool (docker) that interacts with the Docker daemon using REST API.

10. What is Docker Hub?
A public registry where you can store and share Docker images (like GitHub for containers).

11. What is the difference between Docker Hub and Docker Registry?

Docker Hub is a public registry.

Docker Registry can be public or private (self-hosted).

12. What is a Dockerfile?
A text file containing instructions to build a Docker image step by step.

13. What are some common Dockerfile instructions?

FROM – base image

RUN – execute commands

COPY / ADD – copy files

WORKDIR – set working directory

CMD / ENTRYPOINT – define container startup command

EXPOSE – open port

14. What is the difference between CMD and ENTRYPOINT?

| Command      | Purpose                               |
| ------------ | ------------------------------------- |
| `CMD`        | Default command; can be overridden    |
| `ENTRYPOINT` | Main executable; can accept arguments |


15. What is the purpose of the .dockerignore file?
It tells Docker which files to ignore when building an image (similar to .gitignore).

16. How do you build a Docker image?

docker build -t myapp:latest .


17. How do you run a Docker container?

docker run -d -p 8080:80 myapp:latest


18. How do you list running containers?

docker ps


19. How do you stop a running container?

docker stop <container_id>


20. How do you remove a container or image?

docker rm <container_id>
docker rmi <image_id>


21. How do you check logs of a container?

docker logs <container_id>


22. How do you execute a command inside a running container?

docker exec -it <container_id> /bin/bash


23. How do you check container resource usage?

docker stats


24. What are Docker volumes?
Volumes are used to persist data outside containers, allowing data to survive even after the container is deleted.

25. How do you create and use a volume?

docker volume create mydata
docker run -v mydata:/data myapp

15. What is the purpose of the .dockerignore file?
It tells Docker which files to ignore when building an image (similar to .gitignore).

16. How do you build a Docker image?

docker build -t myapp:latest .


17. How do you run a Docker container?

docker run -d -p 8080:80 myapp:latest


18. How do you list running containers?

docker ps


19. How do you stop a running container?

docker stop <container_id>


20. How do you remove a container or image?

docker rm <container_id>
docker rmi <image_id>


21. How do you check logs of a container?

docker logs <container_id>


22. How do you execute a command inside a running container?

docker exec -it <container_id> /bin/bash


23. How do you check container resource usage?

docker stats


24. What are Docker volumes?
Volumes are used to persist data outside containers, allowing data to survive even after the container is deleted.

25. How do you create and use a volume?

docker volume create mydata
docker run -v mydata:/data myapp

to get port number of image

docker run -d   --name nginx-proxy   -p 9000:9000   -v $(pwd)/nginx.conf:/etc/nginx/nginx.conf   nginx 
