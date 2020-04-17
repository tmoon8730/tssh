package sshtemplate

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func Test_AddTemplate(t *testing.T) {
	filePath := "test.json"
	os.Create(filePath)

	type args struct {
		name    string
		command string
	}
	tests := []struct {
		name string
		args args
		want *SSHTemplate
	}{
		{"testName", args{name: "testName", command: "testCommand"}, &SSHTemplate{Name: "testName", Command: "testCommand"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddTemplate(tt.args.name, tt.args.command, filePath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTemplate() = %v, want %v", got, tt.want)
			}
		})
	}

	os.Remove(filePath)
}

func Test_RemoveTemplate(t *testing.T) {
	data := map[string]SSHTemplate{"testName": {Name: "testName", Command: "testCommand"}}
	file, _ := json.MarshalIndent(data, "", " ")
	filePath := "test.json"
	os.Create(filePath)

	ioutil.WriteFile(filePath, file, 0644)

	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *SSHTemplate
	}{
		{"testName", args{name: "testName"}, &SSHTemplate{Name: "testName", Command: "testCommand"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveTemplate(tt.args.name, filePath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeTemplate() = %v, want %v", got, tt.want)
			}
		})
	}

	os.Remove(filePath)
}

func Test_ReadFromFile(t *testing.T) {
	data := map[string]SSHTemplate{"testName": {Name: "testName", Command: "testCommand"}}
	file, _ := json.MarshalIndent(data, "", " ")
	filePath := "test.json"
	os.Create(filePath)

	ioutil.WriteFile(filePath, file, 0644)

	t.Run("test reading", func(t *testing.T) {
		if got := ReadFromFile(filePath); !reflect.DeepEqual(got, data) {
			t.Errorf("readFromFile() = %v, want %v", got, data)
		}
	})

	os.Remove(filePath)
}

func Test_WriteToFile(t *testing.T) {
	data := map[string]SSHTemplate{"testName": {Name: "testName", Command: "testCommand"}}
	filePath := "test.json"
	os.Create(filePath)

	t.Run("test writing", func(t *testing.T) {
		WriteToFile(data, filePath)

		fileData, _ := ioutil.ReadFile(filePath)
		res := map[string]SSHTemplate{}
		json.Unmarshal(fileData, &res)

		if !reflect.DeepEqual(res, data) {
			t.Errorf("writeToFile() = %v, want %v", res, data)
		}
	})

	os.Remove(filePath)
}
