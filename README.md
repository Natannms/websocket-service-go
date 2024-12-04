# WebSocket Service


## Pré-requisitos

Antes de começar, você precisará de:

- Go 1.19+
- Docker e Docker Compose (opcional, mas recomendado para rodar o banco de dados PostgreSQL)
- Node.js e NPM (caso queira usar `air` para auto-reload)

## Instalação

1. **Clone o repositório:**
   ```bash
   git clone https://github.com/Natannms/websocket-service-go
   cd websocket-service
   ```

2. **Configuração do arquivo .env:**
   ```bash
   cp .env.example .env
   ```
   
   Abra o arquivo `.env` e configure as variáveis de ambiente de acordo com suas necessidades.
   
   Exemplo de configuração para um banco de dados local:
   ```
   POSTGRES_USER=root
   POSTGRES_PASSWORD=root
   POSTGRES_HOST=localhost
   POSTGRES_PORT=5432
   POSTGRES_DB=pg_websocket_service
   ```

## Opções para Rodar o Projeto

### 1. Rodando o Banco de Dados com Docker (Opcional)

Se você não tem um banco de dados PostgreSQL configurado, pode rodar o banco de dados utilizando o Docker.

No diretório raiz do projeto, execute o seguinte comando para iniciar o banco de dados com Docker Compose:

```bash
docker-compose up -d
```

Isso irá iniciar um contêiner do PostgreSQL, que estará disponível na porta 5432.

### 2. Rodando o Servidor Go

Após configurar o `.env`, inicie o servidor Go:

```bash
go run main.go
```

O servidor será iniciado na porta 8080. Você pode acessá-lo em http://localhost:8080.

### 3. Usando Air para Auto-reload (Opcional, para Desenvolvimento)

Se você quiser utilizar o Air para auto-reload durante o desenvolvimento, execute o seguinte comando no diretório raiz do projeto:

```bash
air
```

Isso permitirá que a aplicação seja recarregada automaticamente sempre que você fizer alterações no código.

### 4. Acessando o Servidor

Após iniciar o servidor, você pode acessar a aplicação em:

```
http://localhost:8080
```

Se você configurou o banco de dados PostgreSQL corretamente, os eventos serão armazenados na tabela `events`.

## Estrutura do Projeto

- `main.go`: Arquivo principal que inicializa o servidor e a conexão com o banco de dados.
- `db/`: Pacote responsável pela conexão com o banco de dados e migrações.
- `web/`: Pacote que define as rotas e o gerenciamento das conexões WebSocket.