package main

// Check - C unit testing framework
// Homepage: https://libcheck.github.io/check/

import (
	"fmt"
	
	"os/exec"
)

func installCheck() {
	// Método 1: Descargar y extraer .tar.gz
	check_tar_url := "https://github.com/libcheck/check/releases/download/0.15.2/check-0.15.2.tar.gz"
	check_cmd_tar := exec.Command("curl", "-L", check_tar_url, "-o", "package.tar.gz")
	err := check_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	check_zip_url := "https://github.com/libcheck/check/releases/download/0.15.2/check-0.15.2.zip"
	check_cmd_zip := exec.Command("curl", "-L", check_zip_url, "-o", "package.zip")
	err = check_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	check_bin_url := "https://github.com/libcheck/check/releases/download/0.15.2/check-0.15.2.bin"
	check_cmd_bin := exec.Command("curl", "-L", check_bin_url, "-o", "binary.bin")
	err = check_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	check_src_url := "https://github.com/libcheck/check/releases/download/0.15.2/check-0.15.2.src.tar.gz"
	check_cmd_src := exec.Command("curl", "-L", check_src_url, "-o", "source.tar.gz")
	err = check_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	check_cmd_direct := exec.Command("./binary")
	err = check_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gawk")
	exec.Command("latte", "install", "gawk").Run()
}
