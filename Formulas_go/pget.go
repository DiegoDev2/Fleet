package main

// Pget - File download client
// Homepage: https://github.com/Code-Hex/pget

import (
	"fmt"
	
	"os/exec"
)

func installPget() {
	// Método 1: Descargar y extraer .tar.gz
	pget_tar_url := "https://github.com/Code-Hex/pget/archive/refs/tags/v0.2.1.tar.gz"
	pget_cmd_tar := exec.Command("curl", "-L", pget_tar_url, "-o", "package.tar.gz")
	err := pget_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pget_zip_url := "https://github.com/Code-Hex/pget/archive/refs/tags/v0.2.1.zip"
	pget_cmd_zip := exec.Command("curl", "-L", pget_zip_url, "-o", "package.zip")
	err = pget_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pget_bin_url := "https://github.com/Code-Hex/pget/archive/refs/tags/v0.2.1.bin"
	pget_cmd_bin := exec.Command("curl", "-L", pget_bin_url, "-o", "binary.bin")
	err = pget_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pget_src_url := "https://github.com/Code-Hex/pget/archive/refs/tags/v0.2.1.src.tar.gz"
	pget_cmd_src := exec.Command("curl", "-L", pget_src_url, "-o", "source.tar.gz")
	err = pget_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pget_cmd_direct := exec.Command("./binary")
	err = pget_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
