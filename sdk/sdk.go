package sdk

type User struct {
	Id   string
	name string
}

type Sdk struct {
	Id      string
	UserMap map[string]User
}

func NewSdk(id string) (sdk Sdk) {
	return Sdk{id, make(map[string]User)}
}

func (s *Sdk) Run() {

}
