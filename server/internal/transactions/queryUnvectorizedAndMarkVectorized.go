package transactions

import (
	"context"
	"fmt"

	"github.com/hktrib/RetailGo/internal/ent"
	"github.com/hktrib/RetailGo/internal/ent/item"
)

func QueryUnvectorizedItemsAndMarkVectorized(ctx context.Context, dbClient *ent.Client) ([]ent.Item, error) {
	tx, err := dbClient.Tx(ctx)
	if err != nil {
		return nil, rollback(tx, fmt.Errorf("tx_error: starting a transaction: %w", err))
	}

	unvectorizedItems, err := tx.Item.Query().Where(item.Vectorized(false)).All(ctx)

	if err != nil {
		return nil, rollback(tx, fmt.Errorf("tx_error: querying unvectorized items: %w", err))
	}

	previouslyUnvectorizedItems := []ent.Item{}

	for _, itemPtr := range unvectorizedItems {
		previouslyUnvectorizedItems = append(previouslyUnvectorizedItems, *itemPtr)
	}

	fmt.Println("Previously Unvectorized Items:", previouslyUnvectorizedItems)

	_, err = tx.Item.Update().Where(item.Vectorized(false)).SetVectorized(true).Save(ctx)

	if err != nil {
		return nil, rollback(tx, fmt.Errorf("tx_error: marking items vectorized: %w", err))
	}

	// dbClient.Item.Update().Where(item.Vectorized(false)).SetVectorized(true).Save(ctx)

	tx.Commit()

	return previouslyUnvectorizedItems, nil
}
