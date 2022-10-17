package domain

type Core struct {
	ID       uint
	Name     string
	Email    string
	Phone    string
	Addres   string
	Bio      string
	Gender   string
	Location string
}

type Repository interface {
	Insert(newUser Core) (Core, error)  //register
	GetUser(Email string) (Core, error) //login
	Update(updatedData Core, ID uint) (Core, error)
	Get(ID uint) (Core, error)
	Delete(ID uint) (Core, error)
}

type Service interface {
	Register(newUser Core) (Core, error)
	Login(Email string) (Core, error)
	UpdateProfile(updatedData Core, ID uint) (Core, error)
	Profile(ID uint) (Core, error)
	DeleteProfile(ID uint) (Core, error)
}
