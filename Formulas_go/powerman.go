package main

// Powerman - Control (remotely and in parallel) switched power distribution units
// Homepage: https://github.com/chaos/powerman

import (
	"fmt"
	
	"os/exec"
)

func installPowerman() {
	// Método 1: Descargar y extraer .tar.gz
	powerman_tar_url := "https://github.com/chaos/powerman/releases/download/v2.4.3/powerman-2.4.3.tar.gz"
	powerman_cmd_tar := exec.Command("curl", "-L", powerman_tar_url, "-o", "package.tar.gz")
	err := powerman_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	powerman_zip_url := "https://github.com/chaos/powerman/releases/download/v2.4.3/powerman-2.4.3.zip"
	powerman_cmd_zip := exec.Command("curl", "-L", powerman_zip_url, "-o", "package.zip")
	err = powerman_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	powerman_bin_url := "https://github.com/chaos/powerman/releases/download/v2.4.3/powerman-2.4.3.bin"
	powerman_cmd_bin := exec.Command("curl", "-L", powerman_bin_url, "-o", "binary.bin")
	err = powerman_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	powerman_src_url := "https://github.com/chaos/powerman/releases/download/v2.4.3/powerman-2.4.3.src.tar.gz"
	powerman_cmd_src := exec.Command("curl", "-L", powerman_src_url, "-o", "source.tar.gz")
	err = powerman_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	powerman_cmd_direct := exec.Command("./binary")
	err = powerman_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: curl")
exec.Command("latte", "install", "curl")
	fmt.Println("Instalando dependencia: jansson")
exec.Command("latte", "install", "jansson")
}
