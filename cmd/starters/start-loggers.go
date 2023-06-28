package starters

import common "mongodb.com/common/application"

func StartLoggers() {
	common.NewErrorLog()
	common.NewRoutineLog()
	common.NewInfoLog()
}
