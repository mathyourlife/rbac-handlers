# rbac-handlers

Setup some simple http handler endpoints for RBAC calls, and allow for
varied backends.

## Integration Testing

```bash
cd integration
docker-compose up -d --build
API_PORT=$(docker-compose port api 8443 | cut -d ':' -f 2)
curl -v -k "https://localhost:${API_PORT}/auth/login"
```