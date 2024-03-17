FROM golang:1.21-bullseye as build


RUN apt update
RUN apt install -y build-essential
RUN apt-get install ca-certificates -y
RUN gcc --version

WORKDIR /code

COPY go.mod go.sum ./
RUN go mod download -x

COPY main.go ./
RUN CGO_ENABLED=0 go build -v -o main main.go

FROM alpine:3.18 as runner
WORKDIR /code
COPY --from=build /code/main main
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ./main run
