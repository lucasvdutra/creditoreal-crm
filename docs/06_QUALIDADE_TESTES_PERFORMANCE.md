# 06 — Qualidade, testes e performance

## 1. Estratégia de testes

A qualidade deve ser construída desde o primeiro módulo. Não deixar testes para o final.

## 2. Pirâmide mínima

```text
Unit tests: services, mappers, validators, authorization.
Integration tests: repositories, migrations, endpoints críticos.
E2E smoke tests: fluxos principais do MVP.
```

## 3. Testes obrigatórios por tipo de módulo

### Backend CRUD

Testar:

- criação válida;
- campo obrigatório ausente;
- status inválido;
- listagem paginada;
- filtro por tenant;
- acesso negado a tenant não autorizado;
- atualização parcial;
- soft delete;
- auditoria quando aplicável.

### Authorization service

Testar:

- usuário do tenant dono;
- usuário de tenant relacionado com permissão;
- usuário de tenant relacionado sem permissão;
- relacionamento suspenso/expirado;
- recurso público com dados sensíveis omitidos;
- papel sem permissão.

### VRSyncMapper

Testar:

- conversão de tipo de transação;
- conversão de tipo de imóvel;
- conversão de características;
- conversão de garantias;
- erro para mapeamento inexistente;
- garantia de que valores VRSync não aparecem no domínio interno.

### VRSyncValidationService

Testar:

- imóvel sem imagem;
- imóvel sem CEP;
- imóvel sem bairro;
- imóvel sem preço compatível com transação;
- imóvel sem perfil de contato;
- tipo residencial sem campos mínimos quando exigido;
- geração de lista de erros amigáveis.

### Frontend

Testar ou validar manualmente:

- tela carrega lista paginada;
- filtro funciona;
- formulário valida campos obrigatórios;
- erros da API aparecem na interface;
- estado de loading;
- estado vazio;
- permissão remove/oculta ação visual.

## 4. Performance obrigatória desde o MVP

### Banco

- índices por `tenant_id`;
- índices compostos por filtros frequentes;
- evitar `SELECT *` em listagens;
- usar paginação obrigatória;
- analisar queries lentas com `EXPLAIN` quando necessário;
- criar constraints para integridade, não apenas validação na aplicação.

### Backend

- evitar N+1;
- usar `context.Context` com timeout;
- separar listagem de detalhe;
- não retornar payloads grandes sem necessidade;
- comprimir respostas se aplicável;
- jobs assíncronos para tarefas custosas.

### Frontend

- usar TanStack Query para cache;
- debounce em filtros de busca;
- virtualização em tabelas muito grandes;
- paginação server-side;
- formulários divididos por abas/seções quando o imóvel tiver muitos campos;
- evitar re-render desnecessário em tabelas densas.

## 5. Jobs assíncronos

Devem ser assíncronos:

```text
processar_imagens
gerar_feed_vrsync
validar_imovel_vrsync
sincronizar_publicacoes
importar_leads
deduplicar_leads
atribuir_leads
enviar_notificacoes
calcular_dashboards
limpar_tokens_expirados
```

## 6. Observabilidade mínima

Desde o MVP, registrar:

- request_id;
- duração da requisição;
- status HTTP;
- endpoint;
- user_id quando autenticado;
- tenant_id quando aplicável;
- erro padronizado.

## 7. Métricas úteis

```text
latência p95 por endpoint
erros 4xx e 5xx por endpoint
tempo de geração VRSync
quantidade de imóveis com erro de integração
tempo médio de primeira resposta a lead
leads por canal
```

## 8. Definition of Done técnica

Uma tarefa só é concluída quando:

```text
- compila;
- migrations rodam;
- testes focalizados passam;
- lint/format foi aplicado;
- não há valor externo VRSync no domínio;
- consultas sensíveis filtram tenant;
- erros são padronizados;
- logs não expõem dados sensíveis;
- documentação mínima foi atualizada.
```

## 9. Estratégia de otimização

Não adicionar tecnologia por antecipação. A ordem correta é:

```text
1. modelagem correta;
2. índices corretos;
3. paginação;
4. queries otimizadas;
5. cache;
6. jobs assíncronos;
7. particionamento;
8. read replica;
9. OpenSearch quando busca exigir.
```
