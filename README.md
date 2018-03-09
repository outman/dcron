### DCron

Web crontab based on Golang. (Just for study, don't use it on production env.)

### Desp

govendor list

### Config
- conf/conf.yaml  db & hosts

### Usage

- go build
- ./dcron -mode=server for web
- ./dcron -mode=cli  for crontab

```go
func main() {

    // server for web, cli for crontab
	mode := flag.String("mode", "server", "-mode=server|cli")

    // web server host & port
	port := flag.String("port", "8080", "-port=8080")
	host := flag.String("host", "", "-host=")
	flag.Parse()

	if bytes.EqualFold([]byte(*mode), []byte("cli")) {
		fmt.Println("Starting...cli mode...")
		app.RunCrond()
	} else {
		fmt.Println("Starting...server mode...")
		app.RunServer(*host, *port)
	}
}
```

