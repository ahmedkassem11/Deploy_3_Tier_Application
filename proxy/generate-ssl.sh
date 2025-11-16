#!/bin/bash
mkdir -p certs
openssl req -x509 -newkey rsa:4096 -sha256 -nodes \
  -keyout certs/server.key -out certs/server.crt \
  -subj "/CN=localhost" -days 365
