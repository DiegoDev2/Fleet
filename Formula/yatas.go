package main

// Yatas - Tool to audit AWS/GCP infrastructure for misconfiguration or security issues
// Homepage: https://github.com/padok-team/yatas

import (
	"fmt"
	
	"os/exec"
)

func installYatas() {
	// Método 1: Descargar y extraer .tar.gz
	yatas_tar_url := "https://github.com/padok-team/yatas/archive/refs/tags/v1.5.1.tar.gz"
	yatas_cmd_tar := exec.Command("curl", "-L", yatas_tar_url, "-o", "package.tar.gz")
	err := yatas_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	yatas_zip_url := "https://github.com/padok-team/yatas/archive/refs/tags/v1.5.1.zip"
	yatas_cmd_zip := exec.Command("curl", "-L", yatas_zip_url, "-o", "package.zip")
	err = yatas_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	yatas_bin_url := "https://github.com/padok-team/yatas/archive/refs/tags/v1.5.1.bin"
	yatas_cmd_bin := exec.Command("curl", "-L", yatas_bin_url, "-o", "binary.bin")
	err = yatas_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	yatas_src_url := "https://github.com/padok-team/yatas/archive/refs/tags/v1.5.1.src.tar.gz"
	yatas_cmd_src := exec.Command("curl", "-L", yatas_src_url, "-o", "source.tar.gz")
	err = yatas_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	yatas_cmd_direct := exec.Command("./binary")
	err = yatas_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
