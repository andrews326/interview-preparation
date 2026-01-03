1. What is AWS EFS?

Answer:
AWS EFS is a file storage service that automatically scales. It's used to store and share files between multiple EC2 instances in AWS, like a network drive.

2. What are the benefits of using EFS?

Answer:

Scalable: Grows as you store more files.

No management: AWS takes care of the storage.

Shared access: Multiple EC2 instances can access the same data.

3. How does EFS differ from Amazon S3?

Answer:

EFS is a file system that works like a network drive, ideal for apps needing file-based storage.

S3 is for object storage (storing files like images, backups) and doesn’t behave like a file system.

4. What are the use cases of AWS EFS?

Answer:

Shared file storage for web servers.

Big data and analytics.

Content management systems.

Development and testing environments.

5. What types of file systems does EFS support?

Answer:
EFS supports NFS (Network File System), which works with Linux-based systems (like EC2).

6. How do you mount EFS on an EC2 instance?

Answer:

Install NFS tools (sudo yum install nfs-utils).

Mount the file system:

sudo mount -t nfs -o nfsvers=4.1 <EFS-DNS-Name>:/ <mount-point>

7. What is the maximum throughput of AWS EFS?

Answer:
EFS can handle bursting throughput (scales automatically) or provisioned throughput (you set the speed).

8. What are EFS storage classes?

Answer:

Standard: For frequently used files.

Infrequent Access: For rarely used files, cheaper.

9. How does EFS handle data durability?

Answer:
EFS replicates your data across multiple Availability Zones for high availability and durability.

10. How does EFS ensure security?

Answer:

Encryption for data at rest and in transit.

IAM and NFS permissions to control access.

11. What is the maximum size of an EFS file system?

Answer:
EFS is virtually unlimited and automatically grows as you add files.

12. What is the difference between EFS and EBS?

Answer:

EFS: Shared file system, used by multiple EC2 instances.

EBS: Block storage, typically attached to one EC2 instance.

13. Can EFS be used with Windows instances?

Answer:
No, EFS works with Linux instances. For Windows, use Amazon FSx.

14. How do you monitor EFS performance?

Answer:
Use CloudWatch to track EFS performance like throughput and storage usage.

15. Can EFS be used across multiple regions?

Answer:
No, EFS is only available within a single region. For cross-region data, use DataSync.

Summary:

EFS = A scalable file system for sharing files between EC2 instances.

S3 = Object storage for things like backups, photos, etc.

EFS is Linux-only and uses NFS.