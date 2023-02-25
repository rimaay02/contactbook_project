# contactbook_project

API container is using MYSQL and Go:
1. first we need to be in the root of "backend" folder
2. run this command
<code>docker compose up </code> 

Container for frontend:
1. first we need to be in the root of "frontend" folder and run this command
<code>docker build -t frontend:v1 . </code> 
2. Run the frontend container with this command to expose the container with port 3000
<code> docker run -d -p 3000:80 frontend:v1 </code>
