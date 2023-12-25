package main

import (
	"bufio"
	"fmt"
	"gitlab.com/david_mbuvi/go_asterisks"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"vault/color"
	"vault/encryption"
	"vault/help"
	"vault/utils"
)

/*

Add auto update
https://learn.microsoft.com/en-us/windows/win32/taskschd/using-the-task-scheduler


*/

const VERSION = "v1.1.5"
const PREFACE = "abstractionisms\n" // 16 bytes -- minimal block sizefor AES

func main() {

	ex, _ := os.Executable()
	color.Init()

	// Print Vault ${version}
	color.Set("bold")
	fmt.Println("Vault ", VERSION)
	color.Unset()

	response, error := http.Get("https://raw.githubusercontent.com/santos3048/vault/main/version.txt")
	if error != nil {
		color.Set("yellow")
		fmt.Println("Could not check for updates")
		color.Unset()
	} else {
		body, error := ioutil.ReadAll(response.Body)
		if error != nil {
			color.Set("yellow")
			fmt.Println("Could not check for updates")
			color.Unset()
		} else {
			response.Body.Close()

			newv := strings.TrimSpace(string(body))
			//	fmt.Print(newv, VERSION)

			if newv != VERSION {
				//		fmt.Print(string(body), VERSION)
				fmt.Print("New version available: ")
				color.Set("cyan")
				fmt.Println(string(body))
				color.Unset()
				err := utils.DownloadFile(filepath.Dir(ex)+"\\new.exe", "https://raw.githubusercontent.com/santos3048/vault/main/vault"+newv+".exe")
				if err != nil {
					color.Set("yellow")
					fmt.Println("Could not download update")
					color.Unset()
				} else {
					fmt.Println("Run 'update.bat' on the command line")
					os.Exit(0)
				}
			} else {
				color.Set("green")
				fmt.Println("Vault is up to date")
				color.Unset()
			}
		}
	}

	reader := bufio.NewReader(os.Stdin)                //Reader to get the commands
	dat, err := os.ReadFile(filepath.Dir(ex) + "\\db") //Open file

	if err != nil { // File does not exist (Supposedly). Therefore there is no password
		f, _ := os.Create(filepath.Dir(ex) + "\\db") // Create file
		fmt.Print("Welcome to Vault! Create a password: ")
		pwd, _ := reader.ReadString('\n')
		pwd = strings.TrimSpace(pwd)
		pwd = utils.PadRight(pwd, 32) // Pad to 32
		encrypted, err := encryption.EncryptMessage([]byte(pwd), PREFACE)
		utils.Check(err)
		f.WriteString(encrypted)
		fmt.Println("Success, just remember your password!")
		f.Close()
	}

	dat, err = os.ReadFile(filepath.Dir(ex) + "\\db") // read file (encrypted)
	utils.Check(err)                                  // check any errors
	fmt.Println("Welcome to Vault ", VERSION)
	txt := ""
	password := ""
	for {
		fmt.Print("Insert your password: ")
		pwd, err := go_asterisks.GetUsersPassword("", true, os.Stdin, os.Stdout)
		password = string(pwd)
		password = utils.PadRight(password, 32) // pad string until it has sufficient length
		utils.Check(err)
		txt, err = encryption.DecryptMessage([]byte(password), string(dat)) // Try to decrypt the file
		utils.Check(err)

		if !strings.HasPrefix(txt, PREFACE) { // Decryption failed
			color.Set("red")
			fmt.Println("Wrong password.")
			color.Unset()
		} else {
			break
		}
	}

	arr := strings.Split(txt, "\n")

repl:
	for { // Infinite REPL
		fmt.Print(">>> ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if strings.HasPrefix(text, "get") {
			dt := strings.Split(text, " ")
			if len(dt) == 1 { // User just typed get
				help.PrintHelp("get")
				continue repl
			} else if len(dt) == 2 { // get <service>
				searchString := dt[1] + ":"
				found := false
				for _, s := range arr {
					if strings.HasPrefix(s, searchString) {
						fmt.Print("Found password: ")
						color.Set("cyan")
						fmt.Println(strings.Join(strings.Split(strings.Replace(s, searchString, "", 1), ":"), " --> "))
						color.Unset()
						found = true
					}
				}
				if !found {
					fmt.Println("No matching passwords found. You can add passwords with the 'add' command")
				}
				continue repl
			}
			site := dt[1]
			username := dt[2]
			searchString := site + ":" + username + ":" // full query
			for _, s := range arr {
				if strings.HasPrefix(s, searchString) {
					fmt.Print("The password is: ")
					color.Set("cyan")
					fmt.Println(strings.Replace(s, searchString, "", 1))
					color.Unset()
					continue repl
				}
			}
			fmt.Println("Password not found. You can add passwords with the 'add' command")
			continue repl
		}

		if strings.HasPrefix(text, "add") {
			dt := strings.Split(text, " ")
			if len(dt) != 4 {
				help.PrintHelp("add")
				continue repl
			}
			site := dt[1]
			username := dt[2]
			password := dt[3]
			addString := site + ":" + username + ":" + password
			arr = append(arr, addString)
			fmt.Println("Password added succesfully.")
			continue repl
		}

		if strings.HasPrefix(text, "generate") {
			dt := strings.Split(text, " ")
			if len(dt) != 4 {
				help.PrintHelp("generate")
				continue repl
			}
			site := dt[1]
			username := dt[2]
			leng, err := strconv.Atoi(dt[3])
			utils.Check(err)
			password := utils.Generate(leng)
			addString := site + ":" + username + ":" + password
			arr = append(arr, addString)
			fmt.Println("Password " + password + " added succesfully.")
			continue repl
		}

		if strings.HasPrefix(text, "delete") {
			dt := strings.Split(text, " ")
			if len(dt) != 3 {
				help.PrintHelp("delete")
				continue repl
			}
			site := dt[1]
			username := dt[2]
			delString := site + ":" + username + ":"
			for i, s := range arr {
				if strings.HasPrefix(s, delString) {
					arr = utils.Remove(arr, i)
					continue repl
				}
			}
			continue repl
		}

		if strings.HasPrefix(text, "help") {
			dt := strings.Split(text, " ")
			if len(dt) == 1 {
				help.PrintFullHelp()
			} else {
				help.PrintHelp(dt[1])
			}
			continue repl
		}

		switch text {
		case "show":
			fmt.Println("{#?}", arr)
		case "exit":
			fmt.Print("Exiting...")
			os.Truncate(filepath.Dir(ex)+"\\db", 0)
			encrypted, err := encryption.EncryptMessage([]byte(password), strings.Join(arr, "\n"))
			utils.Check(err)
			os.WriteFile(filepath.Dir(ex)+"\\db", []byte(encrypted), 0644)
			os.Exit(0)

		default:
			fmt.Println("Command not recognized. Type 'help' to get a list of the available commands")
			//fmt.Println(text)
		}
	}
}
