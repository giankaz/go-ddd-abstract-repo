package starters

import "mongodb.com/cmd/routines"

func StartRoutines() {
	/*go*/ routines.VerifyUserAge()
	/*go*/ routines.VerifyNullFields()
}
