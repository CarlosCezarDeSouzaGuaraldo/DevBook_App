# FROM golang:1.19.3 AS build
FROM golang:1.19.3

WORKDIR /app

COPY . .

RUN go build -o /server

# FROM gcr.io/distroless/base-debian10

# WORKDIR /

# COPY --from=build /server /server

EXPOSE 5000

# USER nonroot:nonroot

ENTRYPOINT [ "/server" ]