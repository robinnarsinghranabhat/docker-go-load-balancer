## Python server and minimal load balancer in Go deployed as a microservice.


### HOW TO RUN : 
1. `./scripts.sh` 
2. host-machine browser is our client. Try hitting `http://localhost:8000/` multiple times. Each time, we have request routed by load-balancer to different server container.


### HOW IT WORKS : 
Inside python-flask-server, we have dockerfile to create server container.
We spin up two server containers. 

Inside go-load-balancer, we have dockerfile to create a simple load-balancer.
Client requests are handled by this server. And it forwards them to each server in round-robin fashion.   

NOTE : Although not practical, for demonstration purpose, all three containers are in same network-interface i.e. they emulate as if they are on a same network (imagine 3 machines connected to a router ).
