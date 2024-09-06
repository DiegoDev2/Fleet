package main

// Goto - Bash tool for navigation to aliased directories with auto-completion
// Homepage: https://github.com/iridakos/goto

import (
	"fmt"
	
	"os/exec"
)

func installGoto() {
	// Método 1: Descargar y extraer .tar.gz
	goto_tar_url := "https://github.com/iridakos/goto/archive/refs/tags/v2.0.0.tar.gz"
	goto_cmd_tar := exec.Command("curl", "-L", goto_tar_url, "-o", "package.tar.gz")
	err := goto_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	goto_zip_url := "https://github.com/iridakos/goto/archive/refs/tags/v2.0.0.zip"
	goto_cmd_zip := exec.Command("curl", "-L", goto_zip_url, "-o", "package.zip")
	err = goto_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	goto_bin_url := "https://github.com/iridakos/goto/archive/refs/tags/v2.0.0.bin"
	goto_cmd_bin := exec.Command("curl", "-L", goto_bin_url, "-o", "binary.bin")
	err = goto_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	goto_src_url := "https://github.com/iridakos/goto/archive/refs/tags/v2.0.0.src.tar.gz"
	goto_cmd_src := exec.Command("curl", "-L", goto_src_url, "-o", "source.tar.gz")
	err = goto_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	goto_cmd_direct := exec.Command("./binary")
	err = goto_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
