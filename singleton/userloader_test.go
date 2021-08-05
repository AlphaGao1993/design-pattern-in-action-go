package singleton

import (
	"fmt"
	"pattern/singleton/model"
	"testing"
)

func TestSingleton(t *testing.T) {
	instanceA := GetUserLoader(GetUserByID)
	instanceB := GetUserLoader(GetUserByID)
	if instanceA != instanceB {
		t.Fatal("two singleton not equal")
	}
}

func TestLoadUsers(t *testing.T) {
	// 执行改测试，因为 id 默认全是 0，所以按照 key 去重后真正发出的请求
	// 只有一次，那么在控制台里将只能看到打印出 1 次 "fetch user"，所用时间则为 0.02s
	// 因为没有打包预定的 50 个请求（ 去重后只有 1 个），大大减少 IO 压力
	userIDs := make([]uint64, 50)
	for i := range userIDs {
		GetUserLoader(GetUserByID).Load(userIDs[i])
	}
}

func GetUserByID(id uint64) (model.User, error) {
	// 这里应该是通过查询数据库或者查询网络来实现
	fmt.Println("fetch user")
	return model.User{
		ID:     id,
		Name:   "alpha gao",
		Avatar: "https://avatars.githubusercontent.com/u/18735052?s=400&u=3fe7c57771ac26c0726f6d31ce2c2c711b04cdbd&v=4",
	}, nil
}
