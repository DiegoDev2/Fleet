package main

// Ugit - Undo git commands. Your damage control git buddy
// Homepage: https://bhupesh.me/undo-your-last-git-mistake-with-ugit/

import (
	"fmt"
	
	"os/exec"
)

func installUgit() {
	// Método 1: Descargar y extraer .tar.gz
	ugit_tar_url := "https://github.com/Bhupesh-V/ugit/archive/refs/tags/v5.8.tar.gz"
	ugit_cmd_tar := exec.Command("curl", "-L", ugit_tar_url, "-o", "package.tar.gz")
	err := ugit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ugit_zip_url := "https://github.com/Bhupesh-V/ugit/archive/refs/tags/v5.8.zip"
	ugit_cmd_zip := exec.Command("curl", "-L", ugit_zip_url, "-o", "package.zip")
	err = ugit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ugit_bin_url := "https://github.com/Bhupesh-V/ugit/archive/refs/tags/v5.8.bin"
	ugit_cmd_bin := exec.Command("curl", "-L", ugit_bin_url, "-o", "binary.bin")
	err = ugit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ugit_src_url := "https://github.com/Bhupesh-V/ugit/archive/refs/tags/v5.8.src.tar.gz"
	ugit_cmd_src := exec.Command("curl", "-L", ugit_src_url, "-o", "source.tar.gz")
	err = ugit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ugit_cmd_direct := exec.Command("./binary")
	err = ugit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bash")
	exec.Command("latte", "install", "bash").Run()
	fmt.Println("Instalando dependencia: fzf")
	exec.Command("latte", "install", "fzf").Run()
}
