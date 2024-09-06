package main

// ZshAutosuggestions - Fish-like fast/unobtrusive autosuggestions for zsh
// Homepage: https://github.com/zsh-users/zsh-autosuggestions

import (
	"fmt"
	
	"os/exec"
)

func installZshAutosuggestions() {
	// Método 1: Descargar y extraer .tar.gz
	zshautosuggestions_tar_url := "https://github.com/zsh-users/zsh-autosuggestions/archive/refs/tags/v0.7.0.tar.gz"
	zshautosuggestions_cmd_tar := exec.Command("curl", "-L", zshautosuggestions_tar_url, "-o", "package.tar.gz")
	err := zshautosuggestions_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zshautosuggestions_zip_url := "https://github.com/zsh-users/zsh-autosuggestions/archive/refs/tags/v0.7.0.zip"
	zshautosuggestions_cmd_zip := exec.Command("curl", "-L", zshautosuggestions_zip_url, "-o", "package.zip")
	err = zshautosuggestions_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zshautosuggestions_bin_url := "https://github.com/zsh-users/zsh-autosuggestions/archive/refs/tags/v0.7.0.bin"
	zshautosuggestions_cmd_bin := exec.Command("curl", "-L", zshautosuggestions_bin_url, "-o", "binary.bin")
	err = zshautosuggestions_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zshautosuggestions_src_url := "https://github.com/zsh-users/zsh-autosuggestions/archive/refs/tags/v0.7.0.src.tar.gz"
	zshautosuggestions_cmd_src := exec.Command("curl", "-L", zshautosuggestions_src_url, "-o", "source.tar.gz")
	err = zshautosuggestions_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zshautosuggestions_cmd_direct := exec.Command("./binary")
	err = zshautosuggestions_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
