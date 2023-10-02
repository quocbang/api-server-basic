package setupmock

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

// Mock is mock test database.
type Mock struct {
	DB    *gorm.DB
	Redis *redis.Client
}
