package pm

// IProcess ... process interface
type IProcess interface {
	Start() error
	Stop() error
}
