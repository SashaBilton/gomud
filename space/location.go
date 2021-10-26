package space

type Location struct {
	Desc  string
	Exits *Exit
}

func (location Location) AddExit(exit *Exit) {
	location.Exits = exit
}
