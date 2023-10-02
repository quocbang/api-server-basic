# Overview API SERVER BACSIC project

[![Go Report Card](https://goreportcard.com/badge/github.com/quocbang/api-server-basic?style=flat-square)](https://goreportcard.com/report/github.com/quocbang/api-server-basic)

This is personal project using Echo framework to build API

- built_by: @quocbang
- start_built: Jun 21, 2023
- end_built:

## ***DEMO***

To easy for you are demo this project, I have built all of them with images file(using docker) you just need to do a few operations to have all of my project features

Follow this link and do step to step for demo

[Run demo](/docker-compose/)

## Process Flow Diagram

![process image](/image/project-flow-diagram.png)

## Architecture Diagram

![architecture diagram](/image/architecture_diagram.png)

## Project Structure

Main structure

- [api](/api/)
- [cmd](/cmd/)
- [config](/config/)
- [database](/database/)
- [docker-compose](/docker-compose/)
- [errors](/errors/)
- [image](/image/)
- [impl](/impl/)
- [middleware](/middleware/)
- [utils](/utils/)

Detail structure

```bash
api-server-basic/

├── api/
│   │── tasks/
│   │   └── tasks.go     
│   │── users/
│   │   └── users.go
│   └── api.go
├── cmd/
│   │── mockgenerate/
│   │   ├── main.go
│   │   └── mock_generate.go
│   └── cmd.go
├── config/
│   └── config.go
├── database/
│   ├── impl/
│   │   ├── impl.go
│   │   └── service.go
│   ├── orm/
│   │   └── models/
│   │       ├── models.go
│   │       ├── tasks.go
│   │       └── users.go
│   ├── services/
│   │   │── tasks/
│   │   │   └── task.go
│   │   └── users/
│   │       └── user.go
│   ├── servicetest/
│   │   │── internal/
│   │   │   └── suite/
│   │   │       ├── build.go
│   │   │       └── setup.go
│   │   │── tasks/
│   │   │   ├── task_test.go
│   │   │   └── suite_test.go
│   │   │── users/
│   │   │   ├── user_test.go
│   │   │   └── suite_test.go
│   ├── utils/
│   │   │── gorm/
│   │   │   └── postgres/
│   │   └       └── errors.go
│   ├── dm.go
│   └── dmlist.go
├── docker-compose/
│   ├── config/
│   │   └── config.go
│   ├── .env
│   ├── .gitignore
│   ├── docker-compose.yml
│   └── README.md
├── errors/
│   ├── code.pb.go
│   ├── code.proto
│   ├── error.go
│   └── go_generate.go
├── image/
│   └── architecture_diagram.png
├── impl/
│   ├── mock/
│   │   └── some_mock_file.go
│   ├── requests/
│   │   │── common_replies.go
│   │   │── request.go
│   │   └── task.go
│   ├── service/
│   │   │── tasks/
│   │   │   └── task.go
│   │   │── users/
│   │   │   └── user.go
│   │   └── service.go
│   ├── servicetest/
│   │   │── internal/
│   │   │   │── buildtest/
│   │   │   │   └── moskv_test.go
│   │   │   │── suite/
│   │   │   │   └── setup.go
│   │   │── tasks/
│   │   │   ├── task_test.go
│   │   │   └── suite_test.go
│   │   │── users/
│   │   │   ├── user_test.go
│   │   │   └── suite_test.go
│   ├── utils/
│   ├── register.go
│   └── service_func.go
├── middleware/
│   ├── authorization/
│   │   └── authorized.go
│   ├── context/
│   │   └── context.go
│   ├── logging/
│   │   └── logger.go
├── utils/
│   ├── function/
│   │   ├── func_generate.pb.go
│   │   ├── func_generate.proto
│   │   └── go_generate.go
│   ├── roles/
│   │   ├── go_generate.go
│   │   ├── parm.go
│   │   ├── role.pb.go.proto
│   │   └── role.proto
├── .gitignore
├── config.yaml(this file is config file should ignore)
├── Dockerfile
├── go.mod
├── go.sum
├── main.go
└── README.md
```

## Run server

I have to define several flags to run server as below:

Without TLS:

`--dev-mode` : to seting status when run project (use this flag with debug not for production)
`--host`: server IP address
`--port`: server port
`--config`: server config file

Once combine looks like this:

```bash
  go run . --dev-mode --host=localhost --port=your_port --config=your_config_path
```

with TLS:

`--dev-mode` : to seting status when run project (use this flag with debug not for production)
`--host`: server IP address
`--port`: server port
`--tls-cert`: TLS cert file
`--tls-key`: TLS key file
`--config`: server config file

Once combine looks like this:

```bash
  go run . --dev-mode --host=your_host --port=your_port --tls-cert=your_cert_file --tls-key=your_key_file --config=your_config_path
```

## The Tecnique This Project Use

### Web Frameworks

- Echo

### Database

- Relational
  - Postgres

- NoSQL
  - Redis

### ORM

- Gorm

### Logging

- zap

### Validator

- github.com/go-pldayground/validator(struct)

### Testing

- github.com/stretchr/testify/suite
- github.com/stretchr/testify/mock

## Config File

```bash
postgres:
  address: localhost
  port: 5432
  name: test
  schema: test
  username: db_user_name
  password: db_user_password
roles:
  YOUR_FUNCTION:
    - Roles1
    - Roles2
    - Roles3
```
