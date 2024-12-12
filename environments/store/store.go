package store

import (
	"gofr.dev/pkg/gofr"
)

type Store struct {
}

func New() EnvironmentStore {
	return &Store{}
}

func (*Store) InsertEnvironment(ctx *gofr.Context, environment *Environment) (*Environment, error) {
	res, err := ctx.SQL.ExecContext(ctx, INSERTQUERY, environment.Name, environment.Level, environment.ApplicationID)
	if err != nil {
		return nil, err
	}

	environment.ID, _ = res.LastInsertId()

	return environment, nil
}

func (*Store) GetALLEnvironments(ctx *gofr.Context, applicationID int) ([]Environment, error) {
	environments := []Environment{}

	rows, err := ctx.SQL.QueryContext(ctx, GETALLQUERY, applicationID)
	if err != nil {
		return nil, err
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	for rows.Next() {
		var environment Environment

		err := rows.Scan(&environment.ID, &environment.Name, &environment.Level, &environment.ApplicationID,
			&environment.CreatedAt, &environment.UpdatedAt)
		if err != nil {
			return nil, err
		}

		environments = append(environments, environment)
	}

	return environments, nil
}

func (*Store) GetEnvironmentByName(ctx *gofr.Context, applicationID int, name string) (*Environment, error) {
	row := ctx.SQL.QueryRowContext(ctx, GETBYNAMEQUERY, name, applicationID)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var environment Environment

	err := row.Scan(&environment.ID, &environment.Name, &environment.Level, &environment.ApplicationID,
		&environment.CreatedAt, &environment.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &environment, nil
}

func (*Store) UpdateEnvironment(ctx *gofr.Context, environment *Environment) (*Environment, error) {
	_, err := ctx.SQL.ExecContext(ctx, UPDATEQUERY, environment.Name, environment.Level, environment.ID)
	if err != nil {
		return nil, err
	}

	return environment, nil
}
