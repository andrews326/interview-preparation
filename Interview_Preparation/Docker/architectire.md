🖥️ Docker Client

Communicates with the Docker Daemon via REST API (usually over a UNIX socket or TCP).

Example commands:

docker build
docker run
docker pull
docker push

⚙️ Docker Daemon

Runs on the host machine.

Responsible for:

Building images

Running containers

Managing networks, volumes, and logs

It listens for API requests from the client.

📦 Docker Images

Read-only templates used to create containers.

Built from a Dockerfile.

Stored locally or pulled from a registry.

🧩 Docker Containers

Running instances of images.

Isolated environments sharing the host’s kernel using:

Namespaces (process isolation)

Cgroups (resource control)

🌐 Docker Registry

Stores and distributes Docker images.

Public: Docker Hub

Private: Self-hosted registry

Used with:

docker push myimage
docker pull myimage

🔄 2. Docker Workflow (Step-by-Step)
Step 1: Build Phase

Developer writes a Dockerfile (a blueprint).

Command:

docker build -t myapp:v1 .


The daemon reads the instructions and creates an image layer by layer.

Step 2: Push/Pull Phase

Once built, the image can be pushed to a registry for sharing:

docker push username/myapp:v1


Other machines can pull it:

docker pull username/myapp:v1

Step 3: Run Phase

Run a container from the image:

docker run -d -p 8080:80 myapp:v1


The container runs as a separate process with its own filesystem, network stack, and resource limits.

Step 4: Networking

Docker sets up a virtual network bridge (docker0) for container communication.

Containers can be connected via:

bridge

host

none

overlay (in Swarm/K8s)

Step 5: Storage

Docker uses storage drivers (like overlay2) to manage image and container layers.

Persistent data is stored using:

Volumes (recommended)

Bind mounts (direct host path)

Step 6: Monitoring & Logs

Logs can be viewed using:

docker logs <container_id>


Stats monitored with:

docker stats

🧠 3. How Docker Isolates Containers

| Feature        | Description                                                          |
| -------------- | -------------------------------------------------------------------- |
| **Namespaces** | Isolates processes, users, networks, and mounts.                     |
| **Cgroups**    | Controls and limits resource usage (CPU, RAM).                       |
| **UnionFS**    | Combines multiple read-only layers into a single unified filesystem. |

🖇️ 4. Docker Networking Architecture

Each container gets its own virtual Ethernet interface connected to Docker’s bridge network (docker0).
It can communicate:

With other containers on the same host

With external networks through NAT

Across hosts using overlay networks (Swarm or Kubernetes)

🧭 5. End-to-End Docker Workflow Summary
Developer → Dockerfile → docker build → Image
Image → Registry → docker pull → Container
Container → Uses network + volume → Runs app

✅ In Short

Docker =
🧠 Lightweight virtualization using Linux features
⚙️ Client–Server model for image and container management
🚀 Fast, portable, and efficient app deployment