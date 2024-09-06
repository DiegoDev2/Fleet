package main

// MultiGitStatus - Show uncommitted, untracked and unpushed changes for multiple Git repos
// Homepage: https://github.com/fboender/multi-git-status

import (
	"fmt"
	
	"os/exec"
)

func installMultiGitStatus() {
	// Método 1: Descargar y extraer .tar.gz
	multigitstatus_tar_url := "https://github.com/fboender/multi-git-status/archive/refs/tags/2.2.tar.gz"
	multigitstatus_cmd_tar := exec.Command("curl", "-L", multigitstatus_tar_url, "-o", "package.tar.gz")
	err := multigitstatus_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	multigitstatus_zip_url := "https://github.com/fboender/multi-git-status/archive/refs/tags/2.2.zip"
	multigitstatus_cmd_zip := exec.Command("curl", "-L", multigitstatus_zip_url, "-o", "package.zip")
	err = multigitstatus_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	multigitstatus_bin_url := "https://github.com/fboender/multi-git-status/archive/refs/tags/2.2.bin"
	multigitstatus_cmd_bin := exec.Command("curl", "-L", multigitstatus_bin_url, "-o", "binary.bin")
	err = multigitstatus_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	multigitstatus_src_url := "https://github.com/fboender/multi-git-status/archive/refs/tags/2.2.src.tar.gz"
	multigitstatus_cmd_src := exec.Command("curl", "-L", multigitstatus_src_url, "-o", "source.tar.gz")
	err = multigitstatus_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	multigitstatus_cmd_direct := exec.Command("./binary")
	err = multigitstatus_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
