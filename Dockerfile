FROM golang:1.14.4-alpine AS compiler

WORKDIR /app/src

COPY . .

RUN go build -o qrmenum_client

FROM alpine

COPY --from=compiler /app/src/qrmenum_client /qrmenum_client
COPY --from=compiler /app/src/templates /templates

ENTRYPOINT ["/qrmenum_client"]
