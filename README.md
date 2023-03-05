## Python server and minimal load balancer in Go deployed as a microservice.


### HOW TO RUN : 
1. `./scripts.sh` 
2. From host-machine browser, try hitting `http://localhost:8000/` multiple times. EAch time, we have request handled by different server container.


### HOW IT WORKS : 
Inside python-flask-server, we have dockerfile to create server container.
We spin up two server containers. 

Inside go-load-balancer, we have dockerfile to create a simple load-balancer.
Client requests are handled by this server. And it forwards them to each server in round-robin fashion.   

NOTE : Although not practical, for demonstration purpose, We put all three networks in a same network-interface.
