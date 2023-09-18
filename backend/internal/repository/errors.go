package repository

import "errors"

var (
	ErrorRecordNotFound      = errors.New("record not found")
	ErrorConstraintViolation = errors.New("constraint_violation")
	ErrorForeignKeyViolation = errors.New("foreign_key_violation")
)
