package storage

type (
	DB struct {
		data map[string]bool
	}
)

func NewDb(data map[string]bool) *DB {
	return &DB{
		data: data,
	}
}

func (d *DB) Get(key string) bool {
	_, ok := d.data[key]
	return ok
}

func (d *DB) Add(key string) bool {
	d.data[key] = true
	return d.data[key]
}
