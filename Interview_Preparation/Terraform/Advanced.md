# 1️⃣ How do you design Terraform for large projects?

Answer:
For large projects, we use modular structure.

👉 We split into:

VPC module
EC2 module
RDS module

👉 Also:

Use separate folders for dev, prod

👉 Simple:
Break code into small reusable modules

# 2️⃣ What are best practices in Terraform?

Answer:
Some important best practices:

Use remote backend (S3)
Enable state locking (DynamoDB)
Use modules
Use variables
Do NOT hardcode values
Keep code in Git

👉 Simple:
Write clean, reusable, and safe code

# 3️⃣ How do you handle multiple environments (dev, prod)?**

Answer:
We use:

Workspaces OR
Separate folders

👉 Example:

dev/
prod/

👉 Also change variables based on environment

👉 Simple:
Same code, different configs

# 4️⃣ What happens if Terraform state file is deleted?

Answer:
Terraform will lose track of infrastructure.

👉 Problems:

It may try to recreate resources
Infrastructure mismatch

👉 Solution:

Always use remote backend (S3 backup)

👉 Simple:
State is very important — never lose it

# 5️⃣ How do you handle drift in Terraform?

Answer:
Drift means manual changes in infrastructure outside Terraform.

👉 Solution:

Run terraform plan → detect changes
Run terraform apply → fix

👉 Simple:
Bring infra back to Terraform control

# 6️⃣ How do you secure Terraform code?

Answer:
We secure by:

Avoid hardcoding secrets
Use IAM roles
Use Vault / Secrets Manager
Restrict access to state file

👉 Simple:
Protect secrets and access

# 7️⃣ How do you reduce cost using Terraform?

Answer:
We can:

Use smaller instance types
Auto shutdown unused resources
Use autoscaling
Destroy unused infra (terraform destroy)

👉 Simple:
Create only what you need

# 8️⃣ How do you debug Terraform issues?

Answer:
Steps:

Check terraform plan

Enable logs:

export TF_LOG=DEBUG
Check error message
Verify credentials

👉 Simple:
Check logs + plan output

# 9️⃣ What is lifecycle block in Terraform?

Answer:
Lifecycle block controls how resources behave during updates.

👉 Example:

prevent_destroy = true
create_before_destroy = true

👉 Simple:
Control resource changes

# 🔟 What is taint in Terraform?

Answer:
Taint marks a resource to force recreate it.

👉 Command:

terraform taint <resource>

👉 Simple:
Destroy and recreate resource

# 1️⃣1️⃣ What is difference between count and for_each?

Answer:

**count:**

Uses index (0,1,2)
Less flexible

**for_each:**

Uses key-value
More flexible

👉 Simple:
for_each is better for real projects

# 1️⃣2️⃣ Can Terraform handle multi-cloud?

Answer:
Yes, Terraform supports multi-cloud environments.

👉 Example:

AWS + Azure + GCP

👉 Using different providers

👉 Simple:
One tool, multiple clouds