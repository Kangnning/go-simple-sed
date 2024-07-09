package sed_test

import (
	"testing"

	sed "github.com/Kangnning/go-simple-sed"
	"github.com/Kangnning/go-simple-sed/config"
)

func TestRun(t *testing.T) {
	s := sed.New()
	conf := config.New(config.WithFileName("./QueryRoute"), config.WithAction(config.InsertBefore),
		config.WithPattern("test.*"), config.WithDesString("This is inerst before test\n next line"))
	// conf := config.Config{
	// 	FileName:  "./QueryRoute.go",
	// 	Act:       config.InsertBefore,
	// 	Pattern:   "test.*",
	// 	DesString: "This is inerst before test\n next line",
	// }

	s.Run(conf)
	conf.Modify(config.WithAction(config.Delete))
	// conf.Act = config.Delete
	s.Run(conf)
	conf.Modify(config.WithAction(config.Replace))
	s.Run(conf)
}
