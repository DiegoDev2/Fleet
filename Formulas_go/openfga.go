package main

// Openfga - High performance and flexible authorization/permission engine
// Homepage: https://openfga.dev/

import (
	"fmt"
	
	"os/exec"
)

func installOpenfga() {
	// Método 1: Descargar y extraer .tar.gz
	openfga_tar_url := "https://github.com/openfga/openfga/archive/refs/tags/v1.6.0.tar.gz"
	openfga_cmd_tar := exec.Command("curl", "-L", openfga_tar_url, "-o", "package.tar.gz")
	err := openfga_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openfga_zip_url := "https://github.com/openfga/openfga/archive/refs/tags/v1.6.0.zip"
	openfga_cmd_zip := exec.Command("curl", "-L", openfga_zip_url, "-o", "package.zip")
	err = openfga_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openfga_bin_url := "https://github.com/openfga/openfga/archive/refs/tags/v1.6.0.bin"
	openfga_cmd_bin := exec.Command("curl", "-L", openfga_bin_url, "-o", "binary.bin")
	err = openfga_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openfga_src_url := "https://github.com/openfga/openfga/archive/refs/tags/v1.6.0.src.tar.gz"
	openfga_cmd_src := exec.Command("curl", "-L", openfga_src_url, "-o", "source.tar.gz")
	err = openfga_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openfga_cmd_direct := exec.Command("./binary")
	err = openfga_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
