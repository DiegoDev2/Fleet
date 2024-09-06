package main

// Robodoc - Source code documentation tool
// Homepage: https://rfsber.home.xs4all.nl/Robo/index.html

import (
	"fmt"
	
	"os/exec"
)

func installRobodoc() {
	// Método 1: Descargar y extraer .tar.gz
	robodoc_tar_url := "https://rfsber.home.xs4all.nl/Robo/archives/robodoc-4.99.44.tar.bz2"
	robodoc_cmd_tar := exec.Command("curl", "-L", robodoc_tar_url, "-o", "package.tar.gz")
	err := robodoc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	robodoc_zip_url := "https://rfsber.home.xs4all.nl/Robo/archives/robodoc-4.99.44.tar.bz2"
	robodoc_cmd_zip := exec.Command("curl", "-L", robodoc_zip_url, "-o", "package.zip")
	err = robodoc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	robodoc_bin_url := "https://rfsber.home.xs4all.nl/Robo/archives/robodoc-4.99.44.tar.bz2"
	robodoc_cmd_bin := exec.Command("curl", "-L", robodoc_bin_url, "-o", "binary.bin")
	err = robodoc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	robodoc_src_url := "https://rfsber.home.xs4all.nl/Robo/archives/robodoc-4.99.44.tar.bz2"
	robodoc_cmd_src := exec.Command("curl", "-L", robodoc_src_url, "-o", "source.tar.gz")
	err = robodoc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	robodoc_cmd_direct := exec.Command("./binary")
	err = robodoc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
}
