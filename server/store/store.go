package store

import (
	"backup/server/entities"
)

type Store interface {
	GetForms() ([]*entities.Form, error)
	SaveForm(form *entities.Form) error
}
