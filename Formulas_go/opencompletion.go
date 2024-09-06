package main

// OpenCompletion - Bash completion for open
// Homepage: https://github.com/moshen/open-bash-completion

import (
	"fmt"
	
	"os/exec"
)

func installOpenCompletion() {
	// Método 1: Descargar y extraer .tar.gz
	opencompletion_tar_url := "https://github.com/moshen/open-bash-completion/archive/refs/tags/v1.0.5.tar.gz"
	opencompletion_cmd_tar := exec.Command("curl", "-L", opencompletion_tar_url, "-o", "package.tar.gz")
	err := opencompletion_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	opencompletion_zip_url := "https://github.com/moshen/open-bash-completion/archive/refs/tags/v1.0.5.zip"
	opencompletion_cmd_zip := exec.Command("curl", "-L", opencompletion_zip_url, "-o", "package.zip")
	err = opencompletion_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	opencompletion_bin_url := "https://github.com/moshen/open-bash-completion/archive/refs/tags/v1.0.5.bin"
	opencompletion_cmd_bin := exec.Command("curl", "-L", opencompletion_bin_url, "-o", "binary.bin")
	err = opencompletion_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	opencompletion_src_url := "https://github.com/moshen/open-bash-completion/archive/refs/tags/v1.0.5.src.tar.gz"
	opencompletion_cmd_src := exec.Command("curl", "-L", opencompletion_src_url, "-o", "source.tar.gz")
	err = opencompletion_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	opencompletion_cmd_direct := exec.Command("./binary")
	err = opencompletion_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
