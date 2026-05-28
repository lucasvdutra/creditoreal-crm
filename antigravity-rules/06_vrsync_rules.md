# VRSync Rules

## Regra central

VRSync é camada de integração, não domínio interno.

## Permitido usar valores VRSync apenas em

```text
VRSyncMapper
VRSyncValidationService
VRSyncExportService
VRSyncImportService
integration_mappings
snapshots de XML/payload
logs de integração
```

## Proibido

Salvar diretamente em entidades principais:

```text
For Sale
For Rent
Sale/Rent
Residential / Apartment
Pool
STANDARD
SECURITY_DEPOSIT
```

## Fluxo correto

```text
Domínio pt_BR
  ↓
Validação VRSync
  ↓
Mapper pt_BR -> VRSync
  ↓
Geração XML
  ↓
Snapshot/hash
  ↓
Publicação/feed
```

## Testes obrigatórios

- conversão correta;
- erro para mapeamento inexistente;
- validação de imóvel sem imagem;
- validação de endereço mínimo;
- validação de contato mínimo;
- garantia de que domínio continua em pt_BR.
