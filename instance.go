package docker_compose

import (
	"os"
	"os/exec"

	"github.com/cjburchell/go-uatu"
)

type IContainers interface {
	Down()
	Log() error
	Up() error
	Build() error
	Stop()
	LogService(service string) error
}

type containers struct {
	path string
}

func (i containers) Build() error {
	cmd, err := i.dockerCommand("build")
	if err != nil {
		return err
	}

	return cmd.Wait()
}

func Create() IContainers {
	return containers{}
}

func CreateFile(path string) IContainers {
	return containers{path: path}
}

func (i containers) Log() error {
	_, err := i.dockerCommand("logs", "-f", "--tail=0")
	return err
}

func (i containers) LogService(service string) error {
	_, err := i.dockerCommand("logs", "-f", "--tail=0", service)
	return err
}

func (i containers) Up() error {
	cmd, err := i.dockerCommand("up", "-d")
	if err != nil {
		return err
	}

	return cmd.Wait()
}

func (i containers) dockerCommand(command string, args ...string) (*exec.Cmd, error) {
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

func (i containers) Down() {
	cmd, err := i.dockerCommand("down")
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}

func (i containers) Stop() {
	cmd, err := i.dockerCommand("stop")
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}
