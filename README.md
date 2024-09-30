# GO Config

Easily apply environment variables to a struct

## Installation

```shell
go get github.com/robotjoosen/go-config
```

## Usage

```golang
type configuration struct {
	Name string `mapstructure:"NAME"`
}

func example() {
	var cnf configuration
	if _, err := config.Load(&cnf, map[string]any{
		"NAME": "unknown",
	}); err != nil {
		return
	}

	fmt.Println(cnf.Name)
}
```