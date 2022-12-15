# autobooking
a auto booking golang script.
1. Install Go 1.19 (other version is also ok, remember to change golang verion in go.mod)

2. execute
```
go mod tidy
```
to download dependencies

3. change config
config.json.default -> config.json and add your configuration in it.

4. development mode
Use F12 enter development mode, copy url as cURL, then copy cURL to https://mholt.github.io/curl-to-go/ and get go code.

Copy it to `query.go` and fill the body field.


5. run command:
```
go run main.go
```
