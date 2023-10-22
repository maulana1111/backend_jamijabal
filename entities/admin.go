package entities

type Admin struct {
	Id        uint
	Name      string
	Username  string
	Password  string
	RoleAdmin RoleAdmin
	Photo     string
	Status    string
}
