package sshtemplate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"syscall"
)

// SSHTemplate maps a name to a command
type SSHTemplate struct {
	Name    string
	Command string
}

// WriteToFile writes a map of sshTemplate structs to json file
func WriteToFile(data map[string]SSHTemplate, filePath string) {
	file, _ := json.MarshalIndent(data, "", " ")
	err := ioutil.WriteFile(filePath, file, 0644)
	check(err)
}

// ReadFromFile reads a map of sshTemplate structs from a json file
func ReadFromFile(filePath string) map[string]SSHTemplate {
	dat, err := ioutil.ReadFile(filePath)
	check(err)

	res := map[string]SSHTemplate{}
	json.Unmarshal(dat, &res)

	return res
}

// AddTemplate creates a new template and writes it to a file
func AddTemplate(name string, command string, filePath string) *SSHTemplate {
	templateArray := ReadFromFile(filePath)

	newTemplate := SSHTemplate{
		Name:    name,
		Command: command,
	}

	templateArray[name] = newTemplate

	WriteToFile(templateArray, filePath)

	fmt.Println(fmt.Sprintf("New template '%s' added!", newTemplate.Name))
	return &newTemplate
}

// RemoveTemplate searches for a template and removes it from the file
func RemoveTemplate(name string, filePath string) *SSHTemplate {
	templateArray := ReadFromFile(filePath)

	template := templateArray[name]

	delete(templateArray, name)

	WriteToFile(templateArray, filePath)

	fmt.Println(fmt.Sprintf("Template '%s' removed!", name))
	return &template
}

// ExecuteCommand recalls a saved command and attempts to execute it
func ExecuteCommand(template SSHTemplate) {
	fmt.Println("Executing SSH command for", template.Command)

	binary, lookErr := exec.LookPath("ssh")
	check(lookErr)

	args := []string{"ssh", template.Command}
	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	check(execErr)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
