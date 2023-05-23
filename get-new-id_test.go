package sqlite

import (
	"sync"
	"testing"
)

func Test_GetNewID(t *testing.T) {
	idRequired := 30
	wg := sync.WaitGroup{}
	wg.Add(idRequired)

	db := NewConnection("./test_files/", "test.db", false)

	idObtained := make(map[string]int)
	var esperar sync.Mutex

	for i := 0; i < idRequired; i++ {
		go func() {
			defer wg.Done()
			id := db.GetNewID()
			esperar.Lock()
			if cantId, exist := idObtained[id]; exist {
				idObtained[id] = cantId + 1
			} else {
				idObtained[id] = 1
			}
			esperar.Unlock()

		}()
	}
	wg.Wait()

	// fmt.Printf("total id requeridos: %v ob: %v\n", idRequired, len(idObtained))
	// fmt.Printf("%v", idObtained)
	if idRequired != len(idObtained) {
		t.Fail()
	}
}
