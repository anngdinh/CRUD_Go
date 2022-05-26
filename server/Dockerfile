# FROM golang:latest

# RUN mkdir /app
# WORKDIR /app

# RUN go mod install



# builder image
FROM golang:latest as builder
WORKDIR /build
COPY . .
RUN go build -o main

# generate clean, final image for end users
FROM ubuntu:18.04
WORKDIR /app
COPY --from=builder /build/main .
# executable
ENTRYPOINT [ "/main" ]