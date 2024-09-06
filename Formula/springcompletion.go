package main

// SpringCompletion - Bash completion for Spring
// Homepage: https://github.com/jacaetevha/spring_bash_completion

import (
	"fmt"
	
	"os/exec"
)

func installSpringCompletion() {
	// Método 1: Descargar y extraer .tar.gz
	springcompletion_tar_url := "https://github.com/jacaetevha/spring_bash_completion/archive/refs/tags/v0.0.1.tar.gz"
	springcompletion_cmd_tar := exec.Command("curl", "-L", springcompletion_tar_url, "-o", "package.tar.gz")
	err := springcompletion_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	springcompletion_zip_url := "https://github.com/jacaetevha/spring_bash_completion/archive/refs/tags/v0.0.1.zip"
	springcompletion_cmd_zip := exec.Command("curl", "-L", springcompletion_zip_url, "-o", "package.zip")
	err = springcompletion_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	springcompletion_bin_url := "https://github.com/jacaetevha/spring_bash_completion/archive/refs/tags/v0.0.1.bin"
	springcompletion_cmd_bin := exec.Command("curl", "-L", springcompletion_bin_url, "-o", "binary.bin")
	err = springcompletion_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	springcompletion_src_url := "https://github.com/jacaetevha/spring_bash_completion/archive/refs/tags/v0.0.1.src.tar.gz"
	springcompletion_cmd_src := exec.Command("curl", "-L", springcompletion_src_url, "-o", "source.tar.gz")
	err = springcompletion_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	springcompletion_cmd_direct := exec.Command("./binary")
	err = springcompletion_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
