# 1️⃣ What are variables in Ansible?

Answer:
Variables are used to store values that can be reused in playbooks.

Example:

vars:
  package_name: nginx

Usage:

name: "{{ package_name }}"

# 2️⃣ Different ways to define variables?

Answer:

Playbook (vars)
Inventory file
Group vars / host vars
Command line (-e)
Role variables

# 3️⃣ What is variable precedence?

Answer:
It defines which variable value takes priority when the same variable is defined in multiple places.

👉 Highest priority:

Extra vars (-e) > playbook vars > inventory vars > default vars

# 4️⃣ What are loops in Ansible?

Answer:
Loops allow executing a task multiple times.

Example:

- name: Install packages
  apt:
    name: "{{ item }}"
    state: present
  loop:
    - nginx
    - git

# 5️⃣ What are conditionals in Ansible?

Answer:
Used to execute tasks based on conditions using when.

Example:

when: ansible_os_family == "Debian"

# 6️⃣ What are handlers in Ansible?

Answer:
Handlers are special tasks that run only when notified.

Example:

notify: restart nginx

Handler:

- name: restart nginx
  service:
    name: nginx
    state: restarted

# 7️⃣ What is notify in Ansible?

Answer:
notify triggers a handler when a task changes the system.

# 8️⃣ What are templates in Ansible?

Answer:
Templates use Jinja2 to create dynamic configuration files.

Example:

nginx.conf.j2

# 9️⃣ What is Jinja2?

Answer:
A templating engine used in Ansible to insert variables dynamically.

Example:

server_name {{ domain_name }};

# 🔟 What is a role in Ansible?

Answer:
A role is a structured way to organize playbooks into reusable components.

Structure:

roles/
  web/
    tasks/
    handlers/
    templates/
    vars/

# 1️⃣1️⃣ Why do we use roles?

Answer:

Reusability
Clean structure
Easy management
Scalable automation

# 1️⃣2️⃣ What is include/import in Ansible?

Answer:

include_tasks → dynamic
import_tasks → static

Used to reuse tasks across playbooks.

# 1️⃣3️⃣ What is register in Ansible?

Answer:
Used to store output of a task.

Example:

- command: ls
  register: output

# 1️⃣4️⃣ What is debug module?

Answer:
Used to print variables or messages.

- debug:
    var: output

# 1️⃣5️⃣ What is facts in Ansible?

Answer:
Facts are system information collected from managed nodes.

Example:

OS
IP
CPU

# 1️⃣6️⃣ How to disable fact gathering?

Answer:

gather_facts: no

# 1️⃣7️⃣ What is tags in Ansible?

Answer:
Tags allow running specific tasks.

ansible-playbook playbook.yml --tags "install"

# 1️⃣8️⃣ What is Ansible Vault?

Answer:
Used to encrypt sensitive data like:

passwords
API keys

# 1️⃣9️⃣ How to create encrypted file?

Answer:

ansible-vault create secret.yml

2️⃣0️⃣ What is difference between command and shell module?

Answer:

command	shell
No shell	Uses shell
Safe	Less secure
No pipes	Supports pipes
2️⃣1️⃣ What is lineinfile module?

Answer:
Used to add/update/remove a line in a file.

2️⃣2️⃣ What is copy module?

Answer:
Used to copy files from control node to managed node.

2️⃣3️⃣ What is file module?

Answer:
Used to manage file properties like:

permissions
ownership
directories
2️⃣4️⃣ What is service module?

Answer:
Used to manage services (start/stop/restart).

🧠 Interview Tip (VERY IMPORTANT)

If asked:
👉 “How do you manage complex infrastructure in Ansible?”

Answer:

“I use roles, variables, templates, and handlers to create reusable and scalable automation.”



