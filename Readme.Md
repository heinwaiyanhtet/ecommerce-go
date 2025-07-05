# 🧪 User CRUD API Testing (cURL)

This guide provides sample `curl` commands for testing the **Create, Read, Update, and Delete** operations on the `/users` endpoint of your Go backend.

---

## ✅ Create User (POST)

```bash

curl -X POST http://localhost:8080/signup \     
  -H "Content-Type: application/json" \
  -d '{"name":"heinwaiyanhtet", "password":"12345678"}'


curl -X GET http://localhost:8080/users


curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"name":"heinwaiyanhtet2", "password":"12345678"}'


# {"exp":"2025-05-17T20:37:39+06:30","token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDc0OTA4NTksInN1YiI6N30.0SIhXTtzTnZKthr2YfVsA3I1t3JyNMK7BQ7QrQAjBpA"}



curl -X POST http://localhost:8080/orders \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "1",
    "amount": 99.99,
    "status": "pending"
  }'



---

## 🚢 Docker Usage

Build the container image:

```bash


docker buildx build --platform linux/amd64 -t registry.digitalocean.com/ecommerce-go/go-hello:latest --push .


```

Run the image locally:

```bash
docker run --env-file .env -p 8080:8080 registry.digitalocean.com/your-repo/ecommerce-go:latest
```

## ☸️ Deploy to Kubernetes

Apply the manifests in the `k8s/` directory:

```bash
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
```

Ensure your environment variables are set using a `Secret` or `ConfigMap` when deploying to DigitalOcean Kubernetes.

