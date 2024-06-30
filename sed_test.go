package sed_test

import (
	"testing"

	sed "github.com/Kangnning/go-simple-sed"
)

func TestRun(t *testing.T) {
	s := sed.New()
	conf := sed.Config{
		FileName:  "./QueryRoute.go",
		Opt:       sed.InsertBefore,
		Pattern:   "test.*",
		DesString: "This is inerst before test\n next line",
	}
	s.Run(conf)
	conf.Opt = sed.Delete
	s.Run(conf)
	conf.Opt = sed.Replace
	s.Run(conf)
}
