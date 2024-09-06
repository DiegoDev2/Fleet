package main

// PackerCompletion - Bash completion for Packer
// Homepage: https://github.com/mrolli/packer-bash-completion

import (
	"fmt"
	
	"os/exec"
)

func installPackerCompletion() {
	// Método 1: Descargar y extraer .tar.gz
	packercompletion_tar_url := "https://github.com/mrolli/packer-bash-completion/archive/refs/tags/1.4.3.tar.gz"
	packercompletion_cmd_tar := exec.Command("curl", "-L", packercompletion_tar_url, "-o", "package.tar.gz")
	err := packercompletion_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	packercompletion_zip_url := "https://github.com/mrolli/packer-bash-completion/archive/refs/tags/1.4.3.zip"
	packercompletion_cmd_zip := exec.Command("curl", "-L", packercompletion_zip_url, "-o", "package.zip")
	err = packercompletion_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	packercompletion_bin_url := "https://github.com/mrolli/packer-bash-completion/archive/refs/tags/1.4.3.bin"
	packercompletion_cmd_bin := exec.Command("curl", "-L", packercompletion_bin_url, "-o", "binary.bin")
	err = packercompletion_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	packercompletion_src_url := "https://github.com/mrolli/packer-bash-completion/archive/refs/tags/1.4.3.src.tar.gz"
	packercompletion_cmd_src := exec.Command("curl", "-L", packercompletion_src_url, "-o", "source.tar.gz")
	err = packercompletion_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	packercompletion_cmd_direct := exec.Command("./binary")
	err = packercompletion_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
