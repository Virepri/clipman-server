package commands

type Command func(args []string)

var Aliases = map[string]Command{
	"help":   help,
	"config": config,
	"exit":   exit,
	"user":   user,
	"debug":  debug,
}

var HelpList = map[string]string{
	"help": `displays this text`,
	"config": `config save:
saves the config.

config reload:
reloads the config, to undo changes. This is not equivalent to a server restart. 

config buffer [size]:
changes the buffer size.

config bind [port]:
changes where the server listens on.

config password [password]:
changes the admin password stored.`,
	"exit": `Shuts down the server.`,
	"user": `user save:
Saves user info.

user password [password]:
Changes the stored password for the user.
`,
	"debug":`debug is a TEMPORARY utility.

debug clipboard:
Print the current clipboard content

debug devices:
Print the connected devices.

debug printcmd:
Toggle the option to print commands issued by clients.`,
}
