package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"log"
	"strings"
)

func main() {
	SaveExtensions("lol.txt");
	
	CheckIfCodeInPath();
	
	var code int = 0;
	
	for {
		fmt.Println("Install from desktop - 1\nSave to desktop - 2");
		_, err := fmt.Scanf("%d", &code);
		var wrongInput bool = err != nil || (code != 1 && code != 2);

		if wrongInput {
			fmt.Println("\nPlease, input only integer (1 or 2)")
			continue;
		} else {
			break;
		}
	}

	const fileName string = "extensions.txt";
	
	if code == 1 {
		SaveExtensions(fileName);
	}

	if code == 2 {
		InstallExtensions(fileName);
	}

	fmt.Println("Thank you, bye!");
	fmt.Scanln();
}

func GetExtensions() string {
	const executor string = "code";
	const command string = "--list-extensions" ;
	
	cmd := exec.Command(executor, command);
	
	out, err := cmd.CombinedOutput();
	if err != nil {
		log.Fatalf("Failed with %s\n", err);
	}

	return string(out);
}

func SaveExtensions(fileName string) {
	var extensions string = GetExtensions();
	extensions = strings.Replace(extensions, "\n", "\r\n", -1)

	path := GetDesktopPath() + fileName;

	fout, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755);
	
	if err != nil {
		panic("Unable to create or open file");
	}

	defer fout.Close();

	fout.WriteString(extensions);
}

func InstallExtensions(fileName string) {

}

func CheckIfCodeInPath() {
    _, err := exec.LookPath("code");
    if err != nil {
        panic("'code' executable isn't in path, sorry");
    } else {
        fmt.Println("'code' executable is in PATH");
    }
}

func GetDesktopPath() string {
	curUser, error := user.Current();
	
	if error != nil {
	  panic(error);
	}

	return curUser.HomeDir + "/Desktop/";
}