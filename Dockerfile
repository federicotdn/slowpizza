FROM golang:1.22

WORKDIR /build

# Nicer Docker image build caching
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY server server
COPY slowpizza slowpizza

RUN go build -o slowpizza-server ./server

# ENV GRPC_GO_LOG_VERBOSITY_LEVEL=99
# ENV GRPC_GO_LOG_SEVERITY_LEVEL=info
ENTRYPOINT ["/build/slowpizza-server"]
