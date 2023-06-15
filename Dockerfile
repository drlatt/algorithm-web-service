FROM golang as builder

# RUN mkdir /build
WORKDIR /build
ADD . /build/
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o algorithm_web_service

FROM golang:1-alpine3.18
WORKDIR /app
COPY --from=builder /build/algorithm_web_service  .

EXPOSE 3000

CMD ["./algorithm_web_service"]