##

## Descrição das stacks
### Golang
A forma mais idiomática e direta em Go quando o objetivo é obter um []byte para despachar para um serviço externo como o Kafka, RabbitMQ ou até mesmo para salvar em um banco de dados, é usar o **Marshal**. 

#### Marshal e Unmarshal (Memória/Variáveis)
Use esses dois quando você já tem os dados guardados em uma variável na memória
- **json.Marshal**: Pega uma estrutura de dados do Go (como uma struct, ou nesse caso, um loggingpb.LogEntry) e transforma em uma fatia de bytes ([]byte) no formato JSON.
- **json.Unmarshal**: Faz o caminho inverso. Pega um []byte que contém um JSON e preenche uma variável Go com esses dados.

#### NewEncoder vs. NewDecoder (Streams/Fluxos de Dados)
Use esses dois quando você está lidando com I/O (Entrada/Saída), ou seja, lendo ou escrevendo diretamente de um fluxo de dados (um io.Writer ou io.Reader). Isso é muito comum em requisições web, arquivos ou no próprio terminal.
- **json.NewEncoder**: Cria um "codificador" associado a um destino (como um arquivo ou a resposta de uma API - os.Stdout). Quando você chama .Encode(dados), ele transforma os dados em JSON e já envia direto para esse destino, sem precisar guardar o JSON inteiro na memória antes.
  - o json.NewEncoder não foi feito para segurar dados, mas sim para ser um "tubo" que direciona o fluxo deles.
- **json.NewDecoder**: Cria um "decodificador" associado a uma fonte de dados. Quando você chama .Decode(&variavel), ele lê o JSON diretamente dessa fonte (à medida que os dados chegam) e já preenche sua variável.




## Referências
- https://pkg.go.dev/cloud.google.com/go/logging/logadmin#example-Client.Entries-Pagination
- https://pkg.go.dev/cloud.google.com/go/logging/logadmin#example-EntryIterator.Next
- https://gobyexample.com/json
- http://medium.com/@aalves/golang-transformando-dados-marshal-unmarshal-decode-encode-porque-e-quando-2048ecbe1075
- https://go.dev/blog/pipelines


https://hub.docker.com/r/apache/kafka
https://medium.com/@darshak.kachchhi/setting-up-a-kafka-cluster-using-docker-compose-a-step-by-step-guide-a1ee5972b122
https://dev.to/kaike_castro/criando-um-cluster-do-kafka-com-docker-compose-e-desenvolvendo-um-consumer-e-producer-em-golang-403c

https://pkg.go.dev/google.golang.org/protobuf/encoding/protojson