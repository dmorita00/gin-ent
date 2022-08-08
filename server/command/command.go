package command

import "flag"

func Run(name string) error {
	flag.Parse()

	switch name {
	case "gensql":
		return RunGensql()
	}
	return nil
}
