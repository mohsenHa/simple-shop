package pgsql

import (
	"clean-code-structure/entity"
	"fmt"
	"strings"
)

func (p *PGSQLDB) MakeQueryWithFilters(filters []entity.Filter) (string, []any) {
	query := make([]string, len(filters))
	values := make([]any, len(filters))
	for id, filter := range filters {
		values = append(values, filter.GetValue())
		query = append(query, fmt.Sprintf("%s %s $%d", filter.GetName(), filter.GetOperator(), id+1))
	}

	return strings.Join(query, " AND "), values
}
