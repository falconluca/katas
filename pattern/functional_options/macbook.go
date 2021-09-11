package functional_options

type Cpu struct {
	Core  int
	CType string
}

type MacBook struct {
	Name string
	Size float32
	Cpu  *Cpu
}

func NewDefaultMacBook(name string, size float32, cpu *Cpu) (*MacBook, error) {
	if cpu == nil {
		cpu = &Cpu{Core: 16, CType: "M2"}
	}
	return &MacBook{Name: name, Size: size, Cpu: cpu}, nil
}
