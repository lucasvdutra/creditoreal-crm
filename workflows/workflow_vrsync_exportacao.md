# Workflow — Implementar exportação VRSync

## Entrada necessária

- Campos internos do imóvel.
- Mapeamentos necessários.
- Regras de obrigatoriedade.
- Estrutura XML esperada.

## Passos

1. Confirmar que domínio está em pt_BR.
2. Criar/validar `integration_mappings`.
3. Implementar `VRSyncMapper`.
4. Implementar `VRSyncValidationService`.
5. Implementar `VRSyncExportService`.
6. Salvar snapshot/hash em `property_publications`.
7. Criar endpoint de validação.
8. Criar endpoint/feed XML.
9. Criar testes unitários e de integração.

## Regras críticas

- Não salvar valores VRSync em `properties`, `features` ou `property_types`.
- Não gerar XML se validação falhar.
- Erros devem ser claros para usuário operacional.
