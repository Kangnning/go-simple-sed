package sed

type Mode uint8
type Option uint8

const (
	InsertBefore Option = iota
	InsertAfter
	Delete
	Replace
)

type Config struct {
	FileName  string
	Opt       Option
	Pattern   string
	DesString string
}
