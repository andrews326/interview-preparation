26. What is the difference between COPY and ADD in Dockerfile?

COPY: Copies files/directories from the host into the image.

ADD: Similar to COPY, but also supports remote URLs and auto-extracts tar files.
👉 Best practice: Use COPY unless you need ADD features.

27. What is the purpose of EXPOSE in Dockerfile?
It documents which port the container listens on.
Example:

EXPOSE 8080


28. What is the difference between -p and --network?

-p → Maps container ports to host ports.

--network → Connects the container to a specific network.

29. What are Docker networks?
Docker networks enable communication between containers. Types include:

bridge (default)

host

none

overlay (used in Docker Swarm)

30. What is the default network driver in Docker?
bridge — used when no specific network is defined.

31. What is Docker Compose?
A tool for defining and running multi-container applications using a YAML file (docker-compose.yml).

32. What is the command to start containers using Docker Compose?

docker-compose up -d


33. What is the command to stop and remove Docker Compose containers?

docker-compose down


34. What is the purpose of depends_on in Docker Compose?
Specifies dependency order between services — e.g., app depends on database.

35. What is the difference between restart: always and restart: unless-stopped in Docker Compose?

| Option           | Behavior                                              |
| ---------------- | ----------------------------------------------------- |
| `always`         | Restarts container always, even after Docker restarts |
| `unless-stopped` | Restarts until manually stopped                       |


36. What is a Docker volume mount vs. bind mount?

Type	Managed By	Location
Volume	Docker	/var/lib/docker/volumes/...
Bind Mount	User	Any host path

37. How can you inspect a Docker container?

docker inspect <container_id>


38. How do you view container environment variables?

docker inspect --format='{{.Config.Env}}' <container_id>


39. How can you connect two containers?
Create a network:

docker network create mynet


Run containers in that network:

docker run --network=mynet ...


40. What is the use of docker commit?
Creates a new image from a modified container.

docker commit <container_id> myimage:v2


41. What is Docker Swarm?
Docker’s native clustering and orchestration tool — allows managing multiple Docker hosts as a single system.

42. How do you initialize a Docker Swarm?

docker swarm init


43. What are the main Swarm components?

Manager Nodes – control the cluster

Worker Nodes – execute tasks

Services – define how containers run in the swarm

44. What is a Docker service?
A higher-level abstraction for running containers in a swarm (e.g., 3 replicas of a web app).

45. How do you scale a Docker service?

docker service scale myapp=5


46. What is the purpose of Docker secrets?
Used to securely store and manage sensitive data like passwords or API keys in Swarm mode.

47. What is Docker health check?
A command in Dockerfile to test if a container is still working properly.
Example:

HEALTHCHECK CMD curl -f http://localhost:8080 || exit 1


48. How can you clean up unused Docker resources?

docker system prune -a


49. How do you tag and push an image to Docker Hub?

docker tag myapp username/myapp:v1
docker push username/myapp:v1


50. What is multi-stage build in Docker?
Used to optimize image size by using multiple FROM instructions.
Example:

FROM golang:1.20 AS builder
WORKDIR /app
RUN go build -o myapp

FROM alpine:latest
COPY --from=builder /app/myapp .
CMD ["./myapp"]


c group 
namespace