package main

// Rhit - Nginx log explorer
// Homepage: https://dystroy.org/rhit/

import (
	"fmt"
	
	"os/exec"
)

func installRhit() {
	// Método 1: Descargar y extraer .tar.gz
	rhit_tar_url := "https://github.com/Canop/rhit/archive/refs/tags/v2.0.1.tar.gz"
	rhit_cmd_tar := exec.Command("curl", "-L", rhit_tar_url, "-o", "package.tar.gz")
	err := rhit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rhit_zip_url := "https://github.com/Canop/rhit/archive/refs/tags/v2.0.1.zip"
	rhit_cmd_zip := exec.Command("curl", "-L", rhit_zip_url, "-o", "package.zip")
	err = rhit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rhit_bin_url := "https://github.com/Canop/rhit/archive/refs/tags/v2.0.1.bin"
	rhit_cmd_bin := exec.Command("curl", "-L", rhit_bin_url, "-o", "binary.bin")
	err = rhit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rhit_src_url := "https://github.com/Canop/rhit/archive/refs/tags/v2.0.1.src.tar.gz"
	rhit_cmd_src := exec.Command("curl", "-L", rhit_src_url, "-o", "source.tar.gz")
	err = rhit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rhit_cmd_direct := exec.Command("./binary")
	err = rhit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
