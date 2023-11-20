package kv

import (
	"time"

	"github.com/hktrib/RetailGo/internal/ent"
)

type UserCache interface {
	Set(key string, value *ent.User)
	SetX(key string, value *ent.User, expiresAt time.Duration)
	Get(key string) *ent.User
}

type KeyTypes string

const (
	OwnerCreatingStore KeyTypes = "OwnerCreatingStore"
)
