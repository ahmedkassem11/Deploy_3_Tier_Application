# Three-Tier Blog API (Go + MySQL + Nginx)

This repository contains a complete three-tier blog API stack:
- **Backend**: Go REST API serving blog titles
- **Database**: MySQL with persistent storage and initial seed data
- **Proxy**: Nginx reverse proxy exposing HTTPS
- **Local deployment**: Docker Compose
- **Production-like deployment**: Kubernetes manifests

## Repository structure
```
.
├── backend/
│   ├── main.go
│   └── Dockerfile
├── mysql/
│   └── init.sql
├── proxy/
│   ├── nginx.conf
│   └── generate-ssl.sh
├── docker-compose.yml
├── k8s/
│   ├── db-secret.yaml
│   ├── db-data-pv.yaml
│   ├── db-data-pvc.yaml
│   ├── database_deployment.yaml
│   ├── db-service.yaml
│   ├── backend_deployment.yaml
│   ├── backend_service.yaml
│   ├── proxy_deployment.yaml
│   └── proxy_nodeport.yaml
└── architecture.svg
```

## Quickstart — Local (Docker Compose)
1. Generate SSL certs for the proxy:
   ```bash
   cd proxy
   chmod +x generate-ssl.sh
   ./generate-ssl.sh
   cd ..
   ```
2. Build and run:
   ```bash
   docker compose up --build
   ```
3. Visit `https://localhost` (accept self-signed certificate) and call:
   ```
   https://localhost/blogs
   ```

## Quickstart — Kubernetes (minikube / k3s)
1. Apply secrets and PV/PVC:
   ```bash
   kubectl apply -f k8s/db-secret.yaml
   kubectl apply -f k8s/db-data-pv.yaml
   kubectl apply -f k8s/db-data-pvc.yaml
   ```
2. Deploy DB, backend, and proxy:
   ```bash
   kubectl apply -f k8s/database_deployment.yaml
   kubectl apply -f k8s/db-service.yaml
   kubectl apply -f k8s/backend_deployment.yaml
   kubectl apply -f k8s/backend_service.yaml
   kubectl apply -f k8s/proxy_deployment.yaml
   kubectl apply -f k8s/proxy_nodeport.yaml
   ```
3. Access via node port (example `https://<NODE_IP>:30443`)

## Notes
- Replace `yourdockerhub/...` images in `k8s/` manifests with images you build & push.
- Secrets: example uses `mypassword` (base64 `bXlwYXNzd29yZA==`) — change for production.
- This repo is a starting point; see /backend and /mysql for example implementation.

## License
MIT
