package businesslogic

import (
	"fmt"
	"github.com/fdaines/arch-go-sample-project/src/persistence"
)

func GetProcessedData() string {
	dataFromDB := persistence.RetrieveDataFromDatabase()
	return fmt.Sprintf("[%s]", dataFromDB)
}
