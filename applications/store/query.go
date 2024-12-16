package store

const (
	INSERTQUERY    = "INSERT INTO application ( name) VALUES ( ?);"
	GETALLQUERY    = "SELECT id, name, created_at, updated_at FROM application WHERE deleted_at IS NULL;"
	GETBYNAMEQUERY = "SELECT id, name, created_at, updated_at FROM application WHERE name = ? and deleted_at IS NULL;"
)
