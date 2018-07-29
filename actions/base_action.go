package actions

type BaseAction interface {
	GetInfo() map[string]interface{}
	GetName() string
}
