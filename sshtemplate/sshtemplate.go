package sshtemplate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"sshtemplate/utilities"
	"syscall"
)

// SSHTemplate maps a name to a command
type SSHTemplate struct {
	Name      string
	Arguments []string
}

// WriteToFile writes a map of sshTemplate structs to json file
func WriteToFile(data map[string]SSHTemplate, filePath string) {
	file, _ := json.MarshalIndent(data, "", " ")
	err := ioutil.WriteFile(filePath, file, 0644)
	utilities.Check(err)
}

// ReadFromFile reads a map of sshTemplate structs from a json file
func ReadFromFile(filePath string) map[string]SSHTemplate {
	dat, err := ioutil.ReadFile(filePath)
	utilities.Check(err)

	res := map[string]SSHTemplate{}
	json.Unmarshal(dat, &res)

	return res
}

// AddTemplate creates a new template and writes it to a file
func AddTemplate(name string, filePath string, arguments []string) *SSHTemplate {
	templateArray := ReadFromFile(filePath)

	newTemplate := SSHTemplate{
		Name:      name,
		Arguments: arguments,
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

// ListTemplates lists existing templates
func ListTemplates(filePath string) *map[string]SSHTemplate {
	templateMap := ReadFromFile(filePath)

	fmt.Println("=== Saved SSH Templates ===")
	for key, element := range templateMap {
		fmt.Print(key, ":  ssh ")
		for i := 0; i < len(element.Arguments); i++ {
			fmt.Print(element.Arguments[i], " ")
		}
		fmt.Println()
	}
	return &templateMap
}

// ExecuteCommand recalls a saved command and attempts to execute it
func ExecuteCommand(template SSHTemplate) {
	fmt.Println("Executing SSH command for", template.Arguments)

	binary, lookErr := exec.LookPath("ssh")
	utilities.Check(lookErr)

	args := append([]string{"ssh"}, template.Arguments...)
	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	utilities.Check(execErr)
}
