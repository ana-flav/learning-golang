# Taylor Swift Songs API

Este é um projeto simples em Go, uma API REST para gerenciar músicas da Taylor Swift, com operações CRUD.

## Pré-requisitos

Certifique-se de ter o Go instalado na sua máquina. Você pode baixá-lo em [golang.org](https://golang.org/).

Também é necessário ter um banco de dados SQL para estabelecer as conexões necessárias. No diretório do projeto, crie um arquivo `.env` com base no exemplo fornecido em `.env.example`, e preencha as informações do seu banco de dados PostgreSQL.

Exemplo do arquivo `.env`:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=seu_usuario
DB_PASSWORD=sua_senha
DB_NAME=nome_do_banco
```

Certifique-se de criar o banco de dados correspondente no PostgreSQL.

## Instalação e Execução

1. Clone o repositório:

    ```bash
    git clone https://github.com/ana-flav/learning-golang.git
    cd seu-projeto
    ```

2. Baixe as dependências do Go:

    ```bash
    go mod download
    ```

3. Execute o projeto (certifique de estar na dentro do diretório cmd):

    ```bash
    go run main.go
    ```

A API estará disponível em [http://localhost:8000](http://localhost:8000).

## Exemplo de Uso com Postman

Vamos fornecer um exemplo de como usar o Postman para enviar uma solicitação POST e adicionar uma música usando o endpoint da sua API.

1. Abra o Postman e crie uma nova solicitação.

2. Selecione o método POST no canto superior esquerdo.

3. Insira a URL da sua API para adicionar uma música, por exemplo: `http://localhost:8000/add-song`.

4. Vá para a guia "Body" e selecione "raw" e escolha "JSON (application/json)" no menu.

5. No corpo da solicitação, insira os detalhes da música que você deseja adicionar. Por exemplo:

    ```json
    {
        "Name": "mirrorball",
        "LinkSong": "https://open.spotify.com/intl-pt/track/0ZNU020wNYvgW84iljPkPP?si=20728167f34247c4"
    }
    ```

6. Clique no botão "Send" para enviar a solicitação.

Isso enviará uma solicitação POST para o endpoint `http://localhost:8000/add-song` com os detalhes da música fornecidos no corpo da solicitação. Lembre-se de adaptar os detalhes conforme necessário para corresponder aos campos esperados em sua aplicação.

Lembre-se de que este é apenas um exemplo e você pode ajustar os detalhes conforme necessário para a sua implementação específica.