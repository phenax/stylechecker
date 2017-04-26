package linter

func NewJSLinter() (*Linter) {

	return &Linter{

		name: "echo",

		args: []string{
			"foobar",
		},
	}
}
