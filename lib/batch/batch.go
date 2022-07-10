package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

type UsersList struct {
	data          []user
	mx            sync.Mutex
	currentUserId int64
}

func (u *UsersList) Init(capacity int64) {
	u.data = make([]user, 0, capacity)
	u.currentUserId = 0
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	u := new(UsersList)
	u.Init(n)

	var wg sync.WaitGroup

	batchAmount := n / pool

	for i := 0; i < int(pool); i++ {
		wg.Add(1)
		go func(amount int64) {
			defer wg.Done()
			for i := 0; i < int(amount); i++ {
				u.mx.Lock()
				currentUserId := u.currentUserId
				u.currentUserId += 1
				u.mx.Unlock()
				nextUser := getOne(currentUserId)
				u.mx.Lock()
				u.data = append(u.data, nextUser)
				u.mx.Unlock()
			}
		}(batchAmount)
	}
	wg.Wait()
	return u.data
}
