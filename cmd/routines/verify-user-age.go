package routines

import (
	"time"

	common "mongodb.com/common/application"
	users "mongodb.com/users/application/routines"
)

func VerifyUserAge() {

	for {
		common.RoutineLog.Println("Verify user age routine started")

		err := users.VerifyUserAgeRoutine()

		if err != nil {
			common.RoutineLog.Println("Error in routine verify user age", err)
		}

		common.RoutineLog.Println("Verify user age routine ended")

		time.Sleep(30 * time.Second)

	}
}
