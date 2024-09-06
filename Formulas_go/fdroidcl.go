package main

// Fdroidcl - F-Droid desktop client
// Homepage: https://github.com/mvdan/fdroidcl

import (
	"fmt"
	
	"os/exec"
)

func installFdroidcl() {
	// Método 1: Descargar y extraer .tar.gz
	fdroidcl_tar_url := "https://github.com/mvdan/fdroidcl/archive/refs/tags/v0.7.0.tar.gz"
	fdroidcl_cmd_tar := exec.Command("curl", "-L", fdroidcl_tar_url, "-o", "package.tar.gz")
	err := fdroidcl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fdroidcl_zip_url := "https://github.com/mvdan/fdroidcl/archive/refs/tags/v0.7.0.zip"
	fdroidcl_cmd_zip := exec.Command("curl", "-L", fdroidcl_zip_url, "-o", "package.zip")
	err = fdroidcl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fdroidcl_bin_url := "https://github.com/mvdan/fdroidcl/archive/refs/tags/v0.7.0.bin"
	fdroidcl_cmd_bin := exec.Command("curl", "-L", fdroidcl_bin_url, "-o", "binary.bin")
	err = fdroidcl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fdroidcl_src_url := "https://github.com/mvdan/fdroidcl/archive/refs/tags/v0.7.0.src.tar.gz"
	fdroidcl_cmd_src := exec.Command("curl", "-L", fdroidcl_src_url, "-o", "source.tar.gz")
	err = fdroidcl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fdroidcl_cmd_direct := exec.Command("./binary")
	err = fdroidcl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
