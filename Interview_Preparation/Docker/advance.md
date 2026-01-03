🚀 Docker Interview Questions — Advanced Level

51. What is Docker overlay network?
An overlay network allows containers on different hosts to communicate securely as if they were on the same host — mainly used in Docker Swarm or Kubernetes.

52. What is Docker namespace?
Namespaces provide isolation for processes, networks, users, etc., ensuring that containers don’t interfere with each other or the host.

53. What are control groups (cgroups) in Docker?
Cgroups limit and monitor resource usage (CPU, memory, I/O) for containers.

54. How does Docker achieve process isolation?
Through Linux namespaces (for isolation) and cgroups (for resource control).

55. What are Docker storage drivers?
They manage how image layers and container writable layers are stored.
Common drivers: overlay2, aufs, btrfs, zfs.

56. What is the difference between image layers and container layers?

Image layers: Read-only layers created by each Dockerfile instruction.

Container layer: The top writable layer for temporary data and runtime changes.

57. How does Docker use the Union File System?
Docker uses UnionFS to merge multiple layers into one unified view.
Each image layer builds upon the previous one.

58. How do you inspect Docker image layers?

docker history <image_name>


59. What is the difference between Docker restart policies?

| Policy           | Description                                   |
| ---------------- | --------------------------------------------- |
| `no`             | Never restart                                 |
| `always`         | Always restart                                |
| `on-failure`     | Restart only if container exits with non-zero |
| `unless-stopped` | Restart unless manually stopped               |


60. What are Docker plugins?
Extensions that add new functionality (e.g., storage, networking, logging) without modifying Docker itself.

61. How do you configure Docker to use a proxy?
Create or edit /etc/systemd/system/docker.service.d/http-proxy.conf and set:

[Service]
Environment="HTTP_PROXY=http://proxy:8080"


Then reload and restart Docker.

62. What is the difference between docker run and docker start?

docker run → creates and starts a new container.

docker start → starts an existing, stopped container.

63. How can you limit container CPU and memory usage?

docker run -m 512m --cpus="1.5" myapp


64. What is Docker security scanning?
It scans images for known vulnerabilities (CVEs) in OS packages and libraries.

65. What are some Docker security best practices?

Use minimal base images (e.g., Alpine).

Avoid running as root (USER directive).

Use signed images.

Regularly scan for vulnerabilities.

Use secrets instead of environment variables.

66. How do you handle persistent storage for databases in Docker?
Use volumes or bind mounts to store data outside the container lifecycle.

67. What is Docker checkpoint and restore?
Feature (via CRIU) that allows pausing a running container and resuming it later — useful for migration or fault recovery.

68. What is Docker Content Trust (DCT)?
A security feature that signs and verifies image integrity using digital signatures.
Enable it via:

export DOCKER_CONTENT_TRUST=1


69. What is BuildKit in Docker?
An improved build engine that supports parallel builds, caching, and secrets.
Enable it via:

DOCKER_BUILDKIT=1 docker build .


70. What is the purpose of Docker context?
It allows managing multiple Docker environments (local, remote, cloud).

docker context ls
docker context use my-remote


71. How do you connect Docker with Kubernetes?
Kubernetes can use Docker (or containerd) as its container runtime to deploy and manage containers through Pods.

72. What is the difference between Docker and containerd?
Docker is a full platform (build, run, ship), while containerd is the lower-level runtime that handles container execution.

73. What is the difference between Docker Compose and Kubernetes?

| Feature    | Docker Compose                | Kubernetes                  |
| ---------- | ----------------------------- | --------------------------- |
| Purpose    | Local container orchestration | Cluster-level orchestration |
| Scaling    | Manual                        | Automatic                   |
| Networking | Simple bridge                 | Advanced (CNI)              |
| Use Case   | Dev/test                      | Production                  |


74. How can you monitor Docker containers?
Tools like:

Docker stats (built-in)

cAdvisor

Prometheus + Grafana

ELK Stack

75. How do you debug container issues?

View logs: docker logs <id>

Inspect: docker inspect <id>

Enter shell: docker exec -it <id> bash

Check events: docker events