package main

// Authz0 - Automated authorization test tool
// Homepage: https://authz0.hahwul.com/

import (
	"fmt"
	
	"os/exec"
)

func installAuthz0() {
	// Método 1: Descargar y extraer .tar.gz
	authz0_tar_url := "https://github.com/hahwul/authz0/archive/refs/tags/v1.1.2.tar.gz"
	authz0_cmd_tar := exec.Command("curl", "-L", authz0_tar_url, "-o", "package.tar.gz")
	err := authz0_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	authz0_zip_url := "https://github.com/hahwul/authz0/archive/refs/tags/v1.1.2.zip"
	authz0_cmd_zip := exec.Command("curl", "-L", authz0_zip_url, "-o", "package.zip")
	err = authz0_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	authz0_bin_url := "https://github.com/hahwul/authz0/archive/refs/tags/v1.1.2.bin"
	authz0_cmd_bin := exec.Command("curl", "-L", authz0_bin_url, "-o", "binary.bin")
	err = authz0_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	authz0_src_url := "https://github.com/hahwul/authz0/archive/refs/tags/v1.1.2.src.tar.gz"
	authz0_cmd_src := exec.Command("curl", "-L", authz0_src_url, "-o", "source.tar.gz")
	err = authz0_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	authz0_cmd_direct := exec.Command("./binary")
	err = authz0_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
