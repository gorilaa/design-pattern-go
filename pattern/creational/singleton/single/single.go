package single

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct{}

var singleInstance *single
var once sync.Once

func GetInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Create single instance now")
			once.Do(func() {
				singleInstance = &single{}
			})
		} else {
			fmt.Println("Single Instance already created.")
		}
	} else {
		fmt.Println("Single Instance already created.")
	}

	return singleInstance
}
