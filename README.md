# go-service-registry

## Availabe endpoints :

```
GET http://localhost:8671/ --> Dashboard

GET http://localhost:8671/services/[serviceName] --> Get available instances that have name=serviceName

POST http://localhost:8671/services/[serviceName]?ip=[service ip address]&port=[service port] --> Add new service

PUT http://localhost:8671/services/[serviceName]?ip=[service ip address]&port=[service port] --> Change service status to 'up'

DELETE http://localhost:8671/services/[serviceName]?ip=[service ip address]&port=[service port] --> Change service status to 'down'
```

## Run the app :

```
docker build . -t service_registry
docker run -p 8671:8671 -t service_registry
```
