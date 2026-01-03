

| Service Type   | What Port You Use in URL    | Publicly Accessible?  |Notes                                     |
| -------------- | --------------------------- | --------------------- | ----------------------------------------- |
| `NodePort`     | `http://<NodeIP>:nodePort`  | ✅ Yes (with firewall) | You must use the `nodePort` (e.g., 32625) |
| `LoadBalancer` | `http://<EXTERNAL-IP>:port` | ✅ Yes                 | You use the `port`, not the nodePort      |
| `ClusterIP`    | ❌ Not accessible externally | ❌ No                  | For internal cluster use only             |
