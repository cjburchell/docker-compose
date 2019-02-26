package docker_compose

type IInstance interface {
	Down() error
}

type Instance struct {
}

func Up(file File) (IInstance, error) {

	return nil, nil
}

func UpFile(path string) (IInstance, error) {

	file, err := LoadFile(path)
	if err != nil {
		return nil, err
	}

	return Up(file)
}

func (i *Instance) Down() error {
	return nil
}
