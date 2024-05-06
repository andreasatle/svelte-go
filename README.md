# Svelte-Go Webpage

To build the frontend:

```bash
cd svelte
npm run build
cd ..
```

To run the backend (after the frontend is built):

```bash
cd go
go run .
```

# Create fake certificates for TLS (in the go-directory)

```bash
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout key.pem -out cert.pem
```
