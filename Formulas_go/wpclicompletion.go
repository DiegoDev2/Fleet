package main

// WpCliCompletion - Bash completion for Wpcli
// Homepage: https://github.com/wp-cli/wp-cli

import (
	"fmt"
	
	"os/exec"
)

func installWpCliCompletion() {
	// Método 1: Descargar y extraer .tar.gz
	wpclicompletion_tar_url := "https://github.com/wp-cli/wp-cli/archive/refs/tags/v2.11.0.tar.gz"
	wpclicompletion_cmd_tar := exec.Command("curl", "-L", wpclicompletion_tar_url, "-o", "package.tar.gz")
	err := wpclicompletion_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wpclicompletion_zip_url := "https://github.com/wp-cli/wp-cli/archive/refs/tags/v2.11.0.zip"
	wpclicompletion_cmd_zip := exec.Command("curl", "-L", wpclicompletion_zip_url, "-o", "package.zip")
	err = wpclicompletion_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wpclicompletion_bin_url := "https://github.com/wp-cli/wp-cli/archive/refs/tags/v2.11.0.bin"
	wpclicompletion_cmd_bin := exec.Command("curl", "-L", wpclicompletion_bin_url, "-o", "binary.bin")
	err = wpclicompletion_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wpclicompletion_src_url := "https://github.com/wp-cli/wp-cli/archive/refs/tags/v2.11.0.src.tar.gz"
	wpclicompletion_cmd_src := exec.Command("curl", "-L", wpclicompletion_src_url, "-o", "source.tar.gz")
	err = wpclicompletion_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wpclicompletion_cmd_direct := exec.Command("./binary")
	err = wpclicompletion_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
