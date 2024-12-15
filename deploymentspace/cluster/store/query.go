package store

const (
	INSERTQUERY = "INSERT INTO cluster (deployment_space_id,cluster_id,name, region,provider_id," +
		"provider,namespace) VALUES ( ?, ?, ?, ?, ?, ?, ?);"
	GETQUERY = "SELECT id, deployment_space_id, cluster_id, name, region, provider_id, provider, " +
		"namespace, created_at, updated_at FROM cluster WHERE deployment_space_id = ? and deleted_at IS NULL;"
)
