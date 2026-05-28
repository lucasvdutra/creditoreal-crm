# Security Rules — Multi-tenant, LGPD e agente seguro

## Regras críticas

1. Nunca confiar apenas no frontend.
2. Nunca aceitar `tenant_id` do request sem validar se o usuário pode atuar nele.
3. Nunca logar senha, token, documento sensível ou payload pessoal completo.
4. Nunca criar bypass de autorização para facilitar teste.
5. Nunca executar comandos destrutivos sem aprovação.
6. Nunca hardcodar secrets.

## Autorização

Toda operação sensível deve validar:

```text
user_id
tenant atual
tenant dono do recurso
role
permission
relationship
visibility
status
```

## LGPD

Dados de lead devem ser tratados com cuidado:

```text
nome
email
telefone
mensagem
consentimento
origem/campanha
```

## Auditoria

Auditar alterações sensíveis, permissões, exportações e erros de integração.
