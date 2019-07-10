package etcdv3

// Key struct
type Key struct {
	ID     string
	Parent string
	Text   string
}

// Path struct
type Path struct {
	key   string
	Value string
}
