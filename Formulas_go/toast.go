package main

// Toast - Tool for running tasks in containers
// Homepage: https://github.com/stepchowfun/toast

import (
	"fmt"
	
	"os/exec"
)

func installToast() {
	// Método 1: Descargar y extraer .tar.gz
	toast_tar_url := "https://github.com/stepchowfun/toast/archive/refs/tags/v0.47.6.tar.gz"
	toast_cmd_tar := exec.Command("curl", "-L", toast_tar_url, "-o", "package.tar.gz")
	err := toast_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	toast_zip_url := "https://github.com/stepchowfun/toast/archive/refs/tags/v0.47.6.zip"
	toast_cmd_zip := exec.Command("curl", "-L", toast_zip_url, "-o", "package.zip")
	err = toast_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	toast_bin_url := "https://github.com/stepchowfun/toast/archive/refs/tags/v0.47.6.bin"
	toast_cmd_bin := exec.Command("curl", "-L", toast_bin_url, "-o", "binary.bin")
	err = toast_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	toast_src_url := "https://github.com/stepchowfun/toast/archive/refs/tags/v0.47.6.src.tar.gz"
	toast_cmd_src := exec.Command("curl", "-L", toast_src_url, "-o", "source.tar.gz")
	err = toast_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	toast_cmd_direct := exec.Command("./binary")
	err = toast_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
