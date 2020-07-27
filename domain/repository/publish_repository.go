package repository

import (
	"github.com/biosugar0/postmaster/domain/entity"
)

type PublishRepository interface {
	Create(cond entity.Publish) (string, error)
}
