package docker_compose

import (
	"os"
	"os/exec"

	"github.com/cjburchell/go-uatu"
)

type IStack interface {
	Down()
	Log() error
	Up() error
	Build() error
	Stop()
	LogService(service string) error
}

type stack struct {
	path string
}

func (i stack) Build() error {
	cmd, err := i.dockerCommand("build")
	if err != nil {
		return err
	}

	return cmd.Wait()
}

func Create() IStack {
	return stack{}
}

func CreateFile(path string) IStack {
	return stack{path: path}
}

func (i stack) Log() error {
	_, err := i.dockerCommand("logs", "-f", "--tail=0")
	return err
}

func (i stack) LogService(service string) error {
	_, err := i.dockerCommand("logs", "-f", "--tail=0", service)
	return err
}

func (i stack) Up() error {
	cmd, err := i.dockerCommand("up", "-d")
	if err != nil {
		return err
	}

	return cmd.Wait()
}

func (i stack) dockerCommand(command string, args ...string) (*exec.Cmd, error) {
	var cmd *exec.Cmd
	var arguments []string
	if len(i.path) == 0 {
		arguments = append(append(arguments, command), args...)
	} else {
		arguments = append(append(arguments, "--file", i.path, command), args...)
	}

	cmd = exec.Command("docker-compose", arguments...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		return nil, err
	}

	return cmd, nil
}

func (i stack) Down() {
	cmd, err := i.dockerCommand("down")
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}

func (i stack) Stop() {
	cmd, err := i.dockerCommand("stop")
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}
