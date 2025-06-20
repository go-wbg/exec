package exec

type Cmd interface {
	CombinedOutput() ([]byte, error)
	Environ() []string
	Output() ([]byte, error)
	Run() error
	Start() error
	StderrPipe() (io.ReadCloser, error)
	StdinPipe() (io.WriteCloser, error)
	StdoutPipe() (io.ReadCloser, error)
	String() string
	Wait() error
}

func Command(name string, arg ...string) Cmd {

}

func CommandContext(ctx context.Context, name string, arg ...string) Cmd {

}

type Error struct{
	exec.Error
}

type ExitError struct {
	exec.ExitError
}