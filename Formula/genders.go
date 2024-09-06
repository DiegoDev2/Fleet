package main

// Genders - Static cluster configuration database for cluster management
// Homepage: https://github.com/chaos/genders

import (
	"fmt"
	
	"os/exec"
)

func installGenders() {
	// Método 1: Descargar y extraer .tar.gz
	genders_tar_url := "https://github.com/chaos/genders/archive/refs/tags/genders-1-29-1.tar.gz"
	genders_cmd_tar := exec.Command("curl", "-L", genders_tar_url, "-o", "package.tar.gz")
	err := genders_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	genders_zip_url := "https://github.com/chaos/genders/archive/refs/tags/genders-1-29-1.zip"
	genders_cmd_zip := exec.Command("curl", "-L", genders_zip_url, "-o", "package.zip")
	err = genders_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	genders_bin_url := "https://github.com/chaos/genders/archive/refs/tags/genders-1-29-1.bin"
	genders_cmd_bin := exec.Command("curl", "-L", genders_bin_url, "-o", "binary.bin")
	err = genders_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	genders_src_url := "https://github.com/chaos/genders/archive/refs/tags/genders-1-29-1.src.tar.gz"
	genders_cmd_src := exec.Command("curl", "-L", genders_src_url, "-o", "source.tar.gz")
	err = genders_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	genders_cmd_direct := exec.Command("./binary")
	err = genders_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python-setuptools")
	exec.Command("latte", "install", "python-setuptools").Run()
}
