# Workspace Rules — Projeto Plataforma Imobiliária Multi-Tenant

## Escopo do agente

Você está trabalhando em uma plataforma web multi-tenant com backend GoLang, frontend React + TypeScript, PostgreSQL, Redis, S3 compatível e integração VRSync.

## Regras obrigatórias

1. Execute apenas a tarefa solicitada.
2. Não implemente módulos fora do escopo.
3. Não faça refatoração global sem solicitação explícita.
4. Antes de codar, liste brevemente arquivos que pretende criar ou alterar.
5. Trabalhe com diffs pequenos.
6. Preserve valores internos de domínio em pt_BR.
7. Nunca salve valores VRSync nas entidades principais do domínio.
8. Toda consulta sensível deve respeitar tenant, relacionamento, visibilidade e permissão.
9. Não crie secrets hardcoded.
10. Não execute comandos destrutivos sem aprovação explícita.

## Resultado final obrigatório

Ao final de cada tarefa, informe:

```text
- resumo do que foi feito;
- arquivos alterados;
- testes executados;
- como validar manualmente;
- pendências ou riscos.
```

## Baixo consumo de contexto

Não releia o projeto inteiro se a tarefa é local. Leia apenas arquivos relevantes para o módulo.
