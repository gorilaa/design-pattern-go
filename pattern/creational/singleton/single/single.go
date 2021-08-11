package single

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct{}

var singleInstance *single

func GetInstance() interface{} {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Create single instance now")
			singleInstance = &single{}
		} else {
			fmt.Println("Single Instance already created.")
		}
	} else {
		fmt.Println("Single Instance already created.")
	}

	return singleInstance
}
