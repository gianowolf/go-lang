# Building a Service

## ListenAndServe

Starts an HTTP server with:

- given address
- a handler: if handler is nill, usually in net/http lib, DefaultServeMux value is used

### Handler

A Handler is any type that satisfies the Handler interface by providing a ServeHTTP Method

- type Handler interface { ServeHTTP(ResponseWriter, *Request)}

### Mux

Handler acts as a multiplexer. The mux compares the requested URL to the registered patterns and call the handler function
DefaultServeMux is a global value of type serveMux which implements default HTTP multiplexer logic

## Gorilla Mux

1. Create an empty directory
2. Initialize the project using: go mod init example.com/gorilla
3. Add dependencies: go mod tidy
4. cat go.mod to see dewpendencies

If change dependencies, run 'go mod tidy' again to rebuild the go.mod file

### Variables in URI

create paths with variable segments

## Restul Service

