package main

// YarnCompletion - Bash completion for Yarn
// Homepage: https://github.com/dsifford/yarn-completion

import (
	"fmt"
	
	"os/exec"
)

func installYarnCompletion() {
	// Método 1: Descargar y extraer .tar.gz
	yarncompletion_tar_url := "https://github.com/dsifford/yarn-completion/archive/refs/tags/v0.17.0.tar.gz"
	yarncompletion_cmd_tar := exec.Command("curl", "-L", yarncompletion_tar_url, "-o", "package.tar.gz")
	err := yarncompletion_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	yarncompletion_zip_url := "https://github.com/dsifford/yarn-completion/archive/refs/tags/v0.17.0.zip"
	yarncompletion_cmd_zip := exec.Command("curl", "-L", yarncompletion_zip_url, "-o", "package.zip")
	err = yarncompletion_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	yarncompletion_bin_url := "https://github.com/dsifford/yarn-completion/archive/refs/tags/v0.17.0.bin"
	yarncompletion_cmd_bin := exec.Command("curl", "-L", yarncompletion_bin_url, "-o", "binary.bin")
	err = yarncompletion_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	yarncompletion_src_url := "https://github.com/dsifford/yarn-completion/archive/refs/tags/v0.17.0.src.tar.gz"
	yarncompletion_cmd_src := exec.Command("curl", "-L", yarncompletion_src_url, "-o", "source.tar.gz")
	err = yarncompletion_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	yarncompletion_cmd_direct := exec.Command("./binary")
	err = yarncompletion_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bash")
	exec.Command("latte", "install", "bash").Run()
}
