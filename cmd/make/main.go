package main

import (
	"fmt"
	"os/exec"
)

func main() {
	if err := exec.Command("gopherjs", "build", "./cmd/app/main.go", "-o", "./assets/index.js").Run(); err != nil {
		panic(err)
	}
	fmt.Println("Build complete")
}
