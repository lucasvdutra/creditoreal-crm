# 00 — Operação LLM / Antigravity com baixo consumo e melhor performance

## 1. Regra de ouro

O agente não deve receber o escopo completo em toda tarefa. O escopo completo serve como referência-mãe, mas o desenvolvimento deve ser conduzido por **pacotes de contexto pequenos**.

Um bom pacote de contexto deve caber em uma única solicitação curta e conter apenas:

- objetivo da tarefa;
- arquivos prováveis a modificar;
- invariantes de arquitetura;
- contrato esperado;
- critérios de aceite;
- testes obrigatórios.

## 2. Estratégia de contexto mínimo

Use a seguinte regra:

```text
Contexto global fixo + contexto da fase + contexto da tarefa + arquivos diretamente afetados.
```

Não inclua catálogos inteiros de características, tipos de imóveis ou mapeamentos VRSync quando a tarefa não precisar deles. Para catálogos grandes, trabalhe com seed files e mappers isolados.

## 3. Divisão recomendada por agentes

Quando usar agentes paralelos, separe por responsabilidade, não por “partes aleatórias” do código.

### Agente 1 — Backend/domain

Responsável por:

- entidades de domínio;
- services;
- validações;
- autorização;
- casos de uso.

### Agente 2 — Banco/sqlc

Responsável por:

- migrations;
- índices;
- queries sqlc;
- constraints;
- seeds.

### Agente 3 — API/contratos

Responsável por:

- rotas REST;
- DTOs;
- OpenAPI;
- handlers;
- erros HTTP.

### Agente 4 — Frontend

Responsável por:

- telas;
- componentes;
- formulários;
- hooks de API;
- estados de carregamento e erro.

### Agente 5 — Revisão

Responsável por:

- verificar escopo de tenant;
- revisar permissões;
- revisar testes;
- revisar impactos de performance;
- apontar arquivos desnecessariamente alterados.

## 4. Padrão de execução por tarefa

Toda tarefa deve seguir este ciclo:

```text
1. Entender o objetivo.
2. Ler apenas arquivos relevantes.
3. Escrever plano curto.
4. Alterar poucos arquivos.
5. Criar ou ajustar testes.
6. Rodar testes focalizados.
7. Entregar resumo com arquivos alterados e pendências.
```

## 5. Limites práticos por tarefa

Para manter boa performance do LLM:

- alterar no máximo 5 a 8 arquivos por tarefa comum;
- criar no máximo 1 módulo funcional por execução;
- evitar refatorações globais durante implementação de feature;
- não misturar backend, frontend, banco e infraestrutura na mesma tarefa, salvo em uma vertical slice planejada;
- evitar reprocessar catálogos grandes em prompts;
- pedir sempre diffs pequenos e justificáveis.

## 6. Antipadrões que aumentam custo e reduzem performance

Evite prompts como:

```text
Implemente todo o sistema conforme o escopo.
Crie todas as tabelas e todo o frontend.
Analise o repositório inteiro e corrija tudo.
Refatore a arquitetura completa.
```

Prefira prompts como:

```text
Crie apenas o módulo tenant no backend: migration, queries, repository, service, handler e testes. Siga as regras de tenant_id, auditoria básica e slugs pt_BR. Não implemente frontend.
```

## 7. Como reduzir alucinação técnica

O agente deve:

- consultar o contrato OpenAPI antes de criar telas;
- consultar migrations antes de criar queries;
- consultar services antes de criar handlers;
- consultar regras de autorização antes de expor dados;
- nunca inventar valores externos VRSync fora do mapper;
- nunca salvar valores VRSync diretamente no domínio.

## 8. Política de alteração segura

Antes de alterar código existente, o agente deve informar:

```text
Arquivos que pretendo alterar:
1. ...
2. ...

Arquivos que pretendo criar:
1. ...
2. ...

Riscos:
- ...
```

Para tarefas pequenas, esse plano pode ter 5 a 10 linhas.

## 9. Comandos destrutivos

O agente nunca deve executar sem aprovação explícita:

```bash
rm -rf
DROP DATABASE
DROP TABLE
TRUNCATE
DELETE sem WHERE
git reset --hard
git clean -fd
```

Migrations destrutivas devem ser substituídas por estratégia de migração segura:

```text
1. criar novo campo;
2. backfill;
3. validar;
4. trocar leitura/escrita;
5. remover campo antigo em release posterior.
```

## 10. Resultado esperado de cada execução

Toda entrega deve terminar com:

```text
Resumo do que foi feito.
Arquivos alterados.
Como testar.
O que ficou fora do escopo.
Riscos ou pendências.
```
