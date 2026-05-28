# React Frontend Rules

## Stack

- React.
- TypeScript.
- TanStack Query.
- React Hook Form.
- Zod.
- React Router ou roteador definido no projeto.
- UI com MUI ou Tailwind/Shadcn conforme decisão do projeto.

## Regras

1. Usar componentes funcionais.
2. Tipar props e respostas da API.
3. Não duplicar regra de negócio sensível no frontend.
4. Frontend pode ocultar ações, mas backend decide autorização.
5. Toda tela de listagem deve ter loading, erro, vazio e paginação.
6. Formulários devem validar com Zod.
7. Usar TanStack Query para cache e invalidação.
8. Não fazer chamadas diretas sem camada `api.ts`/hooks.

## Estrutura por feature

```text
src/features/[feature]/
├── api.ts
├── hooks.ts
├── schemas.ts
├── types.ts
├── pages/
└── components/
```

## Performance

- Debounce em filtros textuais.
- Paginação server-side.
- Evitar renderizar listas grandes sem paginação ou virtualização.
