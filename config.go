package main

type Config struct {
	Args     []string
	Action   string
	TodoFile string
}

func NewConfig(_args []string, todofile string) Config {
	action, _ := GetArg(_args, 1)
	conf := Config{_args, action, todofile}
	return conf
}
