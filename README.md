# contactbook_project

Database container is using MYSQL:
1. first we need to be in the root of "database" folder and run this command
<code>docker build -t mysql_db . </code> 
2. Run the database container with this command to expose the container with port 3307
<code> docker run -e MYSQL_ROOT_PASSWORD=1234 -d -p 127.0.0.1:3307:3306 mysql_db:latest </code>

Container for frontend:
1. first we need to be in the root of "frontend" folder and run this command
<code>docker build -t frontend:v1 . </code> 
2. Run the frontend container with this command to expose the container with port 3000
<code> docker run -d -p 3000:80 frontend:v1 </code>
