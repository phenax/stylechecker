package linter

func NewPHPLinter() (*Linter) {

	return &Linter{

		Name: "echo",

		Args: []string{
			"phplinter",
		},
	}
}
