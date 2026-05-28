package domain

import "errors"

var (
	ErrInvalidInput       = errors.New("entrada invalida")
	ErrNotFound           = errors.New("nao encontrado")
	ErrInvalidCredentials = errors.New("credenciais invalidas")
	ErrUnauthorized       = errors.New("nao autorizado")
)
