# syntax=docker/dockerfile:1
FROM golang:1.26 AS build
WORKDIR /src
COPY go.work ./
COPY libs ./libs
COPY services ./services
WORKDIR /src/services/ticket-query
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/app ./cmd

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /out/app /app
ENTRYPOINT ["/app"]
