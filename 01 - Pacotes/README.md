# Modulos

## Modulos internos

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

EX
```powershell
./modulo
```

Letra maiuscula serve como um funação publica
Letra minuscula fica a função privada

função com letra minuscula ela só vai ser visivel dentro do pacote que ela está
função com letra maiuscula pode ser "exportada"

## Modulos externos

Para baixar pacotes externos executar
```powershell
go get <diretorio do pacote>
```

EX
```powershell
go get github.com/badoux/checkmail
```

Para remover todas as dependencias não usadas
```powershell
go mod tidy
```