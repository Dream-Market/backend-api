FROM golang:1.19
COPY . /api-gateway
WORKDIR /api-gateway
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN apt update --no-install-recommends
RUN apt install --no-install-recommends -y protobuf-compiler 
RUN make proto
RUN go mod download
RUN go build -o api-gateway cmd/main.go
CMD ./api-gateway
