package main

// Uptimed - Utility to track your highest uptimes
// Homepage: https://github.com/rpodgorny/uptimed/

import (
	"fmt"
	
	"os/exec"
)

func installUptimed() {
	// Método 1: Descargar y extraer .tar.gz
	uptimed_tar_url := "https://github.com/rpodgorny/uptimed/archive/refs/tags/v0.4.6.tar.gz"
	uptimed_cmd_tar := exec.Command("curl", "-L", uptimed_tar_url, "-o", "package.tar.gz")
	err := uptimed_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	uptimed_zip_url := "https://github.com/rpodgorny/uptimed/archive/refs/tags/v0.4.6.zip"
	uptimed_cmd_zip := exec.Command("curl", "-L", uptimed_zip_url, "-o", "package.zip")
	err = uptimed_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	uptimed_bin_url := "https://github.com/rpodgorny/uptimed/archive/refs/tags/v0.4.6.bin"
	uptimed_cmd_bin := exec.Command("curl", "-L", uptimed_bin_url, "-o", "binary.bin")
	err = uptimed_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	uptimed_src_url := "https://github.com/rpodgorny/uptimed/archive/refs/tags/v0.4.6.src.tar.gz"
	uptimed_cmd_src := exec.Command("curl", "-L", uptimed_src_url, "-o", "source.tar.gz")
	err = uptimed_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	uptimed_cmd_direct := exec.Command("./binary")
	err = uptimed_cmd_direct.Run()
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
