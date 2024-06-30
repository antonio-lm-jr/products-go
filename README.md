## Projeto

Criar uma API que possibilite o cadastro de produtos. Esse projeto foi iniciado com `go mod init github.com/antonio-lm-jr/products-go`. O comendo irá criar o arquivo `go.mod`.


## Objetivo

Criar uma pequena API em golang com fober para exercitar alguns aprendizados da linguagem.


## Configurando o ambiente

### Docker

1. Instale o [Docker](https://docs.docker.com/engine/install/).
2. Abra seu shell favorito.
3. Navegue para a a pasta `/dockerfiles`.
4. Execute `docker-compose up -d`.

### Banco de dados

1. Em seu navegador acesse `http://localhost:8080/`.
2. Acesse o admin, com usuário e senha definido no arquivo `docker-compose.yaml`.
3. Crie a tabela abaixo.

   ```
   CREATE TABLE products
   (
     identifier  varchar(50)   NOT NULL,
     name        varchar(50)   NOT NULL,
     description varchar(100)  NOT NULL,
     price       decimal(10,2) NOT NULL
   );
   ```

### Aplicação

1. Instale o [golang](https://go.dev/).
2. Abra seu shell favorito.
3. Para instalar as dependências execute `go mod tidy` ou `go get`.
4. Para iniciar o projeto execute `go run main.go`.


## Aprendizados

- O arquivo go.sum
  
  Assim como o node têm o arquivo package-lock.json o arquivo go.sum é o arquivo que contém as versões exatas dos pacotes.


## Comando uteis

- Instalando um novo pacote.

  ```
  > go get -u {{pacote}}
  > go get -u github.com/gofiber/fiber/v2
  ```

- Removendo um pacote.

  Remova todas as importações e códigos relacionados ao pacote e execute
  
  ```
  > go mod tidy -v
  ```

- Formatando o código.

  ```
  > gofmt -w .
  ```