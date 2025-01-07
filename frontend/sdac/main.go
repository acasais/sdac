package main

import (
	"log"
	"os"

	"github.com/acasais/sdac/frontend/mvc/controller"
)

type sdac struct {
}

func main() {
	newSdac().run()
}

func newSdac() *sdac {
	return &sdac{}
}

func (s *sdac) run() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("[recover]:", err)
			os.Exit(1)
		}
	}()

	mc := controller.NewMainController()
	mc.BuildAndRunUI()
}
