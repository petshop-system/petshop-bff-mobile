# build stage
FROM golang:1.22.3 AS builder
# working directory
WORKDIR /app
COPY ./ /app

RUN go get -d /app/cmd/petshop-bff-mobile

# rebuilt built in libraries and disabled cgo
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app /app/cmd/petshop-bff-mobile
# final stage
FROM alpine:3.19.1

# working directory
WORKDIR /app

# copy the binary file into working directory
COPY --from=builder /app .
# http server listens on port 9997
EXPOSE 9997
# Run the docker_imgs command when the container starts.
CMD ["/app/petshop-bff-mobile"]