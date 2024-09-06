package main

// Scw - Command-line Interface for Scaleway
// Homepage: https://github.com/scaleway/scaleway-cli

import (
	"fmt"
	
	"os/exec"
)

func installScw() {
	// Método 1: Descargar y extraer .tar.gz
	scw_tar_url := "https://github.com/scaleway/scaleway-cli/archive/refs/tags/v2.33.0.tar.gz"
	scw_cmd_tar := exec.Command("curl", "-L", scw_tar_url, "-o", "package.tar.gz")
	err := scw_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	scw_zip_url := "https://github.com/scaleway/scaleway-cli/archive/refs/tags/v2.33.0.zip"
	scw_cmd_zip := exec.Command("curl", "-L", scw_zip_url, "-o", "package.zip")
	err = scw_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	scw_bin_url := "https://github.com/scaleway/scaleway-cli/archive/refs/tags/v2.33.0.bin"
	scw_cmd_bin := exec.Command("curl", "-L", scw_bin_url, "-o", "binary.bin")
	err = scw_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	scw_src_url := "https://github.com/scaleway/scaleway-cli/archive/refs/tags/v2.33.0.src.tar.gz"
	scw_cmd_src := exec.Command("curl", "-L", scw_src_url, "-o", "source.tar.gz")
	err = scw_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	scw_cmd_direct := exec.Command("./binary")
	err = scw_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
