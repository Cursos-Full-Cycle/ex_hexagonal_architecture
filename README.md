# ex_hexagonal_architecture

### Inicialização da aplicação
go mod init go_hexagonal

### Baixar dependencias de pacotes
go mod tidy

### executar testes
go test ./...


### Criando banco de dados
touch sqlite.db
sqlite3 sqlite.db
create table products(id string, name string, price float, status string);
.tables