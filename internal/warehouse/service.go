package warehouse

// Service encapsulates the business logic of a Warehouse.
// As stated by this principle https://golang.org/doc/effective_go#generality,
// since the underlying concrete implementation does not export any other method that is not in the interface,
// we decided to define it where it is implemented rather where it is used (commonly in a handler).
type Service interface {
}

type service struct {
}

func NewService() Service {
	return &service{}
}
