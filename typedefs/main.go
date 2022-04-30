package typedefs

import "gotu/utils"

type ActionFunc = func(args []string)

type CmdModule struct {
	Cmd         string
	Description string
	HelpText    string
	Actions     map[string]Action
}

type Action struct {
	Name     string
	HelpText string
}

func (cmd *CmdModule) AddAction(k string, helpText string) {
	cmd.Actions[k] = Action{Name: k, HelpText: helpText}
}

func (cmd *CmdModule) PrintHelp() {
	utils.PrintHelp(cmd.Description, cmd.HelpText)
}
