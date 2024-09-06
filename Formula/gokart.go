package main

// Gokart - Static code analysis for securing Go code
// Homepage: https://github.com/praetorian-inc/gokart

import (
	"fmt"
	
	"os/exec"
)

func installGokart() {
	// Método 1: Descargar y extraer .tar.gz
	gokart_tar_url := "https://github.com/praetorian-inc/gokart/archive/refs/tags/v0.5.1.tar.gz"
	gokart_cmd_tar := exec.Command("curl", "-L", gokart_tar_url, "-o", "package.tar.gz")
	err := gokart_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gokart_zip_url := "https://github.com/praetorian-inc/gokart/archive/refs/tags/v0.5.1.zip"
	gokart_cmd_zip := exec.Command("curl", "-L", gokart_zip_url, "-o", "package.zip")
	err = gokart_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gokart_bin_url := "https://github.com/praetorian-inc/gokart/archive/refs/tags/v0.5.1.bin"
	gokart_cmd_bin := exec.Command("curl", "-L", gokart_bin_url, "-o", "binary.bin")
	err = gokart_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gokart_src_url := "https://github.com/praetorian-inc/gokart/archive/refs/tags/v0.5.1.src.tar.gz"
	gokart_cmd_src := exec.Command("curl", "-L", gokart_src_url, "-o", "source.tar.gz")
	err = gokart_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gokart_cmd_direct := exec.Command("./binary")
	err = gokart_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
