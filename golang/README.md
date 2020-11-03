# Desafios Backend - Bcredi
### Executando o Projeto

### Testes
*simple*

    go test ./...
*verbose*

    go test ./... -v

*coverage html*

    go test ./... --coverprofile=coverage.out
    go tool cover --html=coverage.out

### Run
    go run  ./cmd/main.go ../test/input/input000.txt
