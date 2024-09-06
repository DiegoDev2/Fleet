package main

// Mytop - Top-like query monitor for MySQL
// Homepage: https://www.mysqlfanboy.com/mytop-3/

import (
	"fmt"
	
	"os/exec"
)

func installMytop() {
	// Método 1: Descargar y extraer .tar.gz
	mytop_tar_url := "https://www.mysqlfanboy.com/mytop-3/mytop-1.9.1.tar.gz"
	mytop_cmd_tar := exec.Command("curl", "-L", mytop_tar_url, "-o", "package.tar.gz")
	err := mytop_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mytop_zip_url := "https://www.mysqlfanboy.com/mytop-3/mytop-1.9.1.zip"
	mytop_cmd_zip := exec.Command("curl", "-L", mytop_zip_url, "-o", "package.zip")
	err = mytop_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mytop_bin_url := "https://www.mysqlfanboy.com/mytop-3/mytop-1.9.1.bin"
	mytop_cmd_bin := exec.Command("curl", "-L", mytop_bin_url, "-o", "binary.bin")
	err = mytop_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mytop_src_url := "https://www.mysqlfanboy.com/mytop-3/mytop-1.9.1.src.tar.gz"
	mytop_cmd_src := exec.Command("curl", "-L", mytop_src_url, "-o", "source.tar.gz")
	err = mytop_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mytop_cmd_direct := exec.Command("./binary")
	err = mytop_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: mysql-client")
	exec.Command("latte", "install", "mysql-client").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
