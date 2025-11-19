# go-testing-course
Project to learn unit testing in Golang

## Coverage
* Ver coverage en archivo de texto:
```bash
go test ./controller -coverprofile=coverage.out
```

* Ver desde navegador el archivo .out:
```bash
go tool cover -html=coverage.out
```
