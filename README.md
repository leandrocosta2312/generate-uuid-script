# Generate UUID for script's sql

### Instale o go em sua maquina
https://golang.org/doc/install

### Baixe as seguintes dependecias

go get github.com/satori/go.uuid

### Para executar localmente:

Necessário colocar dentro da pasta raiz um arquivo chamado **input.sql** \
Obs: Seu script obrigatoriamente precisa ter o texto **$uuid**, pois esse valor que será substituido pelo UUID gerado pelo go.

### Execute o seguinte comando
go run generate_uuid_script.go

Após isso sera criado o arquivo **output.sql** com o UUIDs gerados.
