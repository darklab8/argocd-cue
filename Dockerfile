FROM golang:1.21-bullseye as build

# valid values arm64 or amd64
ARG ARCH

RUN apt update
RUN apt install -y build-essential
RUN apt-get install ca-certificates -y
RUN gcc --version

WORKDIR /code

COPY go.mod go.sum ./
RUN go mod download -x

COPY main.go ./
COPY argocue argocue
RUN CGO_ENABLED=0 go build -v -o main main.go

FROM alpine:3.18 as runner
RUN apk add --no-cache wget
WORKDIR /code
COPY --from=build --chmod=555 /code/main argocue
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ./argocue run
