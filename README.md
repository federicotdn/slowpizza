# SlowPizza 🍕
Like [QuickPizza](https://github.com/grafana/quickpizza) but with gRPC. SlowPizza is a small service to test or debug gRPC connections, locally or in the cloud. Unlike what its name implies, it actually runs quickly.

Based on some of the examples in the [gRPC-Go](https://github.com/grpc/grpc-go/tree/master/examples/features) repository.

This repository is split into two parts: client and server.

## Server

To deploy the SlowPizza server, use the `federicotedin/slowpizza` Docker image. The following environment variables can be used to configure it:
- `SLOWPIZZA_PORT`: Set gRPC server port. Default: 50051.
- `SLOWPIZZA_AUTH_TOKEN`: Value of authorization token required by the sever. Setting this to `disabled` will disable it. If not set, all connections will be rejected.

### Using docker-compose

```yaml
services:
  slowpizza:
    image: federicotedin/slowpizza:latest
    environment:
      - SLOWPIZZA_AUTH_TOKEN=my-auth-token
```

### Using Kubernetes

Service + Deployment: see the [kubernetes.yaml](etc/kubernetes.yaml) file.

## Client

Run `make-client` to build the client binary, then run it with `./slowpizza-client`.

Client usage:
```
Usage of slowpizza-client:
  -H value
    	Extra headers ('Key: Value')
  -addr string
    	The address to connect to (default "localhost:50051")
  -cert string
    	Path to PEM-encoded server certificate (.crt)
  -item value
    	Item to order
  -plaintext
    	Disable TLS
  -token string
    	Authorization token
  -unsafetoken
    	Allow sending authorization token over plaintext
```
