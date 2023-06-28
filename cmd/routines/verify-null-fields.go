package routines

import (
	"time"

	common "mongodb.com/common/application"
	users "mongodb.com/users/application/routines"
)

func VerifyNullFields() {

	for {
		common.RoutineLog.Println("verify null fields routine started")

		err := users.VerifyNullFieldsRoutine()

		if err != nil {
			common.RoutineLog.Println("Error in routine verify null fields", err)
		}

		common.RoutineLog.Println("verify null fields routine ended")

		time.Sleep(20 * time.Second)

	}
}
