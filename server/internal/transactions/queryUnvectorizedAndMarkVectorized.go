package transactions

import (
	"context"
	"fmt"

	"github.com/hktrib/RetailGo/internal/ent"
	"github.com/hktrib/RetailGo/internal/ent/item"
)

func QueryUnvectorizedItemsAndMarkVectorized(ctx context.Context, dbClient *ent.Client) ([]*ent.Item, error) {
	tx, err := dbClient.Tx(ctx)
	if err != nil {
		return nil, rollback(tx, fmt.Errorf("tx_error: starting a transaction: %w", err))
	}

	unvectorizedItems, err := tx.Item.Query().Where(item.Vectorized(false)).All(ctx)

	if err != nil {
		return nil, err
	}

	tx.Item.Update().Where(item.Vectorized(false)).SetVectorized(true).Save(ctx)

	tx.Commit()

	return unvectorizedItems, nil
}
