# Custom-K8-Pod-Scaler
Develop a small-scale application that dynamically scales Kubernetes pods based on custom metrics, such as HTTP request rates or message queue lengths.

Docker Steps:
docker network create my-network
docker build -t my-go-app .
docker build -t load-generator . 
docker run -d --name my-go-app --network my-network my-go-app
docker run -d --name load-generator --network my-network load-generator
docker run -d --name prometheus --network my-network -p 9090:9090 \
  -v /Users/alexcholmsky/Custom-K8-Pod-Scaler/Consumer/Prometheus/prometheus.yml:/etc/prometheus/prometheus.yml  \
  prom/prometheus
