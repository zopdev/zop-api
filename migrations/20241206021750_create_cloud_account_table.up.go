package migrations

import "gofr.dev/pkg/gofr/migration"

func CreateTableObservabilityDistributors() migration.Migrate {
	return migration.Migrate{
		UP: CreateTableCloudAccountMigrateFunc,
	}
}

//nolint:gocritic //migration interface definition
func CreateTableCloudAccountMigrateFunc(d migration.Datasource) error {
	const query = `
CREATE TABLE if not exists cloud_account (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    provider TEXT NOT NULL,
    provider_id TEXT NOT NULL,
    provider_details TEXT,
    credentials TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);
`

	_, err := d.SQL.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
