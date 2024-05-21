# REST API - Golang

## Tecnologias utilizadas
* Go (Linguagem de programação)
* PostgreSQL (Banco de dados relacional)
* Gin-Gonic (Framework web)
* Docker

## Descrição
API com a finalidade de exercitar os principais conceitos de desenvolvimento de software.
A API a ser desenvolvida vai ter as funcionalidades de cadastrar um novo produto, buscar um produto específico e listar todos os produtos.
Ao final, a API e o banco de dados vão estar rodando em um container Docker como forma de estudo na utilização dessa tecnologia.

## Arquitetura da aplicação
A aplicação foi estruturada partindo de uma arquitetura limpa onde teremos a nossa camada de modelo (Model), parte mais interna da aplicação, uma camada de consulta ao banco de dados (Repository), uma camada de lógica de negócio (Use Cases) e a camada responsável pelas rotas da aplicação (Controller), ou seja, a camada de contato com o nosso usuário onde ele irá fazer as requisições.