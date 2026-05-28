# 09 — Checklists de aceite por módulo

## 1. Checklist geral de tarefa

```text
[ ] A tarefa tem escopo pequeno e claro.
[ ] O agente não alterou arquivos fora do escopo sem justificar.
[ ] O código compila.
[ ] Testes focalizados passam.
[ ] Erros seguem o padrão do projeto.
[ ] Logs são estruturados.
[ ] Dados sensíveis não aparecem em logs.
[ ] Migrations são pequenas e revisáveis.
[ ] Queries sensíveis filtram tenant.
[ ] Autorização backend foi aplicada.
[ ] Valores internos estão em pt_BR.
[ ] VRSync não contaminou o domínio interno.
[ ] O resumo final explica como testar.
```

## 2. Tenants

```text
[ ] Tenant pode ser criado.
[ ] Tenant pode ser editado.
[ ] Tenant pode ser inativado/soft deleted.
[ ] Status usa pt_BR.
[ ] Tipo usa pt_BR.
[ ] Há auditoria para ações sensíveis.
[ ] Listagem tem paginação.
```

## 3. Usuários

```text
[ ] Usuário pode pertencer a múltiplos tenants.
[ ] Vínculo usuário-tenant tem status próprio.
[ ] Login não expõe senha/hash.
[ ] Token contém apenas dados necessários.
[ ] Refresh token é tratado com segurança.
```

## 4. RBAC

```text
[ ] Papéis são por tenant quando aplicável.
[ ] Permissões usam chaves em pt_BR técnico.
[ ] Authorization service centraliza decisão.
[ ] Handlers não duplicam regra complexa de permissão.
[ ] Testes cobrem acesso permitido e negado.
```

## 5. Relacionamentos entre tenants

```text
[ ] Tenant origem e destino existem.
[ ] Relacionamento tem tipo, status e permissões.
[ ] Relacionamento ativo sozinho não concede acesso.
[ ] Status suspenso/expirado/revogado bloqueia acesso.
[ ] Relações em grafo são suportadas sem hierarquia fixa.
```

## 6. Imóveis

```text
[ ] Imóvel sempre tem tenant_dono_id.
[ ] Status usa pt_BR.
[ ] Tipo de transação usa pt_BR.
[ ] Tipo de imóvel usa catálogo pt_BR.
[ ] Listagem filtra por tenant/autorização.
[ ] Endereço completo é salvo internamente.
[ ] Exibição de endereço é controlada.
[ ] Preços são normalizados.
```

## 7. Características e garantias

```text
[ ] Catálogos usam pt_BR.
[ ] Mapeamentos externos ficam em integration_mappings ou mapper.
[ ] Associação imóvel-característica funciona.
[ ] Garantias locatícias são salvas em pt_BR.
```

## 8. Mídias

```text
[ ] Arquivo vai para S3 compatível.
[ ] Banco salva metadados e storage_key.
[ ] Hash é calculado.
[ ] Apenas uma imagem principal por imóvel.
[ ] Ordem das imagens é preservada.
[ ] URL muda quando arquivo muda.
[ ] Imóvel exportável tem ao menos uma imagem.
```

## 9. VRSync

```text
[ ] VRSyncMapper cobre transação, publicação, tipo de imóvel, features e garantias.
[ ] VRSyncValidationService retorna erros antes da exportação.
[ ] VRSyncExportService gera XML válido.
[ ] Feed é gerado por tenant.
[ ] Snapshot/hash do XML é salvo.
[ ] Erros de integração são preservados.
[ ] Valores VRSync não são salvos no domínio principal.
```

## 10. Leads

```text
[ ] Lead pertence a tenant.
[ ] Lead pode estar associado a imóvel.
[ ] Telefone/e-mail são normalizados.
[ ] Duplicidade é detectada.
[ ] Consentimento LGPD é registrado quando aplicável.
[ ] Atribuição manual respeita permissão.
[ ] Histórico é registrado.
```

## 11. Comunicação

```text
[ ] Conversa vincula tenant, lead, imóvel e canal.
[ ] Mensagens têm direção, remetente, status e tipo.
[ ] Mensagens são append-only sempre que possível.
[ ] Payload externo bruto, se salvo, fica em JSONB e com cuidado LGPD.
```

## 12. Dashboard

```text
[ ] Métricas respeitam tenant.
[ ] Métricas têm filtros por período.
[ ] Indicadores de VRSync mostram erros.
[ ] Indicadores de leads mostram canal e status.
[ ] Queries agregadas são indexadas ou materializadas se necessário.
```

## 13. Auditoria

```text
[ ] Ações sensíveis são registradas.
[ ] Auditoria é append-only.
[ ] Registra user_id, tenant_id, entity_type, entity_id e ação.
[ ] Antes/depois usam JSONB quando necessário.
[ ] Logs não armazenam segredos.
```
