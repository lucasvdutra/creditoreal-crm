# 07 — Segurança, multi-tenancy e LGPD

## 1. Princípio central

Em um sistema multi-tenant, falha de escopo é falha crítica. Nenhum endpoint deve confiar apenas no frontend.

## 2. Autorização efetiva

Toda operação deve considerar:

```text
usuário autenticado
tenant atual do usuário
tenant dono do recurso
papel do usuário
permissões do papel
relacionamento entre tenants
permissões do relacionamento
visibilidade do recurso
status das entidades envolvidas
```

## 3. Regras obrigatórias para queries

Tabelas por tenant devem ser consultadas com filtro explícito:

```sql
WHERE tenant_id = $1
```

ou, no caso de imóveis:

```sql
WHERE tenant_dono_id = $1
```

Quando houver acesso por relacionamento, a query deve passar por serviço específico de autorização/visibilidade.

## 4. Dados sensíveis

Dados pessoais de leads devem ser tratados como sensíveis:

```text
nome
email
telefone
mensagem
UTMs quando puderem identificar comportamento
consentimento LGPD
```

## 5. LGPD

O sistema deve prever:

- registro de consentimento;
- data/hora do consentimento;
- origem do consentimento;
- política de retenção;
- possibilidade de anonimização ou exclusão conforme base legal;
- controle de acesso a dados pessoais;
- auditoria de acesso e alteração.

## 6. Tokens e segredos

Nunca salvar tokens externos em texto puro.

Obrigatório:

- criptografia em repouso para tokens de integração;
- variáveis de ambiente para segredos;
- rotação de credenciais;
- logs sem tokens.

## 7. Segurança de endpoints públicos

Endpoints públicos, como captação de leads e feed VRSync, devem ter:

- rate limit;
- validação forte;
- proteção contra spam;
- CORS controlado;
- logs de abuso;
- autenticação ou token quando aplicável.

## 8. Proteção contra enumeração

Usar UUID/ULID em identificadores expostos.

Não expor IDs sequenciais internos em URLs públicas.

## 9. Auditoria obrigatória

Auditar:

```text
login
logout
criação/alteração/exclusão de tenant
alteração de permissão
criação/alteração/exclusão de imóvel
publicação de imóvel
exportação VRSync
erro de integração
criação/atribuição de lead
envio de mensagem
alteração de relacionamento entre tenants
```

## 10. Regras para Antigravity/LLM

O agente não deve:

- criar secrets hardcoded;
- imprimir tokens em logs;
- usar comandos destrutivos sem aprovação;
- remover validações de tenant para “simplificar”;
- criar bypass de autorização para testes;
- salvar payload pessoal completo em logs;
- aceitar `tenant_id` do body sem validar se o usuário pode atuar naquele tenant.

## 11. Checklist de revisão de segurança

Antes de aceitar uma tarefa, verificar:

```text
[ ] Há filtro de tenant em todas as queries sensíveis?
[ ] Há autorização backend no service/handler?
[ ] Status de relacionamento é validado?
[ ] Permissões efetivas são consideradas?
[ ] Dados pessoais não aparecem em logs?
[ ] Tokens não estão hardcoded?
[ ] Erros não vazam stack trace?
[ ] Auditoria foi registrada para ação sensível?
[ ] Endpoint público tem rate limit ou mitigação equivalente?
[ ] Testes cobrem acesso negado?
```
