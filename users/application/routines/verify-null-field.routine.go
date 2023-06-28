package users

import (
	common "mongodb.com/common/application"
	mongorepo "mongodb.com/common/infra/mongo"
	infra "mongodb.com/users/infra"
)

func VerifyNullFieldsRoutine() error {
	params := mongorepo.Params{
		Filter: []mongorepo.Filter{
			{Column: "name", Value: "", Type: mongorepo.OPERATOR, Operator: mongorepo.IS_NOT_FILLED},
			{Column: "age", Value: "", Type: mongorepo.OPERATOR, Operator: mongorepo.IS_NOT_FILLED},
		},
		Limit: 100000,
	}

	users := infra.UserRepository.List(params)

	for _, user := range users.Items {
		common.RoutineLog.Printf("User of id %s haves a null field, name: %s, age: %v", user.Id, user.Name, user.Age)
	}

	return nil

}
