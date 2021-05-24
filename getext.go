package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"log"
	"strings"
	"bufio"
)

func main() {

	const executor string = "code";
	const command string = "--list-extensions" ;
	
	_ = exec.Command(executor, command);

	

	/*CheckIfCodeInPath();
	
	var code int = 0;
	
	for {
		fmt.Println("Install from desktop/extensions.txt - 1\nSave to desktop/extensions.txt - 2");
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
		InstallExtensions(fileName);
	}

	if code == 2 {
		SaveExtensions(fileName);
	}

	fmt.Println("Completed, bye!");
*/

	//InstallExtensions("extensions.txt");
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
	path := GetDesktopPath() + fileName;
	const executor string = "code";

	fin, open_err := os.Open(path);

	if open_err != nil {
		panic("Unable to open file");
	}

	defer fin.Close();

	scanner := bufio.NewScanner(fin);
	for scanner.Scan() {
		command := "--install-extension " + scanner.Text();
		fmt.Println(command);
		_ = exec.Command(executor, command);
	}

	err := scanner.Err();

	if err != nil {
		panic("Something went wrong...");
	}
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
