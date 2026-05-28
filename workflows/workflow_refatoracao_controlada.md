# Workflow — Refatoração controlada

## Quando usar

Use apenas quando a tarefa pedir refatoração ou quando a implementação estiver bloqueada por dívida técnica local.

## Regras

1. Não alterar comportamento externo.
2. Não alterar contrato da API.
3. Não alterar migrations existentes sem justificativa.
4. Não mover arquivos de vários módulos ao mesmo tempo.
5. Criar ou manter testes antes/depois.

## Passos

1. Descrever problema local.
2. Listar arquivos afetados.
3. Definir testes de segurança.
4. Refatorar em pequenos commits/diffs.
5. Rodar testes.
6. Informar risco residual.
