package main

// Dbdeployer - Tool to deploy sandboxed MySQL database servers
// Homepage: https://github.com/datacharmer/dbdeployer

import (
	"fmt"
	
	"os/exec"
)

func installDbdeployer() {
	// Método 1: Descargar y extraer .tar.gz
	dbdeployer_tar_url := "https://github.com/datacharmer/dbdeployer/archive/refs/tags/v1.73.0.tar.gz"
	dbdeployer_cmd_tar := exec.Command("curl", "-L", dbdeployer_tar_url, "-o", "package.tar.gz")
	err := dbdeployer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dbdeployer_zip_url := "https://github.com/datacharmer/dbdeployer/archive/refs/tags/v1.73.0.zip"
	dbdeployer_cmd_zip := exec.Command("curl", "-L", dbdeployer_zip_url, "-o", "package.zip")
	err = dbdeployer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dbdeployer_bin_url := "https://github.com/datacharmer/dbdeployer/archive/refs/tags/v1.73.0.bin"
	dbdeployer_cmd_bin := exec.Command("curl", "-L", dbdeployer_bin_url, "-o", "binary.bin")
	err = dbdeployer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dbdeployer_src_url := "https://github.com/datacharmer/dbdeployer/archive/refs/tags/v1.73.0.src.tar.gz"
	dbdeployer_cmd_src := exec.Command("curl", "-L", dbdeployer_src_url, "-o", "source.tar.gz")
	err = dbdeployer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dbdeployer_cmd_direct := exec.Command("./binary")
	err = dbdeployer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
