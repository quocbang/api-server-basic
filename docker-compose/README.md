# Docker Compose

Here is the run docker-compose area, if you run the given command line, it means you building all features of this repository.

To demo this repo you need to do step to step as below:

Clone the repository:

```bash
    git clone https://github.com/quocbang/api-server-basic.git
```

Move to docker-compose:

```bash
    cd docker-compose
```

Run docker-compose to pull images and build project:

```bash
    docker-compose down && docker-compose up --build -d;
```

***Note***: ensure your machine has been installed docker-compose then the docker can be run on this command line

If your machine hasn't installed docker then you can follow these steps to install.
Windows:

[install docker desktop for windows](https://desktop.docker.com/win/main/amd64/Docker%20Desktop%20Installer.exe?utm_source=docker&utm_medium=webreferral&utm_campaign=dd-smartbutton&utm_location=module&_gl=1*1bwjmf3*_ga*NDg3NTExNTEyLjE2ODgxNzczMDA.*_ga_XJWPQMJYHQ*MTY4ODUyOTE1OC41LjEuMTY4ODUyOTIzNi40Ni4wLjA.)

Linux or Ubuntu:

[follow this link](https://docs.docker.com/desktop/install/linux-install/)

## Config File

Config file structure:

- postgres
  - address: `postgres_IP`
  - port: `postgres_PORT`
  - name: `postgres_Name`
  - schema: `postgres_Schema`
  - username: `postgres_UserName`
  - password: `postgres_Password`
- redis
  - address: localhost:6379
  - password: api_server_basic_password
  - database: 1
- demo-account:
  - because creating a user API post `(http://localhost:8810/api/user)` only support for ADMINISTRATOR roles, therefore in the demo part after using `docker-compose down && docker-compose up --build -d;` then the demo-account will be created during server build
  - login with the demo-account to get token.
    - demo login with demo account
    - ![login with demo account](/image/postman-demo-login.png)
  - you will be given 2 demo accounts or more than dependencies in your config file
    - Email: `admin_01@gmail.com` Password: `admin_01_password`
    - Email: `admin_02@gmail.com` Password: `admin_02_password`
- roles:
  - each function should have its own roles
  - example:
    - Function_A only for ADMINISTRATOR can be access, but Function_B can be access by both ADMINISTRATOR and USER we will difine like below:

       ```bash
            roles:
                Function_A:
                    - ADMINISTRATOR
                Function_B
                    - ADMINISTRATOR
                    - USER
       ```

```bash
postgres:
  address: 172.20.0.3
  port: 5432
  name: test
  schema: public  
  username: api_server_basic_user  
  password: api_server_basic_password
redis:
  address: localhost:6379
  password: api_server_basic_password
  database: 1
demo-account: # admin account define here just support for DEMO purpose.
  - admin_01@gmail.com|admin_01_password
  - admin_02@gmail.com|admin_02_password
roles:
  CREATE_ACCOUNT:
    - ADMINISTRATOR
  DELETE_ACCOUNT:
    - ADMINISTRATOR    

  GET_TASKS:
    - ADMINISTRATOR
    - LEADER
    - USER 
  CREATE_TASKS:
    - ADMINISTRATOR
    - LEADER
  UPDATE_TASKS:
    - ADMINISTRATOR
    - LEADER
  DELETE_TASKS:
    - ADMINISTRATOR
    - LEADER
```

## APIs

After completing those steps you can use the Postman to call these APIs as below:

With `USER` service

- Login: see user-name and password in [config/config.yaml](/docker-compose/config/config.yaml)

  ```bash
  POST http://localhost:8810/api/user/login
  request body
      {
          "email": "example@gmail.com",
          "password": "your_password"
      }
  response token 
      {
          "token": "xxxxxxx.xxxxxxxxx.xxxxxxx"
      }
  ```

- Logout: sign-out and save token to redis.

  ```bash
  POST http://localhost:8810/api/user/logout
  authorization
      key = x-server-auth-key
      values = xxxxxxx.xxxxxxxxx.xxxxxxx
  response token 
      "Logout successfully"
  ```

- Create User

  ```bash
  POST http://localhost:8810/api/user
  authorization
      key = x-server-auth-key
      values = xxxxxxx.xxxxxxxxx.xxxxxxx // should is administrator token
  request body
      {
          "email": "user@gmail.com",
          "password": "user_password",
          "roles": [
              0,
              1,
              2 
          ]    
      }
  response
      {
          "RowsAffected": n(row)
      }
  ```

- Delete User

  ```bash
  DELETE http://localhost:8810/api/user
  authorization
      key = x-server-auth-key
      values = xxxxxxx.xxxxxxxxx.xxxxxxx // should is administrator token
  request body
      {
          "emails": [
              "user_01@gmail.com",
              "user_02@gmail.com"
          ]
      }
  response
      {
          "RowsAffected": n(row)
      }
  ```

With `TASKS` service

- Create Tasks:

    ```bash
    POST http://localhost:8810/api/task
    authorization
        key = x-server-auth-key
        values = xxxxxxx.xxxxxxxxx.xxxxxxx
    request body
        {
            "tasks": [
                {
                    "id": 1,
                    "status": "doing",
                    "process_percent": 10, 
                    "content": "create first task today and do now",
                    "end_time": 8 // task time = end_time * time.Hour
                }
            ]
        }
    response
        {
            "RowsAffected": n(row)
        }
    ```

- List Tasks:

    ```bash
    GET http://localhost:8810/api/task
    authorization
        key = x-server-auth-key
        values = xxxxxxx.xxxxxxxxx.xxxxxxx
    request body
        {
            "id": 100 // request without id = get all
        }
    response
        {
        "Tasks": [
                {
                    "ID": 100,
                    "Status": "doing",
                    "ProcessPercent": 10,
                    "Content": "create first task today and do now",
                    "EndTime": "2023-07-05T00:35:37.676863+07:00",
                    "Updated": 1688463337,
                    "Created": 1688463337
                }
            ]
        }
    ```

- Update Task:

    ```bash
    PATCH http://localhost:8810/api/task
    authorization
        key = x-server-auth-key
        values = xxxxxxx.xxxxxxxxx.xxxxxxx
    request body
        {
            "id": 100,
            "status": "complete",
            "process_percent": 100
        }
    response
        {
            "RowsAffected": 1
        }
    ```

- Delete Tasks:

    ```bash
    DELETE http://localhost:8810/api/task
    authorization
        key = x-server-auth-key
        values = xxxxxxx.xxxxxxxxx.xxxxxxx
    request body
        {
            "ids": [1,2,3,4]
        }
    response
        {
            "RowsAffected": 1
        }  
    ```
