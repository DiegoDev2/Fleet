package main

// Atlantis - Terraform Pull Request Automation tool
// Homepage: https://www.runatlantis.io/

import (
	"fmt"
	
	"os/exec"
)

func installAtlantis() {
	// Método 1: Descargar y extraer .tar.gz
	atlantis_tar_url := "https://github.com/runatlantis/atlantis/archive/refs/tags/v0.29.0.tar.gz"
	atlantis_cmd_tar := exec.Command("curl", "-L", atlantis_tar_url, "-o", "package.tar.gz")
	err := atlantis_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	atlantis_zip_url := "https://github.com/runatlantis/atlantis/archive/refs/tags/v0.29.0.zip"
	atlantis_cmd_zip := exec.Command("curl", "-L", atlantis_zip_url, "-o", "package.zip")
	err = atlantis_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	atlantis_bin_url := "https://github.com/runatlantis/atlantis/archive/refs/tags/v0.29.0.bin"
	atlantis_cmd_bin := exec.Command("curl", "-L", atlantis_bin_url, "-o", "binary.bin")
	err = atlantis_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	atlantis_src_url := "https://github.com/runatlantis/atlantis/archive/refs/tags/v0.29.0.src.tar.gz"
	atlantis_cmd_src := exec.Command("curl", "-L", atlantis_src_url, "-o", "source.tar.gz")
	err = atlantis_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	atlantis_cmd_direct := exec.Command("./binary")
	err = atlantis_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
