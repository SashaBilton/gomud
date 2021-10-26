package space

type Location struct {
	Desc  string
	Exits map[string]*Exit
}

func (location *Location) AddExit(exit *Exit) {
	if location.Exits == nil {
		location.Exits = make(map[string]*Exit)
	}
	location.Exits[exit.Name] = exit
}
