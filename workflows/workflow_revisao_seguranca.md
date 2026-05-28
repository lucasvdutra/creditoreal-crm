# Workflow — Revisão de segurança

## Objetivo

Revisar uma implementação antes do merge.

## Checklist

```text
[ ] Query sem tenant_id?
[ ] Handler confia em tenant_id do body?
[ ] Authorization service foi chamado?
[ ] Relacionamento ativo foi validado?
[ ] Permissão efetiva foi validada?
[ ] Dados pessoais aparecem em logs?
[ ] Token/secret foi hardcoded?
[ ] Endpoint público tem rate limit/validação?
[ ] Erro expõe stack trace?
[ ] Auditoria existe para ação sensível?
[ ] Teste cobre acesso negado?
```

## Saída esperada

Lista de achados com severidade:

```text
Crítico
Alto
Médio
Baixo
```

Para cada achado:

```text
arquivo
problema
risco
correção recomendada
```
