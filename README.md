
# Golang REST API with GORM and Swagger

## Descrição

Este projeto demonstra a implementação de uma API RESTful em Go utilizando o framework Gin, com persistência de dados via GORM e documentação da API com Swagger. A aplicação permite operações CRUD (Create, Read, Update, Delete) em usuários, armazenando os dados em um banco de dados SQLite em memória para testes e em MySQL para produção.

## Funcionalidades

- **Criar Usuário**: Adiciona um novo usuário ao banco de dados.
- **Obter Usuário**: Recupera um usuário específico pelo ID.
- **Obter Todos os Usuários**: Recupera todos os usuários do banco de dados.
- **Atualizar Usuário**: Atualiza as informações de um usuário existente.
- **Deletar Usuário**: Remove um usuário do banco de dados.

## Tecnologias Utilizadas

- [Go](https://golang.org/)
- [Gin](https://github.com/gin-gonic/gin) - Framework Web
- [GORM](https://gorm.io/) - ORM para Go
- [Swagger](https://swagger.io/) - Documentação de API
- [Viper](https://github.com/spf13/viper) - Gerenciamento de configuração
- [SQLite](https://www.sqlite.org/index.html) - Banco de dados em memória para testes (CGO support 0 - false)
- [MySQL](https://www.mysql.com/) - Banco de dados para produção

## Estrutura do Projeto

```
golang_rest_orm_unit_test/
├── config/
│   └── config.yaml
├── controllers/
│   └── user_controller.go
├── docs/
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── models/
│   └── user.go
├── database/
│   └── database.go
├── main.go
├── go.mod
├── go.sum
├── test/
│   └── user_test.go
```

## Instalação e Execução

### Pré-requisitos

- Go 1.22.3 ou superior
- MySQL (para ambiente de produção ou para executar localmente)
### Instalar Go

#### Windows

1. Baixe o instalador do Go para Windows em [https://go.dev/dl](https://go.dev/dl).
2. Execute o instalador e siga as instruções na tela.
3. Após a instalação, abra o Prompt de Comando e verifique a instalação com o comando:

    ```sh
    go version
    ```

#### Ubuntu (Linux)

1. Baixe o tarball do Go para Linux em [https://go.dev/dl](https://go.dev/dl).
2. Extraia o tarball para `/usr/local`:

    ```sh
    sudo tar -C /usr/local -xzf go1.22.3.linux-amd64.tar.gz
    ```

3. Adicione o diretório Go ao PATH:

    ```sh
    echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.profile
    source ~/.profile
    ```

4. Verifique a instalação com o comando:

    ```sh
    go version
    ```

### Clonar o Repositório

```sh
git clone https://github.com/gledson85/golang_rest_orm_unit_test.git
cd golang_rest_orm_unit_test
```

### Configurar Variáveis de Ambiente

Adicione o diretório do Go executável (`$GOPATH/bin`) ao PATH do sistema.

### Instalar Dependências

```sh
go mod tidy
```

### Configurar o Banco de Dados

Crie um arquivo `config/config.yaml` com a configuração do banco de dados:

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
### Instalar e Configurar Swagger

#### Windows

1. Baixe o executável do Swag CLI: [swag.exe](https://github.com/swaggo/swag/releases/download/v1.7.1/swag_windows_amd64.exe).
2. Renomeie o arquivo para `swag.exe`.
3. Adicione o diretório onde o `swag.exe` está localizado ao PATH do sistema.
4. Abra um terminal e verifique a instalação com o comando:

    ```sh
    swag --version
    ```

#### Linux

1. Instale o Swag CLI usando o comando:

    ```sh
    go install github.com/swaggo/swag/cmd/swag@latest
    ```

2. Certifique-se de que o diretório `$GOPATH/bin` está incluído no seu PATH.
3. Verifique a instalação com o comando:

    ```sh
    swag --version
    ```

### Gerar a Documentação Swagger

```sh
swag init
```

### Executar a Aplicação

```sh
go run main.go
```

Acesse a API no endereço: `http://localhost:8081/swagger/index.html`

## Endpoints

### Criar Usuário

- **URL**: `/users`
- **Método**: `POST`
- **Exemplo de Requisição**:

  ```sh
  curl -X POST http://localhost:8081/users -H "Content-Type: application/json" -d '{
    "name": "James Silva",
    "email": "james.silva@mail.com"
  }'
  ```

- **Exemplo de Resposta**:

  ```json
  {
    "id": 1,
    "name": "James Silva",
    "email": "james.silva@mail.com"
  }
  ```

### Obter Usuário

- **URL**: `/users/{id}`
- **Método**: `GET`
- **Exemplo de Requisição**:

  ```sh
  curl -X GET http://localhost:8081/users/1
  ```

- **Exemplo de Resposta**:

  ```json
  {
    "id": 1,
    "name": "James Silva",
    "email": "james.silva@mail.com"
  }
  ```

### Obter Todos os Usuários

- **URL**: `/users`
- **Método**: `GET`
- **Exemplo de Requisição**:

  ```sh
  curl -X GET http://localhost:8081/users
  ```

- **Exemplo de Resposta**:

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

### Atualizar Usuário

- **URL**: `/users/{id}`
- **Método**: `PUT`
- **Exemplo de Requisição**:

  ```sh
  curl -X PUT http://localhost:8081/users/1 -H "Content-Type: application/json" -d '{
    "name": "James",
    "email": "james@mail.com"
  }'
  ```

- **Exemplo de Resposta**:

  ```json
  {
    "id": 1,
    "name": "James Silva",
    "email": "james.silva@mail.com"
  }
  ```

### Deletar Usuário

- **URL**: `/users/{id}`
- **Método**: `DELETE`
- **Exemplo de Requisição**:

  ```sh
  curl -X DELETE http://localhost:8081/users/1
  ```

- **Exemplo de Resposta**:

  ```json
  {
    "message": "User deleted"
  }
  ```

## Testes

### Executar Testes

```sh
go test ./...
```

### Descrição dos Testes

Os testes estão localizados no diretório `test/` e cobrem as operações CRUD. Cada teste cria dados de exemplo, realiza as operações e verifica os resultados esperados.

- **TestCreateUser**: Testa a criação de um novo usuário.
- **TestGetUser**: Testa a obtenção de um usuário específico pelo ID.
- **TestGetAllUsers**: Testa a obtenção de todos os usuários.
- **TestUpdateUser**: Testa a atualização de um usuário.
- **TestDeleteUser**: Testa a exclusão de um usuário.

## Contribuição

Sinta-se à vontade para contribuir com este projeto. Para contribuir:

1. Fork este repositório.
2. Crie uma nova branch: `git checkout -b minha-nova-funcionalidade`.
3. Faça suas alterações e commit: `git commit -m 'Adicionar nova funcionalidade'`.
4. Envie para a branch original: `git push origin minha-nova-funcionalidade`.
5. Crie uma pull request.

## Licença

Este projeto está licenciado sob a Licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.
