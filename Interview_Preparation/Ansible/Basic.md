# 1️⃣ What is Ansible?

Answer:
Ansible is an open-source automation tool used for:

Configuration management
Application deployment
Infrastructure provisioning

👉 It automates tasks across multiple servers using simple YAML playbooks.

# 2️⃣ Why do we use Ansible?

Answer:
We use Ansible to:

Automate repetitive tasks
Reduce manual errors
Manage multiple servers efficiently
Ensure consistent configuration

# 3️⃣ What are the advantages of Ansible?

Answer:

Agentless (no installation on target nodes)
Simple YAML syntax
Easy to learn
Idempotent (safe to run multiple times)
Uses SSH (secure)
# 4️⃣ What is agentless architecture?

Answer:
Agentless means:

Ansible does not require any software installation on managed nodes.

👉 It uses SSH to connect and execute tasks.

# 5️⃣ How does Ansible connect to remote machines?

Answer:
Ansible connects using:

SSH (Linux)
WinRM (Windows)

# 6️⃣ What is Ansible Control Node?

Answer:
The Control Node is the machine where:

Ansible is installed
Playbooks are executed

👉 It controls all managed nodes.

# 7️⃣ What are Managed Nodes?

Answer:
Managed Nodes are:

Remote servers that Ansible manages and configures.

# 8️⃣ What is Inventory in Ansible?

Answer:
Inventory is a file that contains:

List of servers
Group of servers

Example:

[web]
server1
server2

# 9️⃣ What are modules in Ansible?

** Answer **:
Modules are:

Small programs that perform specific tasks like installing packages, copying files, etc.

Examples:

apt
yum
copy
service

# 🔟 What is a task in Ansible?

Answer:
A task is:

A single action executed using a module.

Example:

- name: Install nginx
  apt:
    name: nginx
    state: present

# 1️⃣1️⃣ What is a Playbook?

Answer:
A Playbook is:

A YAML file that contains multiple tasks to automate operations on servers.

# 1️⃣2️⃣ What is YAML in Ansible?

Answer:
YAML is:

A human-readable data format used to write Ansible playbooks.

# 1️⃣3️⃣ Structure of a Playbook?

Answer:

- hosts: all
  tasks:
    - name: Install nginx
      apt:
        name: nginx
        state: present

# 1️⃣4️⃣ Difference between ad-hoc and playbook?

Answer:

Ad-hoc	Playbook
One-time command	Reusable
Not structured	Structured
Quick tasks	Complex automation

# 1️⃣5️⃣ How do you run a playbook?

Answer:

ansible-playbook -i inventory.ini playbook.yml

# 1️⃣6️⃣ What is ansible.cfg?

Answer:
Configuration file used to:

Define default settings
Inventory location
SSH options
1️⃣7️⃣ Default inventory file location?

Answer:

/etc/ansible/hosts
1️⃣8️⃣ What is idempotency in Ansible?

Answer:
Idempotency means:

Running the same playbook multiple times will not change the system if it is already in the desired state.

1️⃣9️⃣ What is an ad-hoc command?

Answer:
A one-line command used for quick tasks without playbooks.

Example:

ansible all -m ping
2️⃣0️⃣ Example of ping command?

Answer:

ansible all -m ping

👉 Used to check connectivity between control node and managed nodes.