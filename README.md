# Descrição

Este sistema é um executor de script lua em golang usando a lib `https://github.com/yuin/gopher-lua` usando `https://github.com/urfave/cli` para uma app CLI

# Negócio

Essa ferramenta serve para executar scripts lua dentro do Go para realizar tarefas de forma mais customizavel.

## Casos de uso

### 1 - Listar /tmp e ordenar em ordem alfabética

Ao chamar o comando `list-order` o sistema tem que retornar o conteúdo da pasta `/tmp` ordenado alfabéticamente.
Para isso o script lua `list-order.lua` vai executar os seguintes passos:
1. Chamar a ferramenta implementada em GO que lê o conteúdo da pasta `/tmp` e retorna a lista de nome de arquivo
2. Chamar a ferramenta implementada em GO que ordena a lista de nome de arquivo em ordem alfabética

### 2 - Listar /tmp e colocar tudo em maiusculo e salvar resultado no /tmp

Ao chamar o comando `list-caps` o sistema tem que retornar o conteúdo da pasta `/tmp` com tudo em maiusculo.
Para isso o script lua `list-caps.lua` vai executar os seguintes passos:
1. Chamar a ferramenta implementada em GO que lê o conteúdo da pasta `/tmp` e retorna a lista de nome de arquivo
2. Chamar a ferramenta implementada em GO que varre a lista e transforma tudo em maiusculo


# Esquema de pasta


```text
golang-lua-integration/
├── cmd/
│   └── app/ - lugar que fica a main.go que inicia o servidor
├── internal/
│       ├── di/ - contém código de injeção de dependência
│       ├── adapter/
│       │      ├── lua/
│       │      │    ├── executer/ - script lua executer
│       │      │    └── fileloader/ - script lua loader
│       │      └── cli/
│       │           └── commands/ - cli commands
│       ├── domain/
│       │      ├── services/ - serviços que orquestram as chamas de scripts
│       │      ├── model/ - modelo de dados dos scripts e das ferramentas
│       │      └── tools/ - ferramentas em GO disponível para os scripts lua
│       ├── scripts/ - local dos scripts lua
│       └── util/ - contem codigos utilitários
└── docs/ - documentação
```

Essa ferramenta usa Clean Architecture, onde a lógica de negócio fica na pasta `internal/service` e `internal/model`.

