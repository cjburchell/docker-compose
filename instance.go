package docker_compose

import (
	"io"
	"os"
	"os/exec"
)

type IContainers interface {
	Down() error
	Log() error
	Up() error
	Build() error
	Stop() error
	LogService(service string) error
	LogServiceWithHandler(service string, output io.Writer) error
}

type containers struct {
	path string
}

func (i containers) Build() error {
	cmd, err := i.dockerCommand("build", os.Stdout)
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
	_, err := i.dockerCommand("logs", os.Stdout, "-f", "--tail=0")
	return err
}

func (i containers) LogService(service string) error {
	_, err := i.dockerCommand("logs", os.Stdout, "-f", "--tail=0", service)
	return err
}

func (i containers) LogServiceWithHandler(service string, output io.Writer) error {
	_, err := i.dockerCommand("logs", output, "-f", "--tail=0", service)
	return err
}

func (i containers) Up() error {
	cmd, err := i.dockerCommand("up", os.Stdout, "-d")
	if err != nil {
		return err
	}

	return cmd.Wait()
}

func (i containers) dockerCommand(command string, output io.Writer, args ...string) (*exec.Cmd, error) {
	var cmd *exec.Cmd
	var arguments []string
	if len(i.path) == 0 {
		arguments = append(append(arguments, command), args...)
	} else {
		arguments = append(append(arguments, "--file", i.path, command), args...)
	}

	cmd = exec.Command("docker-compose", arguments...)
	cmd.Stdout = output
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		return nil, err
	}

	return cmd, nil
}

func (i containers) Down() error {
	cmd, err := i.dockerCommand("down", os.Stdout)
	if err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		return err
	}

	return nil
}

func (i containers) Stop() error {
	cmd, err := i.dockerCommand("stop", os.Stdout)
	if err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		return err
	}

	return nil
}
