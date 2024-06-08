package productparam

import (
	"clean-code-structure/entity"
	"database/sql"
	"time"
)

type ProductRepo struct {
	Entity    entity.Product
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
