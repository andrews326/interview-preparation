


sqs simple queue service
   messages will store queue before send to receiver
   message should be within 256kb
   we only keep message 14 days in queue
   message keep maximum 14 days

DLQ= dead letter queue
     messages which is not send to user  after 14 days it will be store on DLQ(we should enable DLQ otherwise it will delete) wich is retention period of aws console for sqs

 messages share only sqs user 
 like one to one chat



 sns 
   one to many 
   components: topic,subscribers  
   sender send the messages to the topic
   the messages goes to topic
   then the subscribers received the messages who are subscribed this topic 

   there are two types to sending messages to sns
     1. FIFO
          subscription protocols: sqs
     2. standated:
          subscription protocols   : sqs,lamda,data firehouse,https,email,mobile sms


 cloud watch :
   this is monitoring tool. we can monitor ec2 by cloud watch   



#### VPC  
 ip address

   ipv4 (currently we use it)
   ipv6  

   public ip  > staic ips which is should be purchased
   private ip  > reusable(internal communication)    

**prerequests**
   route tables
   internet gateway
   nat gateway
   cidr block
   security group
   subnet   

1 tier architecture  => webserver 
                         public subnet
2 tier architecture  => webserver / app server + database  
                              public subnet        private subnet
3 tier architecture  => webserver    +  app server    + database  
                       public subnet   private subnet   private subnet


 **vpc** 
   subnet
     public subnet
     private subnet
   route table
     public rt
     private rt
   internet gateway ==>> to give internet connection to the public subnet
   nat gateway  ==>> to give internet connection to the private subnet
   security group
     public sg
     private sg    


  # we can attah one vpc to internet gateway   
  # route table is a hub to connect internet gateway and public subnet

  # we should connect internet gateway to public route table
  # we should associate public subnet in public route table
  # we should not give internet connection for private subnet its need secure connection (we should user nat gateway)
  # nat gateway internet connection will from public subnet

  # we should mention public securiy group in outbound rule -> destination for private securiy group

  **creating server**
    webserver
       choose myVPC
              public subnet
              enable assign public ip
              public security group

    db-server
       choose myVPC
              private subnet
              enable/disable assign public ip (we could choose either)
              public security group   



first we need to delete nat gateway (its take some time)
then release the elastic ip
then we could delete everything by single shot while we delete vpc