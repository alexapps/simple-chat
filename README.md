# Chat example

## The back folder consist of
1. api service that handles http and ws communication with the Web
2. auth service responsible for the oAuth 
3. chat service - the chat itself.

## The fron folder will consist the Web part. It could be the Angular/React implementation

### The things to be done:
#### common:
1. microservice framework (all microservices routine, e.g. communicating)
2. docker file per each service
3. docker-compose for whole back solution (with 3rd party, e.g. postgres)

#### api:
1. auth handlers for retrieve singUp/singIn data
2. chat handlers: post message, get message history per each chat

#### auth:
1. implenet oauth logic (e.g. password caching, token generation, etc)
2. sorage functionality (CRUD)

#### chat:
1. handler logic
2. storage (CRUD)

