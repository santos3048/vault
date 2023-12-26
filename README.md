### Vault v1.1.5

Vault is a simple password manager. It uses AES-GCM encryption to store passwords.

### Using
Run ```vault``` to start. Available commands:
  ```
  add [service] [username] [password] - Adds a new password to your database
	get [service] [username?] - Gets a password. If no username is given, will print all available passwords
	delete [service] [username] - Deletes a password
	generate [service] [username] [length] - Generates a new password of certain length and adds it to your database
	exit - Exits vault
  *show - Prints the array containing all usernames and passwords
 ```

### Autoupdates
If autoupdates are failing, run 'update.bat'

