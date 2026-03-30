# 1️⃣ What is Dynamic Inventory in Ansible?

Answer:
Dynamic inventory automatically fetches server details from external sources like:

AWS (EC2)
Azure
GCP

👉 Instead of static /etc/ansible/hosts

Example:

ansible-inventory -i aws_ec2.yaml --list

# 2️⃣ Why do we use Dynamic Inventory?

Answer:
Because:

Servers are created/destroyed dynamically (cloud)
No need to manually update inventory
Supports auto-scaling environments

# 3️⃣ How does Ansible work internally?

Answer:

Reads inventory
Connects via SSH
Pushes modules to target node
Executes modules
Returns output

👉 No agent required

# 4️⃣ What is Ansible Galaxy?

Answer:
Ansible Galaxy is a repository for:

Roles
Collections

Command:

ansible-galaxy install geerlingguy.nginx

# 5️⃣ What are Collections in Ansible?

Answer:
Collections are bundles of:

Modules
Plugins
Roles

Used to organize and distribute content.

# 6️⃣ What is Ansible Tower / AWX?

Answer:

AWX → Open-source UI for Ansible
Ansible Tower → Enterprise version

Features:

Web UI
RBAC
Scheduling
API

# 7️⃣ How do you handle secrets securely in Ansible?

Answer:

Use Ansible Vault
Store encrypted variables
Avoid hardcoding secrets

# 8️⃣ What is rolling update in Ansible?

Answer:
Updating servers in batches instead of all at once.

Example:

serial: 2

👉 Updates 2 servers at a time

# 9️⃣ What is parallel execution in Ansible?

Answer:
Ansible runs tasks on multiple hosts simultaneously.

Controlled by:

forks = 5

# 🔟 How to improve Ansible performance?

Answer:

Increase forks
Disable fact gathering
Use pipelining
Use async tasks
Optimize inventory

# 1️⃣1️⃣ What is pipelining in Ansible?

Answer:
Reduces SSH overhead by avoiding multiple connections.

Enable:

pipelining = True

# 1️⃣2️⃣ What are async tasks?

Answer:
Run long-running tasks in background.

Example:

async: 60
poll: 0

# 1️⃣3️⃣ What is delegate_to?

Answer:
Runs a task on a different host.

Example:

delegate_to: localhost

# 1️⃣4️⃣ What is run_once?

Answer:
Runs a task only once even if multiple hosts exist.

# 1️⃣5️⃣ What is block in Ansible?

Answer:
Groups multiple tasks together.

block:
  - name: task1
  - name: task2

# 1️⃣6️⃣ What is rescue and always?

Answer:
Used for error handling.

block:
  - name: task
rescue:
  - name: handle error
always:
  - name: always run

# 1️⃣7️⃣ How do you debug Ansible playbooks?

Answer:

Use -vvv verbose mode
Use debug module
Check logs
Use ansible-playbook --check

# 1️⃣8️⃣ What is check mode?

Answer:
Dry-run mode — shows what will change without applying.

ansible-playbook playbook.yml --check

# 1️⃣9️⃣ What is diff mode?

Answer:
Shows differences before and after changes.

--diff

# 2️⃣0️⃣ How does Ansible integrate with CI/CD?

Answer:

Used in Jenkins pipelines
Automates deployment
Runs playbooks after build

# 2️⃣1️⃣ Ansible vs Terraform?

Answer:
-------------------------------
Ansible	        Terraform     -
-------------------------------
Configuration	Infrastructure
Procedural	    Declarative
Agentless	    Uses providers

👉 Use both together in real projects

# 2️⃣2️⃣ Ansible vs Chef/Puppet?

Answer:

Ansible → Agentless, simple
Chef/Puppet → Agent-based, complex

# 2️⃣3️⃣ How Ansible works with Docker?

Answer:

Can build images
Run containers
Manage Docker services

# 2️⃣4️⃣ How Ansible works with Kubernetes?

Answer:

Deploy manifests
Manage clusters
Use k8s module

# 2️⃣5️⃣ What is idempotency in real scenario?

Answer:
If nginx is already installed:

Ansible will NOT reinstall it
It ensures desired state only

*** Real-Time Scenario Questions ***

# 2️⃣6️⃣ How do you deploy an application using Ansible?

Answer:

Use playbook
Install dependencies
Copy files
Configure services
Restart application

# 2️⃣7️⃣ How do you handle 100+ servers?

Answer:

Use inventory groups
Use roles
Use parallel execution

# 2️⃣8️⃣ How do you manage different environments?

Answer:

Use separate inventory files
Use variables per environment

# 2️⃣9️⃣ What happens if a task fails?

Answer:

Playbook stops by default
Use ignore_errors or rescue

# 3️⃣0️⃣ How do you ensure zero downtime deployment?

Answer:

Use rolling updates (serial)
Use load balancer
Update servers in batches