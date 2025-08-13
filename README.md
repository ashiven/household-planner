# household-planner

Receive reminders of the household tasks in your responsibility

Note: If you want to host this application, replace `http://localhost/api` with `https://<domain-name>/api` in the `docker-compose.yml` file.
If you have a domain name, specify it in the nginx configuration file instead of `localhost`.

Also generate a certificate as follows:

```bash
apt install -y certbot python3-certbot-nginx
```

```bash
certbot certonly --standalone -d domain-name
```

Furthermore, replace `var allowedOrigin = "http://localhost"` in `pkg/backend/server.go` with `var allowedOrigin = "https://<domain-name>"`.
