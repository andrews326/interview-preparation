

| Service Type   | What Port You Use in URL    | Publicly Accessible?  |Notes                                     |
| -------------- | --------------------------- | --------------------- | ----------------------------------------- |
| `NodePort`     | `http://<NodeIP>:nodePort`  | ✅ Yes (with firewall) | You must use the `nodePort` (e.g., 32625) |
| `LoadBalancer` | `http://<EXTERNAL-IP>:port` | ✅ Yes                 | You use the `port`, not the nodePort      |
| `ClusterIP`    | ❌ Not accessible externally | ❌ No                  | For internal cluster use only             |



Why is the service assigned a ClusterIP by default?

In Kubernetes, when you define a Service without explicitly specifying the type, it defaults to ClusterIP. This is the default behavior of Kubernetes.

ClusterIP means that the service is accessible only within the cluster using an internal IP address.

Other service types include NodePort, LoadBalancer, and ExternalName, but unless you explicitly specify one of those, Kubernetes assumes you want a ClusterIP service.