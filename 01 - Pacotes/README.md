Para criar o modulo executar o comando

```powershell
go mod init modulo
```

Para executar o projeto
```powershell
go run main.go
```

Builda o projeto e cria um executavel
```powershell
go build
```

para executar o executavel 
```powershell
./<nome do ficheiro>
```

Letra maiuscula serve como um funação publica
Letra minuscula fica a função privada

função com letra minuscula ela só vai ser visivel dentro do pacote que ela está
função com letra maiuscula pode ser "exportada"