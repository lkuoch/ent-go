package resolvers

import (
	"lkuoch/ent-todo/ent/schema"
	"lkuoch/ent-todo/ent/schema/types/pulid"
	"sync"
)

type ResolverNameService struct{}

var (
	resolverMap    = make(map[string]string)
	configMapMutex sync.RWMutex
)

func (ResolverNameService) SetName(tableName string) {
	configMapMutex.Lock()
	key := pulid.NewPrefix(tableName)
	resolverMap[key] = tableName
	configMapMutex.Unlock()
}

func (ResolverNameService) GetName(tableName string) *string {
	value, ok := resolverMap[tableName]

	if !ok {
		return nil
	}

	return &value
}

func (r ResolverNameService) Init() {
	r.SetName(schema.User{}.Name())
	r.SetName(schema.Todo{}.Name())
}
