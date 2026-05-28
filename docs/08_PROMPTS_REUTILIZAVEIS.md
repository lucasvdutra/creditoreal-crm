# 08 — Prompts reutilizáveis para LLM / Antigravity

## 1. Prompt mestre do projeto

```text
Você é um agente especialista em GoLang, React, PostgreSQL, sqlc, arquitetura multi-tenant e integrações imobiliárias. Trabalhe neste projeto de forma incremental, com baixo consumo de contexto. Não implemente além do escopo da tarefa. Preserve valores internos em pt_BR e trate VRSync apenas como camada de integração. Toda consulta sensível deve respeitar tenant, relacionamento, visibilidade e permissão efetiva. Produza código testável, modular e auditável.
```

## 2. Prompt para planejar uma fase

```text
Com base no roadmap da fase [X], transforme a fase em no máximo 8 task cards pequenos. Para cada card, indique objetivo, arquivos prováveis, migrations necessárias, endpoints, testes e critério de aceite. Não escreva código ainda.
```

## 3. Prompt para backend Go

```text
Implemente apenas o backend da tarefa abaixo. Use Go, PostgreSQL e sqlc. Crie migrations, queries, repository, service, handler REST, DTOs, validações e testes focalizados. Não implemente frontend. Antes de codar, liste os arquivos que pretende criar/alterar. Garanta tenant_id, autorização quando aplicável, erros padronizados e auditoria para ações sensíveis.

Tarefa:
[descrever tarefa]

Critérios de aceite:
[colar critérios]
```

## 4. Prompt para migration + sqlc

```text
Crie apenas a migration e as queries sqlc para [recurso]. Não implemente handlers ou frontend. A tabela deve seguir os padrões do projeto: UUID/ULID, created_at, updated_at, deleted_at quando aplicável, índices por tenant_id, constraints e slugs internos em pt_BR. Inclua queries nomeadas para criar, buscar por id, listar paginado, atualizar e soft delete.
```

## 5. Prompt para service e authorization

```text
Implemente o service de [recurso] usando as queries existentes. O service deve validar permissões pelo authorization service, garantir escopo de tenant, retornar erros padronizados e registrar auditoria quando aplicável. Não altere migrations nesta tarefa, salvo se encontrar inconsistência crítica e justificar antes.
```

## 6. Prompt para handler REST

```text
Implemente os handlers REST de [recurso] usando o service existente. Não coloque regra de negócio no handler. O handler deve validar entrada, chamar o service, traduzir erros para HTTP e retornar JSON no padrão do projeto. Atualize OpenAPI apenas para estes endpoints.
```

## 7. Prompt para frontend React

```text
Implemente apenas a tela frontend de [recurso]. Use React + TypeScript, TanStack Query, React Hook Form e Zod. Consuma os endpoints já existentes. Não altere backend. A tela deve ter loading, erro, estado vazio, paginação e validação. Respeite permissões recebidas da API para exibir/ocultar ações.
```

## 8. Prompt para VRSyncMapper

```text
Implemente apenas o VRSyncMapper para [tipo de mapeamento]. O domínio interno usa pt_BR e o VRSync só pode aparecer no mapper, em integration_mappings, nos snapshots de payload/XML e nos logs de integração. Crie testes para conversões válidas e para erro de mapeamento inexistente. Não altere entidades principais do domínio para salvar valores externos.
```

## 9. Prompt para validação VRSync

```text
Implemente o VRSyncValidationService para validar se um imóvel pode ser exportado. O serviço deve retornar lista estruturada de erros, sem gerar XML. Validar campos obrigatórios, mídia mínima, endereço mínimo, preço compatível com tipo de transação, tipo de imóvel, perfil de contato e mapeamentos existentes. Inclua testes para cada erro.
```

## 10. Prompt para exportação VRSync

```text
Implemente o VRSyncExportService. Ele deve receber um tenant, buscar imóveis elegíveis, validar, mapear valores pt_BR para VRSync, gerar XML, salvar snapshot/hash em property_publications e retornar o feed. Não salvar valores VRSync nas entidades principais. Inclua testes de XML mínimo válido e erro de imóvel inválido.
```

## 11. Prompt para revisão de segurança

```text
Revise a implementação abaixo procurando falhas de multi-tenancy, autorização, LGPD, logs sensíveis, comandos destrutivos, queries sem tenant_id, endpoints públicos sem proteção e valores VRSync salvos no domínio. Retorne uma lista objetiva de problemas, severidade, arquivo/linha e correção recomendada. Não faça alterações ainda.
```

## 12. Prompt para refatoração controlada

```text
Refatore apenas [módulo/arquivo] para melhorar [objetivo]. Não altere comportamento externo. Antes de alterar, liste riscos e testes que devem continuar passando. Preserve contratos públicos, nomes de endpoints, migrations existentes e regras de autorização. Ao final, informe diffs relevantes e como validar.
```

## 13. Prompt para gerar testes

```text
Crie testes focalizados para [módulo]. Cubra casos de sucesso, validação, erro, acesso negado por tenant, permissão insuficiente e auditoria quando aplicável. Não altere a implementação, exceto pequenos ajustes necessários para testabilidade, justificando cada ajuste.
```

## 14. Prompt para corrigir bug

```text
Investigue e corrija apenas o bug descrito. Primeiro explique a causa provável, depois altere o menor número de arquivos possível. Inclua teste que reproduz o bug e passa após a correção. Não faça refatorações oportunistas.

Bug:
[descrever bug]
```
