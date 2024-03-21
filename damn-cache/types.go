//@Author: wulinlin
//@Description:
//@File:  types
//@Version: 1.0.0
//@Date: 2024/03/21 20:21

package damn_cache

import (
	"context"
	"time"
)

type Cache interface {
	Get(ctx context.Context, key string) (any, error)

	Set(ctx context.Context, key string, value any, expire time.Duration) error

	Delete(ctx context.Context, key string) error

	Clear(ctx context.Context)
}
