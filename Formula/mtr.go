package main

// Mtr - 'traceroute' and 'ping' in a single tool
// Homepage: https://www.bitwizard.nl/mtr/

import (
	"fmt"
	
	"os/exec"
)

func installMtr() {
	// Método 1: Descargar y extraer .tar.gz
	mtr_tar_url := "https://github.com/traviscross/mtr/archive/refs/tags/v0.95.tar.gz"
	mtr_cmd_tar := exec.Command("curl", "-L", mtr_tar_url, "-o", "package.tar.gz")
	err := mtr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mtr_zip_url := "https://github.com/traviscross/mtr/archive/refs/tags/v0.95.zip"
	mtr_cmd_zip := exec.Command("curl", "-L", mtr_zip_url, "-o", "package.zip")
	err = mtr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mtr_bin_url := "https://github.com/traviscross/mtr/archive/refs/tags/v0.95.bin"
	mtr_cmd_bin := exec.Command("curl", "-L", mtr_bin_url, "-o", "binary.bin")
	err = mtr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mtr_src_url := "https://github.com/traviscross/mtr/archive/refs/tags/v0.95.src.tar.gz"
	mtr_cmd_src := exec.Command("curl", "-L", mtr_src_url, "-o", "source.tar.gz")
	err = mtr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mtr_cmd_direct := exec.Command("./binary")
	err = mtr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: jansson")
	exec.Command("latte", "install", "jansson").Run()
}
