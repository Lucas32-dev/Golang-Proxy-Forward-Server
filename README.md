
# Golang Proxy Forward 

Just an simple proxy forward server.
 

## Instalation

Be aware to have golang PATH system variable configurated.


```bash
  cd my-project
  go mod tidy // Install dependencies.
  go run main.go // Run server
```

## API Documentation

#### Dog picture

```http
  GET /dog
```
#### Json object

```http
  GET /jsonplaceholder
```

## Reference

 - [Fiber API](https://docs.gofiber.io/api/fiber)
 - [Fiber Proxy Middleware](https://docs.gofiber.io/api/middleware/proxy)


## Author

- [@Lucas32-dev](https://github.com/Lucas32-dev)

