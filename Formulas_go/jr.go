package main

// Jr - CLI program that helps you to create quality random data for your applications
// Homepage: https://jrnd.io/

import (
	"fmt"
	
	"os/exec"
)

func installJr() {
	// Método 1: Descargar y extraer .tar.gz
	jr_tar_url := "https://github.com/jrnd-io/jr/archive/refs/tags/v0.3.9.tar.gz"
	jr_cmd_tar := exec.Command("curl", "-L", jr_tar_url, "-o", "package.tar.gz")
	err := jr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jr_zip_url := "https://github.com/jrnd-io/jr/archive/refs/tags/v0.3.9.zip"
	jr_cmd_zip := exec.Command("curl", "-L", jr_zip_url, "-o", "package.zip")
	err = jr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jr_bin_url := "https://github.com/jrnd-io/jr/archive/refs/tags/v0.3.9.bin"
	jr_cmd_bin := exec.Command("curl", "-L", jr_bin_url, "-o", "binary.bin")
	err = jr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jr_src_url := "https://github.com/jrnd-io/jr/archive/refs/tags/v0.3.9.src.tar.gz"
	jr_cmd_src := exec.Command("curl", "-L", jr_src_url, "-o", "source.tar.gz")
	err = jr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jr_cmd_direct := exec.Command("./binary")
	err = jr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
