package constant


const (
	ProdScope Scope = "prod"
	TestScope Scope = "test"
	DevScope Scope = "dev"
)

type Scope string

func (s Scope) Value() string {
	return string(s)
}