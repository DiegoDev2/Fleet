package main

// Ctail - Tool for operating tail across large clusters of machines
// Homepage: https://github.com/pquerna/ctail

import (
	"fmt"
	
	"os/exec"
)

func installCtail() {
	// Método 1: Descargar y extraer .tar.gz
	ctail_tar_url := "https://github.com/pquerna/ctail/archive/refs/tags/ctail-0.1.0.tar.gz"
	ctail_cmd_tar := exec.Command("curl", "-L", ctail_tar_url, "-o", "package.tar.gz")
	err := ctail_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ctail_zip_url := "https://github.com/pquerna/ctail/archive/refs/tags/ctail-0.1.0.zip"
	ctail_cmd_zip := exec.Command("curl", "-L", ctail_zip_url, "-o", "package.zip")
	err = ctail_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ctail_bin_url := "https://github.com/pquerna/ctail/archive/refs/tags/ctail-0.1.0.bin"
	ctail_cmd_bin := exec.Command("curl", "-L", ctail_bin_url, "-o", "binary.bin")
	err = ctail_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ctail_src_url := "https://github.com/pquerna/ctail/archive/refs/tags/ctail-0.1.0.src.tar.gz"
	ctail_cmd_src := exec.Command("curl", "-L", ctail_src_url, "-o", "source.tar.gz")
	err = ctail_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ctail_cmd_direct := exec.Command("./binary")
	err = ctail_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: apr")
	exec.Command("latte", "install", "apr").Run()
	fmt.Println("Instalando dependencia: apr-util")
	exec.Command("latte", "install", "apr-util").Run()
}
