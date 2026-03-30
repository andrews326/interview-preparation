# 1️⃣ What is Terraform?

Answer:
Terraform is an Infrastructure as Code (IaC) tool used to create, manage, and update infrastructure automatically.

👉 In simple words:
Instead of creating servers manually, we write code and Terraform creates everything.

# 2️⃣ Why do we use Terraform?

Answer:
We use Terraform to:

Automate infrastructure creation
Avoid manual errors
Save time
Maintain same setup in all environments

👉 Simple: Write once → create anywhere

# 3️⃣ What is Infrastructure as Code (IaC)?**

Answer:
Infrastructure as Code means managing infrastructure using code instead of manual steps.

👉 Example:

Without IaC → manually create EC2
With IaC → write code → Terraform creates EC2

# 4️⃣ What is a Terraform Provider?**

Answer:
A provider is a plugin that allows Terraform to interact with cloud services.

👉 Examples:

AWS provider
Azure provider
GCP provider

👉 Simple: Provider = connection between Terraform and cloud

# 5️⃣ What is a Terraform Resource?**

Answer:
A resource is any infrastructure component created using Terraform.

👉 Examples:

EC2 instance
S3 bucket
VPC

👉 Simple: Resource = actual thing you create

# 6️⃣ What is a Terraform Module?**

Answer:
A module is a collection of Terraform files used together.

👉 Simple:

Small reusable code block

👉 Example:

One module for EC2
One module for VPC

# 7️⃣ What is terraform init?**

Answer:
terraform init is used to initialize the Terraform project.

👉 It will:

Download providers
Prepare working directory

👉 Simple: First command to start Terraform

# 8️⃣ What is terraform plan?**

Answer:
terraform plan shows what changes Terraform will make.

👉 It does NOT create anything
👉 Just preview

👉 Simple: Dry run (check before apply)

# 9️⃣ What is terraform apply?**

Answer:
terraform apply is used to create or update infrastructure.

👉 It executes the plan

👉 Simple: Real execution command

# 🔟 What is terraform destroy?**

Answer:
terraform destroy is used to delete all created resources.

👉 Simple: Clean everything



👉 **Remember flow:**  init → plan → apply → destroy