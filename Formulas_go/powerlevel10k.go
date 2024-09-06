package main

// Powerlevel10k - Theme for zsh
// Homepage: https://github.com/romkatv/powerlevel10k

import (
	"fmt"
	
	"os/exec"
)

func installPowerlevel10k() {
	// Método 1: Descargar y extraer .tar.gz
	powerlevel10k_tar_url := "https://github.com/romkatv/powerlevel10k/archive/refs/tags/v1.20.0.tar.gz"
	powerlevel10k_cmd_tar := exec.Command("curl", "-L", powerlevel10k_tar_url, "-o", "package.tar.gz")
	err := powerlevel10k_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	powerlevel10k_zip_url := "https://github.com/romkatv/powerlevel10k/archive/refs/tags/v1.20.0.zip"
	powerlevel10k_cmd_zip := exec.Command("curl", "-L", powerlevel10k_zip_url, "-o", "package.zip")
	err = powerlevel10k_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	powerlevel10k_bin_url := "https://github.com/romkatv/powerlevel10k/archive/refs/tags/v1.20.0.bin"
	powerlevel10k_cmd_bin := exec.Command("curl", "-L", powerlevel10k_bin_url, "-o", "binary.bin")
	err = powerlevel10k_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	powerlevel10k_src_url := "https://github.com/romkatv/powerlevel10k/archive/refs/tags/v1.20.0.src.tar.gz"
	powerlevel10k_cmd_src := exec.Command("curl", "-L", powerlevel10k_src_url, "-o", "source.tar.gz")
	err = powerlevel10k_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	powerlevel10k_cmd_direct := exec.Command("./binary")
	err = powerlevel10k_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
