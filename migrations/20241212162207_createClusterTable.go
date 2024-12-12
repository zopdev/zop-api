package migrations

import (
	"gofr.dev/pkg/gofr/migration"
)

func createClusterTable() migration.Migrate {
	return migration.Migrate{
		UP: func(d migration.Datasource) error {
			const query = `
				CREATE TABLE if not exists cluster (
					id INTEGER PRIMARY KEY AUTOINCREMENT,
					type VARCHAR(255) NOT NULL,
				    environment_id INTEGER NOT NULL,
				    cloudaccount_id INTEGER NOT NULL,
				    details TEXT NOT NULL,
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
		},
	}
}
