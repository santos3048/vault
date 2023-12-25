package help

import (
	"fmt"
	"vault/color"
)



func PrintFullHelp(){
	fmt.Println(`
Available commands:
	add [service] [username] [password] - Adds a new password to your database
	get [service] [username?] - Gets a password. If no username is given, will print all available passwords
	delete [service] [username] - Deletes a password
	generate [service] [username] [length] - Generates a new password of certain length and adds it to your database
	exit - Exits vault		`)
}
func PrintHelp(command string){
	switch command {
	case "add":
		fmt.Println("The command <add> is used to add a new password to your database. Usage: ")
		color.Set("cyan")
		fmt.Println("add [service] [username] [password]")
		color.Unset()
	case "get":
		fmt.Println("The command <get> is used to get a password from your database. If no username is given, it will print all password registered in a service. Usage: ")
		color.Set("cyan")
		fmt.Println("get [service] [username?]")
		color.Unset()
	
	case "exit":
		fmt.Println("The command <exit> is used to exit the Vault")
	case "generate":
		fmt.Println("The command <generate> is used to generate a password of certain length. Usage: ")
		color.Set("cyan")
		fmt.Println("generate [service] [username] [length]")
		color.Unset()
	case "delete":
		fmt.Println("The command <delete> is used to delete a password. Usage: ")
		color.Set("cyan")
		fmt.Println("delete [service] [username?]")
		color.Unset()
	default:
		fmt.Println("There is no such a command. Type 'help' to have a list of the available commands.")
	}
}