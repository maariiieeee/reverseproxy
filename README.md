**How can someone build and run your code?**
1) edit the config.env file, with your desired configuration for the reverse proxy (proxy adress, backend adress, max concurrent requests and max retries)
2) run `go get` command in the root directory to get the required dependencies
3) run `go run reverse_proxy.go` to get the proxy software started
4) connect to the specified proxy address in `config.env` via a browser or postman to test and use the reverse proxy
