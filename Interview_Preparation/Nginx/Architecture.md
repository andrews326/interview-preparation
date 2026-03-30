👉 Important DevOps idea:
Users talk only to NGINX, never directly to app servers.


📊 Simple Architecture Diagram (Real-world)

Browser / Mobile App
        |
        |  https://example.com
        |
      NGINX
  (Reverse Proxy)
        |
        | forwards request
        |
 -------------------------
 |           |           |
App-1      App-2      App-3
(Golang)   (Golang)   (Golang)


🧠 What NGINX is doing here?

Accepts client requests

Decides where to send traffic

Hides backend servers

Improves performance & security

🧩 Very Simple NGINX Config Example

server {
    listen 80;

    server_name example.com;

    location / {
        proxy_pass http://localhost:8080;
    }
}


🔍 Line-by-line explanation (slow & clear)
server { }

➡️ Defines a virtual server (one website / one app)

listen 80;

➡️ NGINX listens on port 80 (HTTP)

User hits:

http://example.com

server_name example.com;

➡️ Domain name for this app

location / { }

➡️ “For all requests coming to / ”

Examples:

/

/api

/login

proxy_pass http://localhost:8080;