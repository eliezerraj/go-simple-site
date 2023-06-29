# go-simple-site

Ingress Controller POC for technical purposes

## How to use

1-) Create a image for each endpoint

    docker build -t go_simple_site_bike .
    docker build -t go_simple_site_pay .
    docker build -t go_simple_site_food .
    docker build -t go_simple_site_stream .

2-) Deploy the deployments/services/ingress

    ./kubernetes

3-) Call the endpoints

    http://localhost/bike
    http://localhost/food
    http://localhost/pay
    http://localhost/stream


## Endpoint

GET /pay

GET /food

GET /bike

GET /stream

