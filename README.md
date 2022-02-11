# go-service-registry

## Availabe endpoints :

```
GET http://localhost:3000/ --> Dashboard

GET http://localhost:3000/services/[serviceName] --> Get available instances that have name=serviceName

POST http://localhost:3000/services/[serviceName]?ip=[service ip address] --> Add new service

PUT http://localhost:3000/services/[serviceName]?ip=[service ip address] --> Change service status to 'up'

DELETE http://localhost:3000/services/[serviceName]?ip=[service ip address] --> Change service status to 'down'
```

## Run the app :

```
docker build . -t golang_container
docker-compose up
```
