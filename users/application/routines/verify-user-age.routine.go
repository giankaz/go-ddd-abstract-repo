package users

import (
	common "mongodb.com/common/application"
	mongorepo "mongodb.com/common/infra/mongo"
	infra "mongodb.com/users/infra"
)

func VerifyUserAgeRoutine() error {
	params := mongorepo.Params{
		Filter: []mongorepo.Filter{
			{Column: "age", Value: "18", Type: mongorepo.OPERATOR, Operator: mongorepo.IS_LESS_THAN},
		},
		Limit: 100000,
	}

	users := infra.UserRepository.List(params)

	for _, user := range users.Items {
		common.RoutineLog.Printf("User of id %s and name %s is not an adult, the age is %v", user.Id, user.Name, user.Age)
	}

	return nil

}
