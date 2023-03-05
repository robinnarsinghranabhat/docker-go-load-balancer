# build and run two server containers
docker build -t python-server .
docker run --name s1 -p 2222:8000 -d python-server:latest
docker run --name s2 -p 3333:8000 -d python-server:latest


# build and run loab-balancer server in another container
docker build -t go-load-balancer .
docker run --name lb -p 8000:8000 -d go-load-balancer:latest


# Put all load-balancer as well as 
# containers inside same namespace.
docker network create lb-server
docker network connect lb-server s1
docker network connect lb-server s2
docker network connect lb-server lb

