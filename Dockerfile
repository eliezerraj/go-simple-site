# docker build -t go_simple_site_bike .
# docker run -dit --name go_simple_site_bike -p 3000:3000 go_simple_site_bike

# docker build -t go_simple_site_pay .
# docker run -dit --name go_simple_site_pay -p 3001:3000 go_simple_site_pay

# docker build -t go_simple_site_food .
# docker run -dit --name go_simple_site_food -p 3002:3000 go_simple_site_food

# docker build -t go_simple_site_stream .
# docker run -dit --name go_simple_site_stream -p 3003:3000 go_simple_site_stream

FROM golang:1.18 As builder

RUN apt-get update && apt-get install bash && apt-get install -y --no-install-recommends ca-certificates

WORKDIR /app

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-w' -o go_simple_site

#FROM scratch
FROM alpine

WORKDIR /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/go_simple_site .
COPY --from=builder /app/html/*.html /app/html/
COPY --from=builder /app/img/*.jpg  /app/img/

ENV APP_MODE=stream

CMD ["/app/go_simple_site"]