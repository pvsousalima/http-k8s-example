FROM golang:alpine as build
WORKDIR /app
COPY main.go .
RUN CGO_ENABLED=0 go build -o main

FROM scratch
COPY --from=build /app/main .
EXPOSE 8080
ENTRYPOINT ["./main"]
