package main

// Dcp - Docker cp made easy
// Homepage: https://github.com/exdx/dcp

import (
	"fmt"
	
	"os/exec"
)

func installDcp() {
	// Método 1: Descargar y extraer .tar.gz
	dcp_tar_url := "https://github.com/exdx/dcp/archive/refs/tags/v0.4.1.tar.gz"
	dcp_cmd_tar := exec.Command("curl", "-L", dcp_tar_url, "-o", "package.tar.gz")
	err := dcp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dcp_zip_url := "https://github.com/exdx/dcp/archive/refs/tags/v0.4.1.zip"
	dcp_cmd_zip := exec.Command("curl", "-L", dcp_zip_url, "-o", "package.zip")
	err = dcp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dcp_bin_url := "https://github.com/exdx/dcp/archive/refs/tags/v0.4.1.bin"
	dcp_cmd_bin := exec.Command("curl", "-L", dcp_bin_url, "-o", "binary.bin")
	err = dcp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dcp_src_url := "https://github.com/exdx/dcp/archive/refs/tags/v0.4.1.src.tar.gz"
	dcp_cmd_src := exec.Command("curl", "-L", dcp_src_url, "-o", "source.tar.gz")
	err = dcp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dcp_cmd_direct := exec.Command("./binary")
	err = dcp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
