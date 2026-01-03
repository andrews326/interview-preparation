🔐 AWS IAM — Interview Questions & Answers
🌱 Basic Level

1. What is IAM?
IAM (Identity and Access Management) is a service that controls who can access AWS resources and what actions they can perform.

2. What are IAM users?
An IAM user represents a single person or application that interacts with AWS. Each user has unique credentials (username, password, access keys).

3. What are IAM groups?
Groups are collections of users. Permissions applied to a group are inherited by all its users.

4. What are IAM policies?
Policies are JSON documents that define what actions are allowed or denied for a user, group, or role.

5. What is an IAM role?
A role is like a temporary identity with a set of permissions that can be assumed by AWS services or users.

6. What is the root account?
The root account is the main AWS account owner with unrestricted access. It should be used only for account setup, not daily tasks.

7. What is MFA in IAM?
MFA (Multi-Factor Authentication) adds an extra layer of security by requiring a code from a device or app in addition to your password.

8. How do you create an IAM user?
Through the AWS Console, CLI, or SDK — specifying username, permissions, and access type (Console or Programmatic).

9. What is the difference between IAM user and IAM role?
A user has long-term credentials; a role provides temporary access with permissions that can be assumed by entities.

10. What is the IAM password policy?
Rules that enforce password strength — like minimum length, complexity, and rotation frequency.

⚙️ Intermediate Level

1. How do IAM policies work?
Policies are evaluated based on explicit allow and deny rules. The final decision follows: explicit deny > allow > default deny.

2. What is the principle of least privilege?
Give users and services only the permissions they need to perform their job — nothing more.

3. What is a managed policy?
Predefined policies provided by AWS (AWS-managed) or created by you (customer-managed) for easy reuse.

4. What are inline policies?
Policies directly attached to a single user, group, or role — not reusable like managed policies.

5. What is the difference between AWS-managed and customer-managed policies?
AWS-managed = created and maintained by AWS.
Customer-managed = created and controlled by you.

6. How do you allow an EC2 instance to access S3 securely?
Attach an IAM role to the EC2 instance with an S3 access policy — no need for credentials in the app.

7. What is IAM federation?
It allows users from external identity systems (like Google, Active Directory, or Okta) to access AWS without creating IAM users.

8. What are IAM access keys?
Access keys (ID and Secret) are used for programmatic access (CLI, SDK, API). They should be rotated regularly.

9. How do you audit IAM activities?
Use AWS CloudTrail to monitor and log all IAM actions like user creation, policy changes, or login attempts.

10. How do you restrict console login access?
Use password policies, MFA, and limit “AWS Management Console access” when creating IAM users.

🚀 Advanced Level

1. How do you design IAM for a large organization?
Use groups, roles, and customer-managed policies. Combine IAM with AWS Organizations and Service Control Policies (SCPs).

2. What are Service Control Policies (SCPs)?
SCPs are policies used in AWS Organizations to set permission boundaries for accounts — they apply organization-wide.

3. How do you secure IAM roles from being abused?
Use least privilege, restrict sts:AssumeRole permissions, and monitor with CloudTrail.

4. What is a permission boundary?
It defines the maximum permissions an IAM role or user can have — useful for delegating control safely.

5. How do you give cross-account access?
Create a role in Account A and allow a user from Account B to assume it using the trust policy.

6. How do you enforce MFA for specific actions?
Use conditions like "aws:MultiFactorAuthPresent": "true" in the IAM policy.

7. How can you test IAM policy permissions?
Use the IAM Policy Simulator or AWS CLI simulate-principal-policy to validate before applying.

8. How do you temporarily elevate permissions for a user?
Use roles with STS (Security Token Service) for temporary access rather than permanent permission changes.

9. How do you protect IAM from accidental deletions or changes?
Use CloudTrail monitoring, IAM Access Analyzer, and restrict IAM:* permissions to admins only.

10. What are IAM best practices?

Never use the root account for daily tasks

Enable MFA everywhere

Rotate access keys

Apply least privilege

Use IAM roles for applications instead of static credentials

Monitor with CloudTrail and Access Analyzer