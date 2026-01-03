🌐 AWS VPC — Interview Questions & Answers
🌱 Basic Level

1. What is a VPC?
A VPC (Virtual Private Cloud) is a virtual network dedicated to your AWS account where you can launch AWS resources with complete control over IP addressing, subnets, route tables, and network gateways.

2. What is a subnet?
A subnet is a segment of a VPC’s IP address range where you can place EC2 instances or other resources. Subnets can be public or private.

3. What is the difference between public and private subnets?
Public subnet = resources have direct access to the internet via an Internet Gateway.
Private subnet = resources do not have direct internet access; can use NAT Gateway to reach the internet.

4. What is an Internet Gateway (IGW)?
IGW is a horizontally scaled, redundant gateway that allows communication between instances in your VPC and the internet.

5. What is a NAT Gateway?
A NAT Gateway allows instances in private subnets to access the internet for updates or downloads while keeping them unreachable from the internet.

6. What is a route table?
A route table contains rules that determine how network traffic is directed in your VPC, including routes to subnets, IGWs, NAT Gateways, and peering connections.

7. What is a VPC peering connection?
It allows two VPCs (same or different AWS accounts) to communicate with each other privately without going over the internet.

8. What is a security group in VPC?
A security group acts as a virtual firewall controlling inbound and outbound traffic at the instance level.

9. What is a Network ACL (NACL)?
A NACL is an optional firewall at the subnet level, controlling traffic entering or leaving the subnet.

10. What is the default VPC?
AWS provides a default VPC in each region so you can launch instances without manual network configuration.

⚙️ Intermediate Level

1. What is CIDR in VPC?
CIDR (Classless Inter-Domain Routing) defines the IP address range for a VPC or subnet (e.g., 10.0.0.0/16).

2. How do public and private subnets differ in routing?
Public subnets have a default route pointing to the Internet Gateway. Private subnets route external traffic through a NAT Gateway.

3. What is an Elastic IP in a VPC?
A static, public IPv4 address that can be associated with an instance or NAT Gateway for consistent internet access.

4. How do you connect VPCs across accounts or regions?
Use VPC peering, AWS Transit Gateway, or VPN connections for cross-account or cross-region communication.

5. How do you implement high availability in VPC?
Deploy resources across multiple Availability Zones and use subnets in each AZ with redundancy for critical services.

6. What is VPC flow logs?
Flow logs capture information about the IP traffic going to and from network interfaces in a VPC. Useful for monitoring and troubleshooting.

7. How do you restrict access to a specific EC2 instance in a VPC?
Use Security Groups for instance-level rules and Network ACLs for subnet-level rules.

8. What is a VPN Gateway?
A VPN Gateway enables secure communication between an on-premises network and your VPC over an encrypted VPN connection.

9. How do you allow private instances to access S3 without using the internet?
Use a VPC endpoint (Gateway type) for S3 so traffic stays within AWS’s network.

10. How do you implement multi-tier architecture in a VPC?
Place web servers in public subnets, application servers and databases in private subnets, and control access via security groups and NACLs.

🚀 Advanced Level

1. How do you connect multiple VPCs efficiently?
Use AWS Transit Gateway for centralized routing or set up multiple VPC peering connections with proper route tables.

2. How do you monitor VPC network traffic?
Enable VPC Flow Logs and integrate with CloudWatch or S3 for analysis.

3. How do you implement network segmentation for security?
Use multiple subnets, security groups, and NACLs to separate layers like web, application, and database.

4. How can you implement hybrid cloud connectivity?
Use VPN Gateway for site-to-site VPN or Direct Connect for private high-speed links between on-premises and AWS.

5. How do you optimize VPC for cost and performance?
Right-size subnets, use NAT Gateway only when necessary, consolidate VPCs with Transit Gateway, and enable flow logs selectively.

6. How do you secure data-in-transit within a VPC?
Use TLS for application-level traffic, enforce security groups, and encrypt traffic for VPC endpoints if available.

7. How do you implement zero-trust networking in AWS?
Use least privilege security groups, NACLs, IAM policies, and private subnets for all sensitive workloads.

8. How do you handle overlapping CIDR ranges for VPC peering?
You cannot peer VPCs with overlapping CIDR blocks; plan non-overlapping IP ranges before connecting.

9. How do you automate VPC creation?
Use Terraform, CloudFormation, or AWS CLI scripts to define subnets, route tables, IGWs, and security groups programmatically.

10. What are best practices for VPC security?

Use private subnets for critical workloads

Apply least privilege security groups

Enable flow logs and monitoring

Limit public exposure and use NAT Gateways strategically