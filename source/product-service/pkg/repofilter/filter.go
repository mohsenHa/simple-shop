package repofilter

type Operator string

type Filter struct {
	name     string
	operator Operator
	value    string
}

func New(name string, operator Operator, value string) Filter {
	return Filter{
		name:     name,
		operator: operator,
		value:    value,
	}
}

func (f Filter) GetValue() string {
	return f.value
}

func (f Filter) GetOperator() string {
	return string(f.operator)
}

func (f Filter) GetName() string {
	return f.name
}

const (
	CompareTypeEqual Operator = "="
	CompareTypeLike  Operator = "like"
)
