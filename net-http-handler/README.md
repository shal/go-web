# Handler

> Client Server architecture

Entrypoint to `net/http` standart library.

```go
type Handler interface {
  ServeHTTP(ResponseWriter, *Request)
}
```
