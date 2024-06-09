package pgsqlproduct

import (
	"clean-code-structure/entity"
	"clean-code-structure/param/productparam"
	"clean-code-structure/pkg/errmsg"
	"clean-code-structure/pkg/richerror"
	"context"
	"database/sql"
	"fmt"
	"time"
)

func (d *DB) GetListProducts(ctx context.Context, page int, prePage int, filters ...entity.Filter) (products []productparam.ProductRepo, hasMore bool, err error) {
	const op = "pgsqlproduct.GetListProducts"
	start := (page - 1) * prePage
	query, values := d.conn.MakeQueryWithFilters(filters)
	if query != "" {
		query = " WHERE " + query
	}
	query = query + fmt.Sprintf(` OFFSET $%d LIMIT $%d `, len(values)+1, len(values)+2)
	values = append(values, start)
	values = append(values, prePage)
	rows, err := d.conn.Conn().QueryContext(ctx, `SELECT * FROM products `+query, values...)

	if err != nil {
		return nil, false, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}

	products, err = scanProducts(rows)
	if err != nil {
		return nil, false, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}

	total, err := d.GetTotalProductCount(ctx, filters...)
	if err != nil {
		return nil, false, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}
	lastResultNumber := prePage * page

	hasMore = total > lastResultNumber

	return products, hasMore, nil
}

func scanProducts(scanner *sql.Rows) ([]productparam.ProductRepo, error) {
	var products []productparam.ProductRepo

	for scanner.Next() {
		var createdAt time.Time
		var updatedAt time.Time
		var deletedAt sql.NullTime
		var product entity.Product

		err := scanner.Scan(&product.Id, &product.Name, &product.Quantity, &createdAt, &updatedAt, &deletedAt)

		if err != nil {

			return nil, err
		}
		products = append(products, productparam.ProductRepo{
			Entity:    product,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			DeletedAt: deletedAt,
		})
	}

	return products, nil
}
