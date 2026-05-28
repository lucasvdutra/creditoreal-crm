# Go Backend Rules

## Stack

- GoLang.
- REST API.
- PostgreSQL.
- sqlc para queries.
- Migrations com goose, golang-migrate ou ferramenta definida no projeto.
- Logs estruturados com zap ou zerolog.
- Testes com `testing` e `testify`.

## Arquitetura

Separar:

```text
handler
DTO
service/use case
repository/sqlc
authorization
audit
```

Handlers não devem conter regra de negócio complexa.

## Contexto e banco

- Toda chamada de banco recebe `context.Context`.
- Toda operação composta deve usar transaction.
- Queries de dados por tenant devem filtrar `tenant_id` ou `tenant_dono_id`.
- Não retornar erro bruto de banco para HTTP.

## Erros

Usar erros padronizados com código técnico em pt_BR:

```text
imovel.nao_encontrado
permissao.negada
tenant.invalido
vrsync.mapeamento_inexistente
```

## Testes

Todo service novo deve ter testes para:

- sucesso;
- validação;
- permissão negada;
- tenant incorreto;
- erro de repository quando relevante.
