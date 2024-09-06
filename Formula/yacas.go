package main

// Yacas - General purpose computer algebra system
// Homepage: https://www.yacas.org/

import (
	"fmt"
	
	"os/exec"
)

func installYacas() {
	// Método 1: Descargar y extraer .tar.gz
	yacas_tar_url := "https://github.com/grzegorzmazur/yacas/archive/refs/tags/v1.9.1.tar.gz"
	yacas_cmd_tar := exec.Command("curl", "-L", yacas_tar_url, "-o", "package.tar.gz")
	err := yacas_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	yacas_zip_url := "https://github.com/grzegorzmazur/yacas/archive/refs/tags/v1.9.1.zip"
	yacas_cmd_zip := exec.Command("curl", "-L", yacas_zip_url, "-o", "package.zip")
	err = yacas_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	yacas_bin_url := "https://github.com/grzegorzmazur/yacas/archive/refs/tags/v1.9.1.bin"
	yacas_cmd_bin := exec.Command("curl", "-L", yacas_bin_url, "-o", "binary.bin")
	err = yacas_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	yacas_src_url := "https://github.com/grzegorzmazur/yacas/archive/refs/tags/v1.9.1.src.tar.gz"
	yacas_cmd_src := exec.Command("curl", "-L", yacas_src_url, "-o", "source.tar.gz")
	err = yacas_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	yacas_cmd_direct := exec.Command("./binary")
	err = yacas_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
