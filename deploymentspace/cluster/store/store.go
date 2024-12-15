package store

import "gofr.dev/pkg/gofr"

type Store struct{}

func New() ClusterStore {
	return &Store{}
}

func (s *Store) Insert(ctx *gofr.Context, cluster *Cluster) (*Cluster, error) {
	res, err := ctx.SQL.ExecContext(ctx, INSERTQUERY, cluster.DeploymentSpaceID, cluster.Identifier,
		cluster.Name, cluster.Region, cluster.ProviderID, cluster.Provider, cluster.Namespace.Name)
	if err != nil {
		return nil, err
	}

	cluster.ID, err = res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return cluster, nil
}

func (s *Store) GetByDeploymentSpaceID(ctx *gofr.Context, deploymentSpaceID int) (*Cluster, error) {
	cluster := Cluster{}

	err := ctx.SQL.QueryRowContext(ctx, GETQUERY, deploymentSpaceID).Scan(&cluster.ID, &cluster.DeploymentSpaceID, &cluster.Identifier,
		&cluster.Name, &cluster.Region, &cluster.ProviderID, &cluster.Provider, &cluster.Namespace.Name, &cluster.CreatedAt, &cluster.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &cluster, nil
}
