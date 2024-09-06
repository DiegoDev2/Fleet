package main

// Cmctl - Command-line tool to manage cert-manager
// Homepage: https://cert-manager.io

import (
	"fmt"
	
	"os/exec"
)

func installCmctl() {
	// Método 1: Descargar y extraer .tar.gz
	cmctl_tar_url := "https://github.com/cert-manager/cmctl/archive/refs/tags/v2.1.0.tar.gz"
	cmctl_cmd_tar := exec.Command("curl", "-L", cmctl_tar_url, "-o", "package.tar.gz")
	err := cmctl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cmctl_zip_url := "https://github.com/cert-manager/cmctl/archive/refs/tags/v2.1.0.zip"
	cmctl_cmd_zip := exec.Command("curl", "-L", cmctl_zip_url, "-o", "package.zip")
	err = cmctl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cmctl_bin_url := "https://github.com/cert-manager/cmctl/archive/refs/tags/v2.1.0.bin"
	cmctl_cmd_bin := exec.Command("curl", "-L", cmctl_bin_url, "-o", "binary.bin")
	err = cmctl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cmctl_src_url := "https://github.com/cert-manager/cmctl/archive/refs/tags/v2.1.0.src.tar.gz"
	cmctl_cmd_src := exec.Command("curl", "-L", cmctl_src_url, "-o", "source.tar.gz")
	err = cmctl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cmctl_cmd_direct := exec.Command("./binary")
	err = cmctl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
