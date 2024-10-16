package cli

import (
	"fmt"
	"os/exec"
)

func Install(pkg string) {
	fmt.Println("Installing " + pkg)
	exec.LookPath("Formulas/" + pkg + ".go")
	exec.Command("go", "run", "Formulas/"+pkg+".go").Run()
	fmt.Println("Done")
}
