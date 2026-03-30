# 1️⃣ What is Terraform State?

Answer:
Terraform state is a file that stores the current status of infrastructure.

👉 It keeps details like:

What resources are created
Their configuration
Their IDs

👉 Simple:
State = Terraform memory of your infrastructure

# 2️⃣ What is terraform.tfstate file?

Answer:
terraform.tfstate is the default file where Terraform stores state.

👉 Example:

EC2 instance ID
S3 bucket name

👉 Simple:
It tracks what Terraform already created

# 3️⃣ What is Remote Backend?

Answer:
Remote backend is used to store Terraform state in a remote location instead of local.

👉 Examples:

AWS S3
Azure Storage

👉 Why?

Team collaboration
Avoid state loss

👉 Simple:
Store state safely in cloud

# 4️⃣ Why do we use state locking?

Answer:
State locking prevents multiple users from changing infrastructure at the same time.

👉 Example:

2 people run terraform apply → conflict

👉 Solution:

Lock state using DynamoDB (in AWS)

👉 Simple:
Avoid conflicts in team

# 5️⃣ What is a Variable in Terraform?

Answer:
A variable is used to pass input values into Terraform code.

👉 Example:

instance_type = t2.micro

👉 Simple:
Variables make code reusable and dynamic

# 6️⃣ What is Output in Terraform?

Answer:
Output is used to display values after Terraform execution.

👉 Example:

Public IP of EC2

👉 Simple:
Output = result you want to see

# 7️⃣ What is a Data Source in Terraform?

Answer:
A data source is used to fetch existing resources from cloud.

👉 Example:

Get existing VPC
Get AMI ID

👉 Simple:
Data source = read existing resources (not create)

# 8️⃣ What is dependency in Terraform?

Answer:
Dependency means one resource depends on another resource.

👉 Example:

EC2 depends on VPC

👉 Terraform handles automatically, but we can use:

depends_on

👉 Simple:
Order of resource creation

# 9️⃣ What is terraform refresh?

Answer:
terraform refresh updates the state file with real infrastructure data.

👉 Example:

If someone changed resource manually

👉 Simple:
Sync state with real world

# 🔟 What is terraform import?

Answer:
terraform import is used to bring existing resources into Terraform management.

👉 Example:

Existing EC2 → import → Terraform manages it

👉 Simple:
Add already-created resources into Terraform

# 1️⃣1️⃣ What is workspace in Terraform?

Answer:
Workspace is used to manage multiple environments using same code.

👉 Example:

dev
staging
prod

👉 Simple:
Same code, different environments

# 1️⃣2️⃣ How do you handle secrets in Terraform?

Answer:
We should NOT hardcode secrets.

👉 Use:

Environment variables
AWS Secrets Manager
Vault

👉 Simple:
Keep secrets outside code