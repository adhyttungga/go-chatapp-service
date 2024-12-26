# GO-CHATAPP-SERVICE

###### version: v1.0

###### Last Update: 2024-12-26

### Programming Language

- [go1.23 windows/amd64]

### Framework

- [GIN]

### How to Deploy on Localhost

```
go run .
```

### How to Test on Localhost or Development Server

```
go test github.com/adhyttungga/go-chatapp-service/test/
```

### Environment Variables (Development)

Merupakan suatu variabel yang ditetapkan diluar program melalui fungsionalitas yang dibangun ke dalam OS atau Microservices. Tim DevOps perlu mengetahui Nama Env. Variables, Value Env. Variables, Description dari Env. Variables, dan Locationnya.

| Name         | Description         |
| ------------ | ------------------- |
| GIN_MODE     | Gin mode            |
| MONGO_DB_URI | Database URL        |
| SERVICE_HOST | Service's host      |
| SERVICE_PORT | Service's port      |
| ALLOW_ORIGIN | Allow origin        |
| PRIVATE_KEY  | Private key for jwt |
| PUBLIC_KEY   | Public key for jwt  |
