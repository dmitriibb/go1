package context

type Bean struct {
	Metadata *BeanMetadata
	Instance *any
}
type BeanStatus string

const (
	Created BeanStatus = "created"
	Init    BeanStatus = "init"
	Ready   BeanStatus = "ready"
	Dead    BeanStatus = "dead"
)

type BeanMetadata struct {
	typeName string
	status   BeanStatus
}
