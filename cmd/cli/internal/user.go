package internal

import (
	"context"
	"github.com/hvturingga/ya/ent"
)

func GetUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	query, err := client.User.Query().WithProvider().WithSubscribe().WithDaemon().First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			create := client.User.Create().SaveX(ctx)
			return create, nil
		}
		return nil, err
	}
	return query, nil
}
