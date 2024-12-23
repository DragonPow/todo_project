# Architecture

## Todo service

### Summary
- Description: management `todo` and `item`
- Config file: `app/todo/config.yaml`
- Port:
  - gRPC: `443`
  - Restful: `8080`
- Database:
  - Port: `5433`
  - User: `postgres`
  - Password: `123456`
  - Database: `todo`

### API
- See full API in: `app/todo/api/api.proto`

Create todo:
```bash
curl --location 'http://localhost:8080/api/v1/todos' \
--header 'Content-Type: application/json' \
--data '{
    "todo": {
        "title": "ticket",
        "description": "ok",
        "completed": false,
        "items": [
            {
                "title": "item 1",
                "completed": false
            },
            {
                "title": "item 2",
                "completed": false
            }
        ]
    }
}'
```

Update todo:
```bash
curl --location --request PUT 'http://localhost:8080/api/v1/todos/{todoId}' \
--header 'Content-Type: application/json' \
--data '{
    "title": "my title",
    "description": "new description",
    "completed": true
}'
```

Get List todo:
```bash
curl --location 'http://localhost:8080/api/v1/todos?like_title=t'
```

Get Detail todo:
```bash
curl --location 'http://localhost:8080/api/v1/todos/{todoId}'
```

Delete todo:
```bash
curl --location --request DELETE 'http://localhost:8080/api/v1/todos/{todoId}'
```

## Iam service

### Summary
- Description: management user session login
- Config file: `app/iam/config.yaml`
- Port:
  - gRPC: `444`
  - Restful: `8081`
- Database:
  - Port: `5434`
  - User: `postgres`
  - Password: `123456`
  - Database: `iam`

### API
Updating...


# Starting services

Start an service:
```bash
cd path/to/service
# Example: cd app/todo

# Starting service
make run
```

Start all services:
```bash
make run-all
```

When start service in the first time, need to init database
```bash
# Init database todo
cd app/todo
make migrate-up

# Init database iam
cd app/iam
make migrate-up
```

# How to develop ?

## Implement more API

- Add new API into proto file `path/to/service/api/*.proto`
- Then run command `make proto` in directory of service
- Implement new API features in `service` and `repository` module

## Change database structure
Some query such as (`ALTER`, `INSERT TABLE`)

- Add new file with format `YYYYHHddHHhhmm_xxx.up.sql` in dirtory `/path/to/service/migrations`
- Then run command `make migrate` in directory of service

## Change config

1. Change config in docker file `docker-compose.yml`
   - Config of every services will in `enviroment` file
2. Change config in `/path/to/service/cofig.yaml`
   - Config when load from module `config` will be override by this field

Concept load config:
- Load default config in module `config`
- If exists file config, new value will load and override from file `config.yaml`
- If exists env variable, new value will load and overrid from file `docker` (Dockerfile or docker-compose)