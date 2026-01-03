
# KMS = Key Management System
  its used in s3 (encrypion enable disable)

# there are two type of KMS
1. Symetric => used for windows
2. asymetric => used for linux

**Syemetric**  = to lock(encrypt) and unlock(decrypt) with same key(symetric)
**Asymetric**  = its used  different keys to encrypt and decrypt (public and private key) public key for encryption and private key for decryption. after the encrypted data which is called cipher text 
plain text (publick key) cipher text (private key)  plain text

## algorithm for asymetric 
   algorithm means rule to use   

## cryptography
    which means encryption and decryption      


    admin = ec2
    user  = lamda               


## CLOUD TRAIL
   monitioring the aws user activities what they are doing in console   


**AWS KMS (Key Management Service)**

Q1: What is AWS KMS?
A1: AWS KMS is a service that helps you create and control encryption keys for your data. It’s used to protect sensitive information by encrypting it.

Q2: Why do we need AWS KMS?
A2: AWS KMS ensures your data is encrypted and secure, preventing unauthorized access to sensitive information.

Q3: Can I manage my own encryption keys?
A3: Yes, you can create and manage your own keys or use AWS-managed keys.

**AWS CloudTrail**

Q1: What is AWS CloudTrail?
A1: AWS CloudTrail records all actions (API calls) in your AWS account, providing a history of events to track user activity.

Q2: Why is CloudTrail important?
A2: CloudTrail helps with security and auditing by letting you see who did what and when, making it easier to troubleshoot or investigate issues.

Q3: How long does CloudTrail store logs?
A3: CloudTrail stores logs for 90 days by default, but you can configure it to store logs in S3 for as long as you want.

**AWS Organizations**

Q1: What is AWS Organizations?
A1: AWS Organizations lets you manage multiple AWS accounts from a single place. You can organize accounts, apply policies, and manage billing.

Q2: Why use AWS Organizations?
A2: It helps you organize accounts, manage permissions, and apply security policies across all your AWS accounts.

Q3: Can I apply policies to all accounts in an organization?
A3: Yes, you can use Service Control Policies (SCPs) to apply security and permissions across all accounts.

**AWS WAF (Web Application Firewall)**

Q1: What is AWS WAF?
A1: AWS WAF is a security service that protects your web applications from attacks like SQL injection and cross-site scripting (XSS).

Q2: Why use AWS WAF?
A2: AWS WAF helps secure your website by filtering out malicious web traffic and preventing common attacks.

Q3: How does AWS WAF work?
A3: AWS WAF uses rules to block harmful traffic, such as requests from suspicious IP addresses or known attack patterns.    