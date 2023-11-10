package kvRedis

import "github.com/hktrib/RetailGo/ent"

type UserCache interface {
	Set(key string, entity *ent.User)
	Get(key string) *ent.User
}
