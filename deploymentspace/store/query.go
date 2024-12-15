package store

const (
	INSERTQUERY     = "INSERT INTO deployment_space (cloud_account_id, environment_id, type) VALUES ( ?, ?, ?);"
	GETQUERYBYENVID = "SELECT id, cloud_account_id, environment_id, type, created_at, updated_at FROM deployment_space WHERE environment_id = ? and deleted_at IS NULL;"
)
