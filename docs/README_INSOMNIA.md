# Importando a Coleção do PsyGrow API no Insomnia

Este diretório contém arquivos para facilitar o teste e uso da API PsyGrow.

## Arquivo de Coleção do Insomnia

O arquivo `insomnia_collection.json` contém uma coleção completa de todas as rotas da API PsyGrow, organizadas por recursos (autenticação, usuários, pacientes, etc.) e com exemplos de requisições.

## Como Importar

1. Abra o Insomnia
2. Clique em "Create" (ou no ícone "+") no canto superior
3. Selecione "Import from File"
4. Navegue até o diretório onde você salvou o arquivo `insomnia_collection.json`
5. Selecione o arquivo e clique em "Open"
6. A coleção "PsyGrow API" será importada com todas as rotas organizadas em pastas

## Estrutura da Coleção

A coleção está organizada nas seguintes pastas:

- **Auth**: Rotas de autenticação (login)
- **Users**: Gerenciamento de usuários (criar, listar, atualizar)
- **Patients**: Gerenciamento de pacientes
- **Appointments**: Agendamentos e consultas
- **Anamnese**: Templates e campos de anamnese
- **Financial**: Centro de custos e pagamentos

## Variáveis de Ambiente

A coleção já vem configurada com as seguintes variáveis de ambiente:

- `base_url`: URL base da API (padrão: http://localhost:8080/api/v1)
- `token`: Token de autenticação (você precisará preencher após fazer login)

Para usar o token de autenticação:

1. Execute a requisição "Login" na pasta Auth
2. Copie o token da resposta
3. Clique em "Environment" no canto superior direito
4. Cole o token no campo "token" da variável de ambiente

## Alternativa: Importar via Swagger

Se preferir, você também pode importar diretamente o arquivo Swagger:

1. No Insomnia, clique em "Create" > "Import From" > "File"
2. Selecione o arquivo `swagger.json` neste diretório
3. O Insomnia irá criar automaticamente as requisições baseadas na especificação OpenAPI