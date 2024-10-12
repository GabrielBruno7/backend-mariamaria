

# Maria Maria Confeitaria (em desenvolvimento)

## Sobre

Este projeto é o backend da **Maria Maria Confeitaria**, um sistema desenvolvido para gerenciar os processos internos de uma confeitaria. O objetivo do projeto é servir como base para a organização dos produtos, pedidos e precificação de bolos, roscas, e outros produtos confeitados.

O backend está sendo desenvolvido em **Go**, e este projeto também é uma oportunidade de estudo e aprimoramento da linguagem. Go foi escolhido pela sua simplicidade, desempenho e suporte para desenvolvimento de APIs eficientes e escaláveis, tornando-o ideal para projetos de backend como este.

## Tecnologias Utilizadas

- **Go**: Linguagem principal utilizada no desenvolvimento do backend.
- **Gin**: Framework web em Go para criação de APIs RESTful.
- **PostgreSQL** (ou outro banco de dados): Utilizado para armazenar os dados da confeitaria (produtos, pedidos, clientes, etc.).
- **Docker**: Para facilitar o processo de desenvolvimento e deploy.
- **Swagger**: Para documentação automática da API.

## Como rodar o projeto

1. Clone o repositório:
   ```bash
   git clone https://github.com/usuario/backend-mariamaria.git

2. Entre no projeto
   ```bash
   cd backend-mariamaria
   
3. Inicie o docker
   ```bash
   docker compose up --build
   docker compose up -d
