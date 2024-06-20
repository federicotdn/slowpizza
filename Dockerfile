FROM golang:1.22

WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o slowpizza-server ./server

# ENV GRPC_GO_LOG_VERBOSITY_LEVEL=99
# ENV GRPC_GO_LOG_SEVERITY_LEVEL=info
ENTRYPOINT ["/build/slowpizza-server"]