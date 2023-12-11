package transactions

import (
	"context"
	"fmt"
	. "github.com/hktrib/RetailGo/cmd/api/stripe-components"

	"github.com/google/uuid"
	clerkstorage "github.com/hktrib/RetailGo/internal/clerk"
	"github.com/hktrib/RetailGo/internal/ent"
	"github.com/hktrib/RetailGo/internal/ent/usertostore"
)

func StoreAndOwnerCreationTx(ctx context.Context, reqStore *ent.Store, reqUser *ent.User, DB *ent.Client, clerkStore *clerkstorage.ClerkStorage) (error, *ent.Store, *ent.User) {
	tx, err := DB.Tx(ctx)
	if err != nil {
		return rollback(tx, fmt.Errorf("tx_error: starting a transaction: %w", err)), nil, nil
	}

	// Generating uuid for store and converting to string
	id := uuid.New().String()
	// Creating Stripe Connected Account
	stripeAccount, err := CreateConnectedAccount()
	if err != nil {
		return rollback(tx, fmt.Errorf("tx_error: Unable to create stripe account: %w", err)), nil, nil
	}

	newStore, err := tx.Store.Create().
		SetUUID(id).
		SetStoreName(reqStore.StoreName).
		SetOwnerEmail(reqStore.OwnerEmail).
		SetStoreAddress(reqStore.StoreAddress).
		SetStoreType(reqStore.StoreType).
		SetStripeAccountID(stripeAccount.ID).
		SetCreatedBy(reqUser.Email).
		Save(ctx)

	if err != nil {
		return rollback(tx, fmt.Errorf("tx_error: Unable to create store: %w", err)), nil, nil
	}

	newUser, err := tx.User.Create().
		SetClerkUserID(reqUser.ClerkUserID).
		SetEmail(reqUser.Email).
		SetIsOwner(true).
		SetFirstName(reqUser.FirstName).
		SetLastName(reqUser.LastName).
		SetStoreID(newStore.ID).AddStoreIDs(newStore.ID).
		Save(ctx)
	if err != nil {
		return rollback(tx, fmt.Errorf("tx_error: Unable to create owner: %w", err)), nil, nil
	}

	_, err = tx.UserToStore.Update().
		Where(usertostore.StoreID(newStore.ID), usertostore.UserID(newUser.ID)).
		SetClerkUserID(newUser.ClerkUserID).
		SetPermissionLevel(0).
		SetStoreName(newStore.StoreName).
		Save(ctx)
	if err != nil {
		return rollback(tx, fmt.Errorf("tx_error: Unable to Update UserToStore Junction table: %w", err)), nil, nil
	}

	newUTS, err := tx.UserToStore.Query().
		Where(usertostore.StoreID(newStore.ID), usertostore.UserID(newUser.ID)).
		Only(ctx)
	if err != nil {
		return rollback(tx, fmt.Errorf("tx_error: Unable to get query user to store table: %w", err)), nil, nil
	}
	err = clerkStore.AddMetadata(newUTS)
	if err != nil {
		return rollback(tx, fmt.Errorf("tx_error: Unable to user metadata to Clerk Store: %w", err)), nil, nil
	}

	return tx.Commit(), newStore, newUser
}
