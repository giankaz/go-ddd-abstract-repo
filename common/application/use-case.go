package common

type UseCase[Input interface{}, Output interface{}] interface {
	execute(input Input) (Output, error)
}
