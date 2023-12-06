package transactions

import (
	"fmt"

	"github.com/hktrib/RetailGo/internal/ent"
)

func rollback(tx *ent.Tx, err error) error {

	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
		fmt.Println(err)
	}
	return err
}
