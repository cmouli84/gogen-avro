package schema

type Node interface {
	Name() string
	Package() string
	Children() []AvroType
}
