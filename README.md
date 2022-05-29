## Go Rest DDD V2

### Run App With Docker
```sh
    git clone git@github.com:geekbim/Go-Rest-DDD-V2.git
    cd Go-Rest-DDD-V2
    docker-compose up
```

### Run App Without Docker
```sh
    git clone git@github.com:geekbim/Go-Rest-DDD-V2.git
    cd Go-Rest-DDD-V2
    cp .env.example .env
    go mod tidy
    sh run-service.sh
```

### Run Migration
```sh
    sh run-migration.sh
```

### Run Test
```sh
    sh run-test.sh
```

### Demo Account
```sh
    SELLER : 
    - email: seller@email.com
    - password: qweasd123

    BUYER : 
    - email: buyer@email.com
    - password: qweasd123
```

### Docs Swagger
```sh
    http://localhost:8080/doc/index.html
```

### Docs postman
```sh
    https://documenter.getpostman.com/view/1850032/Uz5CLJ9Q
```
