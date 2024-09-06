package main

// Licensefinder - Find licenses for your project's dependencies
// Homepage: https://github.com/pivotal/LicenseFinder

import (
	"fmt"
	
	"os/exec"
)

func installLicensefinder() {
	// Método 1: Descargar y extraer .tar.gz
	licensefinder_tar_url := "https://github.com/pivotal/LicenseFinder.git"
	licensefinder_cmd_tar := exec.Command("curl", "-L", licensefinder_tar_url, "-o", "package.tar.gz")
	err := licensefinder_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	licensefinder_zip_url := "https://github.com/pivotal/LicenseFinder.git"
	licensefinder_cmd_zip := exec.Command("curl", "-L", licensefinder_zip_url, "-o", "package.zip")
	err = licensefinder_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	licensefinder_bin_url := "https://github.com/pivotal/LicenseFinder.git"
	licensefinder_cmd_bin := exec.Command("curl", "-L", licensefinder_bin_url, "-o", "binary.bin")
	err = licensefinder_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	licensefinder_src_url := "https://github.com/pivotal/LicenseFinder.git"
	licensefinder_cmd_src := exec.Command("curl", "-L", licensefinder_src_url, "-o", "source.tar.gz")
	err = licensefinder_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	licensefinder_cmd_direct := exec.Command("./binary")
	err = licensefinder_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ruby")
	exec.Command("latte", "install", "ruby").Run()
}
