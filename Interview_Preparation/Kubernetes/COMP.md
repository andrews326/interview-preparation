⚙️ Kubernetes Components (Complete Explanation)
🧠 1. Control Plane Components

These are the “brains” of the Kubernetes cluster — they manage the cluster’s overall state and make scheduling, scaling, and deployment decisions.

a. kube-apiserver

Acts as the front door to the Kubernetes cluster.

All commands from CLI (kubectl), UI (Dashboard), or other components go through the API Server.

It validates and processes REST requests (like creating Pods, Deployments, Services, etc.).

Communicates with etcd to store and retrieve cluster data.

👉 Interview Tip:
If the API server is down, you can’t make any changes to the cluster, but running workloads continue.

b. etcd

A distributed key-value database used to store cluster state and configuration.

It contains information like:

What Pods exist

Which Nodes they’re running on

Configurations, Secrets, and Service details

Highly consistent and fault-tolerant.

👉 Example:
If you run kubectl get pods, that information ultimately comes from etcd.

c. kube-scheduler

Decides which Node a Pod should run on.

It considers resource requirements (CPU, memory), taints/tolerations, affinity rules, and node availability.

If no suitable Node is found, the Pod remains pending.

👉 Example:
When you create a new Deployment with 5 replicas, the scheduler finds appropriate Nodes for each Pod.

d. kube-controller-manager

Runs multiple controllers — background processes that continuously check the cluster’s state and try to move it toward the desired state.

Key controllers include:

Node Controller: Monitors Node health.

Replication Controller: Ensures the desired number of Pods are running.

Endpoint Controller: Updates Service–Pod relationships.

Service Account & Token Controller: Manages default accounts and tokens for new Pods.

👉 Analogy:
Think of controllers as the “autopilot” that corrects the system whenever something goes off balance.

e. cloud-controller-manager

Manages interactions with cloud providers (like AWS, GCP, Azure).

Handles tasks such as:

Creating Load Balancers

Managing Volumes (EBS, etc.)

Assigning Node IP addresses

👉 Example:
If you deploy a Service of type LoadBalancer on AWS, this component requests and attaches an AWS ELB automatically.

🖥️ 2. Node Components

These run on every worker node and handle Pod execution and communication with the control plane.

a. kubelet

The primary agent running on each Node.

Ensures the containers described in PodSpecs are running and healthy.

Communicates with the API Server to receive Pod definitions and reports status back.

👉 Example:
If a container crashes, kubelet will restart it automatically.

b. kube-proxy

Handles network communication within and outside the cluster.

Maintains network rules (iptables/ipvs) to route traffic to the right Pod based on Services.

Supports load balancing between Pods.

👉 Example:
When you hit a Service IP, kube-proxy ensures your request reaches one of the backend Pods.

c. Container Runtime

Responsible for running the containers themselves.

Common runtimes: containerd, CRI-O, and formerly Docker.

kubelet uses this runtime to start, stop, and manage containers.

👉 Interview Tip:
Docker was deprecated as a runtime in Kubernetes 1.24 — replaced by containerd or CRI-O.

📦 3. Add-on Components (Optional but Important)

These aren’t mandatory, but are usually present in production clusters for observability and networking.

CoreDNS: Provides DNS-based service discovery (resolves Pod/Service names).

Metrics Server: Collects resource usage data (used by Horizontal Pod Autoscaler).

Ingress Controller: Manages external access via HTTP/HTTPS routes.

Dashboard: Web UI for Kubernetes management.

Network Plugins (CNI): Implement Pod networking (e.g., Calico, Flannel, Weave).

4. How They All Work Together

Here’s a simple flow:

You run kubectl apply -f deployment.yaml → sent to API Server

etcd stores the new desired state (e.g., 3 Pods running)

Scheduler assigns Pods to Nodes

Controller Manager ensures Pods match desired count

kubelet runs the Pods on assigned Nodes

kube-proxy routes network traffic correctly