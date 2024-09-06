package main

// Yh - YAML syntax highlighter to bring colours where only jq could
// Homepage: https://github.com/andreazorzetto/yh

import (
	"fmt"
	
	"os/exec"
)

func installYh() {
	// Método 1: Descargar y extraer .tar.gz
	yh_tar_url := "https://github.com/andreazorzetto/yh/archive/refs/tags/v0.4.0.tar.gz"
	yh_cmd_tar := exec.Command("curl", "-L", yh_tar_url, "-o", "package.tar.gz")
	err := yh_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	yh_zip_url := "https://github.com/andreazorzetto/yh/archive/refs/tags/v0.4.0.zip"
	yh_cmd_zip := exec.Command("curl", "-L", yh_zip_url, "-o", "package.zip")
	err = yh_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	yh_bin_url := "https://github.com/andreazorzetto/yh/archive/refs/tags/v0.4.0.bin"
	yh_cmd_bin := exec.Command("curl", "-L", yh_bin_url, "-o", "binary.bin")
	err = yh_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	yh_src_url := "https://github.com/andreazorzetto/yh/archive/refs/tags/v0.4.0.src.tar.gz"
	yh_cmd_src := exec.Command("curl", "-L", yh_src_url, "-o", "source.tar.gz")
	err = yh_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	yh_cmd_direct := exec.Command("./binary")
	err = yh_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
