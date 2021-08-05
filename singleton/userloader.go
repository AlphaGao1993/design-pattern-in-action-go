package singleton

import (
	"sync"
	"time"

	"pattern/singleton/model"
)

// userOnce 用于保证单例只会被初始化一次
// userLoader 就是最核心的 UserLoader 的唯一实例
var (
	userOnce   sync.Once
	userLoader *UserLoader
)

func GetUserLoader(fetch func(id uint64) (model.User, error)) *UserLoader {
	userOnce.Do(func() {
		userLoader = NewUserLoader(UserLoaderConfig{
			Fetch: func(keys []uint64) ([]model.User, []error) {
				users := make([]model.User, len(keys))
				errs := make([]error, len(keys))
				wg := sync.WaitGroup{}
				wg.Add(len(keys))
				for index, key := range keys {
					// 为了保证积压的请求能够迅速返回结果，这里用并发的方式来进行
					// 如果 GetUserByID 是通过查询 db 或者 http 等接口实现，则并发会缩短 Batch 的响应时间
					go func(i int, k uint64) {
						defer wg.Done()
						users[i], errs[i] = fetch(k)
					}(index, key)
				}
				wg.Wait()
				return users, errs
			},
			Wait:     time.Millisecond * 20,
			MaxBatch: 50,
		})
	})
	return userLoader
}
