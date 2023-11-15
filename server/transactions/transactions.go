package transactions

import (
	"context"
	"fmt"

	"github.com/hktrib/RetailGo/ent"
)

type StoreKey struct {
	Store string
}

type OwnerKey struct {
	Owner string
}

func StoreAndOwnerCreationTx(ctx context.Context, dbClient *ent.Client) error {
	tx, err := dbClient.Tx(ctx)
	if err != nil {
		return rollback(tx, fmt.Errorf("tx_error: starting a transaction: %w", err))
	}

	reqStore := ctx.Value(StoreKey{"store"}).(*ent.Store)
	reqUser := ctx.Value(OwnerKey{"owner"}).(*ent.User)

	_, err = tx.Store.Create().SetStoreName(reqStore.StoreName).Save(ctx)

	if err != nil {
		return rollback(tx, fmt.Errorf("tx_error: Unable to create owner: %w", err))
	}
	_, err = tx.User.Create().
		SetEmail(reqUser.Email).
		SetIsOwner(reqUser.IsOwner).
		SetRealName(reqUser.RealName).
		SetStoreID(reqUser.StoreID).Save(ctx)

	if err != nil {
		return rollback(tx, fmt.Errorf("tx_error: Unable to create store: %w", err))
	}
	return tx.Commit()
}

func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}