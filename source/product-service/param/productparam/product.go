package productparam

import (
	"clean-code-structure/entity"
	"database/sql"
	"encoding/json"
	"time"
)

type ProductRepo struct {
	Entity    entity.Product `json:"entity"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt sql.NullTime   `json:"deleted_at"`
}

type productJson struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Quantity  int       `json:"quantity"`
	IsDeleted bool      `json:"is_deleted"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (p *ProductRepo) MarshalJSON() ([]byte, error) {
	isDeleted := p.DeletedAt.Valid
	deletedAt := p.DeletedAt.Time

	return json.Marshal(&productJson{
		Id:        p.Entity.Id,
		Name:      p.Entity.Name,
		Quantity:  p.Entity.Quantity,
		IsDeleted: isDeleted,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		DeletedAt: deletedAt,
	})
}
