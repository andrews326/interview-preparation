🐳 Kubernetes — Interview Questions & Answers
🌱 Basic Level

1. What is Kubernetes?
Kubernetes (K8s) is an open-source container orchestration platform that automates deployment, scaling, and management of containerized applications.

2. What are Pods in Kubernetes?
A Pod is the smallest deployable unit in Kubernetes, which can contain one or more containers sharing the same network namespace and storage.

3. What is a Node?
A Node is a worker machine (physical or virtual) in Kubernetes that runs Pods. Nodes are managed by the control plane.

4. What is a Cluster?
A Cluster is a set of Nodes managed by a Kubernetes control plane, providing high availability and scalability.

5. What is the role of kube-apiserver?
The API server is the front-end of the control plane. It exposes the Kubernetes API and processes requests from CLI, UI, or other components.

6. What is a Deployment?
A Deployment manages Pods and ReplicaSets, ensuring the desired number of Pods are running and updating them in a controlled way.

7. What is a Service in Kubernetes?
A Service exposes Pods to internal or external traffic and provides a stable IP and DNS name, load balancing requests across Pods.

8. What are Labels and Selectors?
Labels are key-value pairs attached to objects like Pods. Selectors allow filtering and grouping of objects using these labels.

9. What is a ConfigMap?
ConfigMap stores non-sensitive configuration data (like environment variables, config files) to decouple configuration from application code.

10. What is a Secret?
A Secret stores sensitive data (passwords, tokens, keys) in a secure way for use by Pods.

⚙️ Intermediate Level

1. What is the difference between ReplicaSet and Deployment?
ReplicaSet ensures a specified number of Pod replicas are running. Deployment manages ReplicaSets and provides rolling updates, rollbacks, and versioning.

2. What is a DaemonSet?
DaemonSet ensures a copy of a Pod runs on every Node (or selected Nodes) for services like logging, monitoring, or network agents.

3. What is a StatefulSet?
StatefulSet manages stateful applications, providing stable network identities and persistent storage for each Pod.

4. What is a PersistentVolume (PV) and PersistentVolumeClaim (PVC)?

PV: Cluster-level storage resource.

PVC: Pod-level request for storage from a PV.

5. What is a Namespace?
Namespace is a virtual cluster inside a Kubernetes Cluster for isolating resources between teams, projects, or environments.

6. How do you update an application without downtime?
Use Rolling Updates in Deployments to replace Pods gradually while keeping the application available.

7. How do Services work internally?
Services use kube-proxy and iptables/ipvs rules to route traffic to the correct Pods based on labels and selectors.

8. What is the difference between ClusterIP, NodePort, and LoadBalancer Services?

ClusterIP: Internal access only.

NodePort: Exposes service on each Node’s port.

LoadBalancer: Creates an external cloud load balancer for public access.

9. What is Horizontal Pod Autoscaler (HPA)?
HPA automatically scales the number of Pods in a Deployment/ReplicaSet based on CPU/memory or custom metrics.

10. How do you store application logs in Kubernetes?
Logs are written to stdout/stderr by containers and collected by logging agents like Fluentd or EFK/ELK stack.

🚀 Advanced Level

1. How does Kubernetes scheduling work?
The scheduler assigns Pods to Nodes based on resource requests, constraints, affinity/anti-affinity rules, and available resources.

2. How do you perform a Blue-Green deployment in Kubernetes?
Maintain two Deployments (Blue = current, Green = new), switch Services to point to the Green version after testing.

3. What is the difference between Init Containers and regular Containers?
Init Containers run before application containers and perform setup tasks (like configuration, waiting for dependencies).

4. How do you manage secrets securely in Kubernetes?

Use Kubernetes Secrets with encryption at rest.

Use tools like SealedSecrets, Vault, or AWS Secrets Manager.

Restrict RBAC access to Secrets.

5. How does Kubernetes handle high availability?
Deploy control plane components across multiple master Nodes (HA masters), use multiple worker Nodes, and spread Pods across Availability Zones.

6. How do you perform a rolling rollback?
Use kubectl rollout undo deployment <name> to revert to a previous version of a Deployment.

7. What are Taints and Tolerations?
Taints mark Nodes to repel certain Pods. Tolerations allow Pods to schedule on Nodes with matching taints.

8. What is Kubernetes Ingress?
Ingress manages external HTTP/HTTPS traffic routing to Services, often using an Ingress Controller.

9. How do you integrate Kubernetes with CI/CD pipelines?
Use Jenkins, GitHub Actions, or CodePipeline to build Docker images, push to ECR/DockerHub, and deploy to Kubernetes with manifests or Helm charts.

10. What are some Kubernetes best practices for production?

Use Namespaces for isolation

Apply Resource Limits and Requests

Enable HPA and Cluster Autoscaler

Use RBAC for access control

Encrypt Secrets

Monitor with Prometheus/Grafana or CloudWatch

Use Helm or GitOps for declarative deployments