package main

// Osc - Command-line interface to work with an Open Build Service
// Homepage: https://openbuildservice.org

import (
	"fmt"
	
	"os/exec"
)

func installOsc() {
	// Método 1: Descargar y extraer .tar.gz
	osc_tar_url := "https://github.com/openSUSE/osc/archive/refs/tags/1.9.1.tar.gz"
	osc_cmd_tar := exec.Command("curl", "-L", osc_tar_url, "-o", "package.tar.gz")
	err := osc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	osc_zip_url := "https://github.com/openSUSE/osc/archive/refs/tags/1.9.1.zip"
	osc_cmd_zip := exec.Command("curl", "-L", osc_zip_url, "-o", "package.zip")
	err = osc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	osc_bin_url := "https://github.com/openSUSE/osc/archive/refs/tags/1.9.1.bin"
	osc_cmd_bin := exec.Command("curl", "-L", osc_bin_url, "-o", "binary.bin")
	err = osc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	osc_src_url := "https://github.com/openSUSE/osc/archive/refs/tags/1.9.1.src.tar.gz"
	osc_cmd_src := exec.Command("curl", "-L", osc_src_url, "-o", "source.tar.gz")
	err = osc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	osc_cmd_direct := exec.Command("./binary")
	err = osc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cryptography")
exec.Command("latte", "install", "cryptography")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
