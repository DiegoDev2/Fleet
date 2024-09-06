package main

// Brev - CLI tool for managing workspaces provided by brev.dev
// Homepage: https://docs.brev.dev

import (
	"fmt"
	
	"os/exec"
)

func installBrev() {
	// Método 1: Descargar y extraer .tar.gz
	brev_tar_url := "https://github.com/brevdev/brev-cli/archive/refs/tags/v0.6.287.tar.gz"
	brev_cmd_tar := exec.Command("curl", "-L", brev_tar_url, "-o", "package.tar.gz")
	err := brev_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	brev_zip_url := "https://github.com/brevdev/brev-cli/archive/refs/tags/v0.6.287.zip"
	brev_cmd_zip := exec.Command("curl", "-L", brev_zip_url, "-o", "package.zip")
	err = brev_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	brev_bin_url := "https://github.com/brevdev/brev-cli/archive/refs/tags/v0.6.287.bin"
	brev_cmd_bin := exec.Command("curl", "-L", brev_bin_url, "-o", "binary.bin")
	err = brev_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	brev_src_url := "https://github.com/brevdev/brev-cli/archive/refs/tags/v0.6.287.src.tar.gz"
	brev_cmd_src := exec.Command("curl", "-L", brev_src_url, "-o", "source.tar.gz")
	err = brev_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	brev_cmd_direct := exec.Command("./binary")
	err = brev_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
