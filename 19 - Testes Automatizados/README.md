### Executar testes
```
go test

go test -v
```

### Gerar cobertura dos testes
```
go test ./... --coverprofile coverage.txt
```

```
go tool cover --html=coverage.txt
```