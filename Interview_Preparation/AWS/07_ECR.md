🐳 AWS ECR — Interview Questions & Answers
🌱 Basic Level

1. What is AWS ECR?
ECR (Elastic Container Registry) is a fully managed Docker container registry that makes it easy to store, manage, and deploy Docker container images.

2. How does ECR integrate with Docker?
You can use Docker CLI commands (docker push, docker pull) with ECR repositories to store and retrieve images.

3. What is a repository in ECR?
A repository is a collection of Docker images. Each repository can have multiple image tags representing different versions.

4. How do you authenticate Docker to ECR?
Use AWS CLI with aws ecr get-login-password to authenticate Docker to push or pull images.

5. What is the difference between public and private ECR repositories?
Private repositories are accessible only with AWS credentials. Public repositories can be accessed by anyone without authentication.

6. How do you tag images in ECR?
You tag images using docker tag before pushing, for example:
docker tag myapp:latest 123456789012.dkr.ecr.us-east-1.amazonaws.com/myapp:latest.

7. Can you store multiple versions of an image in ECR?
Yes, each image can have multiple tags representing different versions.

8. What regions support ECR?
ECR is available in most AWS regions, and you should use the same region as your deployment resources for lower latency.

9. What is image scanning in ECR?
ECR can scan images for known vulnerabilities using Amazon Inspector.

10. How do you delete an image from ECR?
Use aws ecr batch-delete-image or delete it manually from the AWS Console.

⚙️ Intermediate Level

1. How do you integrate ECR with Jenkins?
Jenkins can push and pull Docker images from ECR using AWS credentials and Docker CLI. You can also use the AWS ECR plugin for Jenkins.

2. How do you make images available across accounts?
Use repository policies to grant cross-account permissions, allowing users or services in another AWS account to pull images.

3. How do you automate Docker image builds and pushes to ECR?
Use CI/CD pipelines (Jenkins, GitHub Actions, CodePipeline) with commands: docker build, docker tag, docker push.

4. What is image lifecycle policy in ECR?
Lifecycle policies automatically expire old images based on rules, helping manage storage costs.

5. How can you secure ECR repositories?
Use IAM policies to control access and enable encryption at rest with AWS KMS.

6. What is image replication in ECR?
ECR can replicate images to other regions automatically for high availability and disaster recovery.

7. How do you reduce image storage costs?
Use lifecycle policies to remove old or unused images and compress image layers when building.

8. How can ECR integrate with ECS or Kubernetes?
ECS tasks or Kubernetes pods can pull images directly from ECR using IAM roles for authentication.

9. What is the difference between ECR and DockerHub?
ECR is AWS-managed and tightly integrated with AWS services. DockerHub is public and independent but supports private repositories too.

10. How do you handle image versioning in ECR?
Use tags like v1.0, v1.1, or latest to version your images consistently.

🚀 Advanced Level

1. How do you secure image pull for Kubernetes deployments?
Use IAM roles for service accounts (IRSA) in EKS or attach an EC2 role with AmazonEC2ContainerRegistryReadOnly policy for kubelets to pull images.

2. How do you scan images in ECR for vulnerabilities automatically?
Enable image scanning on push or schedule periodic scans using Amazon Inspector integration.

3. How can you implement CI/CD pipelines using ECR for microservices?
Build Docker images per microservice, push to ECR, tag appropriately, and deploy with Kubernetes (kubectl) or ECS tasks in an automated Jenkins pipeline.

4. How do you replicate ECR images across multiple AWS regions?
Configure cross-region replication rules in ECR to copy images to target regions automatically.

5. How do you handle rollbacks using ECR images?
Use previous image tags in deployment manifests (like v1.0) to roll back pods or ECS services.

6. How do you audit access to ECR repositories?
Enable CloudTrail logging for ECR API calls to track who pushed, pulled, or deleted images.

7. How do you manage permissions for multiple teams?
Use fine-grained IAM policies, repository-level policies, and service roles to give only the necessary access per team.

8. How do you optimize ECR for large-scale microservices deployment?

Use minimal base images

Remove unnecessary layers

Use multi-stage builds

Enable cross-region replication for availability

9. How do you integrate ECR with Kubernetes secrets?
Create an imagePullSecret containing ECR credentials so pods can pull private images securely.

10. What are best practices for using ECR in production CI/CD?

Use automated builds and scans

Apply lifecycle policies

Secure access via IAM roles, not static keys

Tag images with semantic versioning

Monitor usage and costs