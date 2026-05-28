# Instruções LLM / Antigravity — Projeto Plataforma Imobiliária Multi-Tenant

Este pacote reorganiza o escopo técnico da plataforma em arquivos menores para uso com LLMs e Antigravity, com foco em baixo consumo de contexto, redução de retrabalho e melhor performance de desenvolvimento.

## Objetivo

Orientar o desenvolvimento incremental de uma plataforma web multi-tenant para gestão imobiliária, publicação de imóveis, integração VRSync, captação de leads, comunicação comercial, CRM, auditoria e dashboards.

## Stack-base

- Backend: GoLang.
- Frontend: React + TypeScript.
- Banco principal: PostgreSQL.
- Cache/fila inicial: Redis.
- Storage: S3 compatível.
- Integração externa principal: VRSync / Grupo OLX-ZAP.
- Idioma funcional e valores internos de domínio: pt_BR.

## Como usar com LLM / Antigravity

Não envie o escopo completo em toda solicitação. Trabalhe sempre com um **Context Pack** mínimo:

1. `01_CONTEXT_PACK_TEMPLATE.md`.
2. O arquivo de fase correspondente em `03_ROADMAP_FASEADO.md`.
3. O arquivo de regras aplicável em `antigravity-rules/`.
4. Apenas os trechos do domínio necessários para a tarefa atual.
5. O critério de aceite da tarefa.

## Ordem recomendada de leitura

1. `00_OPERACAO_LLM_BAIXO_CONSUMO.md` — como economizar tokens e evitar agentes perdidos.
2. `01_CONTEXT_PACK_TEMPLATE.md` — modelo de prompt enxuto para cada tarefa.
3. `02_ARQUITETURA_E_DECISOES.md` — decisões técnicas obrigatórias.
4. `03_ROADMAP_FASEADO.md` — fases de desenvolvimento.
5. `04_BACKLOG_MVP_TASK_CARDS.md` — cartões de tarefas prontos para execução.
6. `05_CONTRATOS_E_PADROES.md` — padrões de API, banco, código e nomenclatura.
7. `06_QUALIDADE_TESTES_PERFORMANCE.md` — testes, performance e qualidade.
8. `07_SEGURANCA_MULTITENANT_LGPD.md` — segurança, autorização e LGPD.
9. `08_PROMPTS_REUTILIZAVEIS.md` — prompts para agentes.
10. `09_CHECKLISTS_DE_ACEITE.md` — checklists finais por módulo.

## Organização do pacote

```text
llm_antigravity_instrucoes/
├── README.md
├── 00_OPERACAO_LLM_BAIXO_CONSUMO.md
├── 01_CONTEXT_PACK_TEMPLATE.md
├── 02_ARQUITETURA_E_DECISOES.md
├── 03_ROADMAP_FASEADO.md
├── 04_BACKLOG_MVP_TASK_CARDS.md
├── 05_CONTRATOS_E_PADROES.md
├── 06_QUALIDADE_TESTES_PERFORMANCE.md
├── 07_SEGURANCA_MULTITENANT_LGPD.md
├── 08_PROMPTS_REUTILIZAVEIS.md
├── 09_CHECKLISTS_DE_ACEITE.md
├── antigravity-rules/
│   ├── 01_workspace_rules.md
│   ├── 02_go_backend_rules.md
│   ├── 03_react_frontend_rules.md
│   ├── 04_database_rules.md
│   ├── 05_security_rules.md
│   └── 06_vrsync_rules.md
└── workflows/
    ├── workflow_criar_modulo_backend.md
    ├── workflow_criar_tela_frontend.md
    ├── workflow_vrsync_exportacao.md
    ├── workflow_revisao_seguranca.md
    └── workflow_refatoracao_controlada.md
```

## Princípio central

Cada tarefa deve ser pequena, verificável e concluída por uma entrega concreta: migration, endpoint, service, teste, tela, componente, mapper, validação ou documentação.

Evite comandos amplos como “implemente o sistema inteiro”. Prefira comandos como:

> Implemente apenas o CRUD de tenants no backend, com migration, queries sqlc, service, handler REST, testes de service e validação de tenant_id. Não implemente frontend nesta tarefa.
