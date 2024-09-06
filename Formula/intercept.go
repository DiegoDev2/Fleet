package main

// Intercept - Static Application Security Testing (SAST) tool
// Homepage: https://intercept.cc

import (
	"fmt"
	
	"os/exec"
)

func installIntercept() {
	// Método 1: Descargar y extraer .tar.gz
	intercept_tar_url := "https://github.com/xfhg/intercept/archive/refs/tags/v1.6.1.tar.gz"
	intercept_cmd_tar := exec.Command("curl", "-L", intercept_tar_url, "-o", "package.tar.gz")
	err := intercept_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	intercept_zip_url := "https://github.com/xfhg/intercept/archive/refs/tags/v1.6.1.zip"
	intercept_cmd_zip := exec.Command("curl", "-L", intercept_zip_url, "-o", "package.zip")
	err = intercept_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	intercept_bin_url := "https://github.com/xfhg/intercept/archive/refs/tags/v1.6.1.bin"
	intercept_cmd_bin := exec.Command("curl", "-L", intercept_bin_url, "-o", "binary.bin")
	err = intercept_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	intercept_src_url := "https://github.com/xfhg/intercept/archive/refs/tags/v1.6.1.src.tar.gz"
	intercept_cmd_src := exec.Command("curl", "-L", intercept_src_url, "-o", "source.tar.gz")
	err = intercept_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	intercept_cmd_direct := exec.Command("./binary")
	err = intercept_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
