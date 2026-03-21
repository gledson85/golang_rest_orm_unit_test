
# Golang REST API with GORM and Swagger

## Descrição

API RESTful em Go para gerenciamento de usuários com operações CRUD (Create, Read, Update, Delete). Utiliza o framework Gin, persistência via GORM com MySQL em produção e SQLite in-memory para testes, e documentação interativa com Swagger.

## Tecnologias Utilizadas

- [Go 1.26+](https://golang.org/)
- [Gin](https://github.com/gin-gonic/gin) - Framework Web
- [GORM](https://gorm.io/) - ORM para Go
- [MySQL](https://www.mysql.com/) - Banco de dados (produção)
- [SQLite](https://www.sqlite.org/index.html) - Banco de dados in-memory (testes, sem CGO)
- [Swagger](https://swagger.io/) / [swaggo](https://github.com/swaggo/swag) - Documentação de API
- [Viper](https://github.com/spf13/viper) - Gerenciamento de configuração
- [Testify](https://github.com/stretchr/testify) - Assertions para testes

## Estrutura do Projeto

```
golang_rest_orm_unit_test/
├── config/
│   └── config.yaml          # Configuração do servidor e banco de dados
├── controllers/
│   └── user_controller.go   # Handlers dos endpoints (CRUD)
├── database/
│   └── database.go          # Conexão MySQL (produção) e SQLite (testes)
├── docs/
│   ├── docs.go              # Swagger docs (gerado automaticamente)
│   ├── swagger.json
│   └── swagger.yaml
├── models/
│   └── user.go              # Model User (ID, Name, Email)
├── test/
│   └── user_test.go         # Testes de integração
├── main.go                  # Entrypoint, rotas e Swagger
├── go.mod
└── go.sum
```

## Instalação e Execução

### Pré-requisitos

- Go 1.26.1 ou superior
- MySQL rodando (para ambiente de produção)

### Instalar Go

#### Windows

1. Baixe o instalador em [https://go.dev/dl](https://go.dev/dl).
2. Execute o instalador e siga as instruções.
3. Verifique a instalação:

    ```sh
    go version
    ```

#### Linux

1. Baixe e extraia o tarball:

    ```sh
    sudo tar -C /usr/local -xzf go1.26.1.linux-amd64.tar.gz
    ```

2. Adicione ao PATH:

    ```sh
    echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.profile
    source ~/.profile
    ```

3. Verifique a instalação:

    ```sh
    go version
    ```

### Clonar o Repositório

```sh
git clone https://github.com/gledson85/golang_rest_orm_unit_test.git
cd golang_rest_orm_unit_test
```

### Instalar Dependências

```sh
go mod tidy
```

### Configurar o Banco de Dados

Edite o arquivo `config/config.yaml` com suas credenciais MySQL:

```yaml
server:
  port: 8081

database:
  user: seu_usuario
  password: sua_senha
  host: 127.0.0.1
  port: 3306
  name: seu_banco_de_dados
```

### Instalar Swagger CLI

#### Windows

1. Baixe o executável do [swag releases](https://github.com/swaggo/swag/releases).
2. Adicione o diretório ao PATH do sistema.
3. Verifique:

    ```sh
    swag --version
    ```

#### Linux

```sh
go install github.com/swaggo/swag/cmd/swag@latest
```

### Gerar a Documentação Swagger

```sh
swag init
```

### Executar a Aplicação

```sh
go run main.go
```

Acesse a documentação Swagger: http://localhost:8081/swagger

## Endpoints

### Criar Usuário

- **URL**: `POST /users`
- **Status**: `201 Created`
- **Request**:

  ```sh
  curl -X POST http://localhost:8081/users \
    -H "Content-Type: application/json" \
    -d '{"name": "James Silva", "email": "james.silva@mail.com"}'
  ```

- **Response**:

  ```json
  {
    "id": 1,
    "name": "James Silva",
    "email": "james.silva@mail.com"
  }
  ```

- **Erros**: `400` JSON inválido | `500` Erro interno

### Obter Usuário por ID

- **URL**: `GET /users/{id}`
- **Status**: `200 OK`
- **Request**:

  ```sh
  curl http://localhost:8081/users/1
  ```

- **Response**:

  ```json
  {
    "id": 1,
    "name": "James Silva",
    "email": "james.silva@mail.com"
  }
  ```

- **Erros**: `404` Usuário não encontrado

### Listar Todos os Usuários

- **URL**: `GET /users`
- **Status**: `200 OK`
- **Request**:

  ```sh
  curl http://localhost:8081/users
  ```

- **Response**:

  ```json
  [
    {
      "id": 1,
      "name": "James Silva",
      "email": "james.silva@mail.com"
    },
    {
      "id": 2,
      "name": "Maria Silva",
      "email": "maria.silva@mail.com"
    }
  ]
  ```

- **Erros**: `500` Erro interno

### Atualizar Usuário

- **URL**: `PUT /users/{id}`
- **Status**: `200 OK`
- **Request**:

  ```sh
  curl -X PUT http://localhost:8081/users/1 \
    -H "Content-Type: application/json" \
    -d '{"name": "James", "email": "james@mail.com"}'
  ```

- **Response**:

  ```json
  {
    "id": 1,
    "name": "James",
    "email": "james@mail.com"
  }
  ```

- **Erros**: `400` JSON inválido | `404` Usuário não encontrado | `500` Erro interno

### Deletar Usuário

- **URL**: `DELETE /users/{id}`
- **Status**: `200 OK`
- **Request**:

  ```sh
  curl -X DELETE http://localhost:8081/users/1
  ```

- **Response**:

  ```json
  {
    "message": "User deleted"
  }
  ```

- **Erros**: `404` Usuário não encontrado | `500` Erro interno

## Testes

### Executar Testes

```sh
go test ./test/ -v
```

### Cobertura de Testes

Os testes usam SQLite in-memory, sem necessidade de MySQL. Cada teste limpa o banco antes de executar, garantindo isolamento.

| Teste | Cenário | Status esperado |
|---|---|---|
| `TestCreateUser` | Criar usuário com dados válidos | `201` |
| `TestCreateUser_InvalidJSON` | Criar com JSON inválido | `400` |
| `TestGetUser` | Buscar usuário existente por ID | `200` |
| `TestGetUser_NotFound` | Buscar ID inexistente | `404` |
| `TestGetAllUsers` | Listar com usuários cadastrados | `200` |
| `TestGetAllUsers_Empty` | Listar sem nenhum usuário | `200` (array vazio) |
| `TestUpdateUser` | Atualizar usuário existente | `200` |
| `TestUpdateUser_NotFound` | Atualizar ID inexistente | `404` |
| `TestUpdateUser_InvalidJSON` | Atualizar com JSON inválido | `400` |
| `TestDeleteUser` | Deletar usuário existente | `200` |
| `TestDeleteUser_NotFound` | Deletar ID inexistente | `404` |

## Contribuição

1. Fork este repositório.
2. Crie uma nova branch: `git checkout -b minha-nova-funcionalidade`.
3. Faça suas alterações e commit: `git commit -m 'Adicionar nova funcionalidade'`.
4. Envie para a branch original: `git push origin minha-nova-funcionalidade`.
5. Crie uma pull request.

## Licença

Este projeto está licenciado sob a Licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.
