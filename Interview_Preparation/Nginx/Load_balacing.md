🎓 LECTURE 6

🧠 Why do we need load balancing?

Imagine:

One backend container cannot handle all traffic

You run multiple copies of the same app

🔁 Architecture

Client
  ↓
NGINX
  ↓
---------------------
|       |           |
App-1   App-2      App-3

NGINX decides where each request goes.


🧩 NGINX Load Balancing methods
1️⃣ Round Robin (default)

Requests distributed evenly

2️⃣ Least Connections

Server with fewer active connections

3️⃣ IP Hash

Same client → same backend


🧠 NGINX concept: upstream

This is the heart of load balancing.

upstream backend_servers {
    server backend1;
    server backend2;
    server backend3;
}


🔁 Using upstream in server block
server {
    listen 9000;

    location / {
        proxy_pass http://backend_servers;
    }
}
