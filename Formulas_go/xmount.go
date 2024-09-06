package main

// Xmount - Convert between multiple input & output disk image types
// Homepage: https://www.sits.lu/xmount

import (
	"fmt"
	
	"os/exec"
)

func installXmount() {
	// Método 1: Descargar y extraer .tar.gz
	xmount_tar_url := "https://code.sits.lu/foss/xmount/-/archive/1.1.1/xmount-1.1.1.tar.bz2"
	xmount_cmd_tar := exec.Command("curl", "-L", xmount_tar_url, "-o", "package.tar.gz")
	err := xmount_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xmount_zip_url := "https://code.sits.lu/foss/xmount/-/archive/1.1.1/xmount-1.1.1.tar.bz2"
	xmount_cmd_zip := exec.Command("curl", "-L", xmount_zip_url, "-o", "package.zip")
	err = xmount_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xmount_bin_url := "https://code.sits.lu/foss/xmount/-/archive/1.1.1/xmount-1.1.1.tar.bz2"
	xmount_cmd_bin := exec.Command("curl", "-L", xmount_bin_url, "-o", "binary.bin")
	err = xmount_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xmount_src_url := "https://code.sits.lu/foss/xmount/-/archive/1.1.1/xmount-1.1.1.tar.bz2"
	xmount_cmd_src := exec.Command("curl", "-L", xmount_src_url, "-o", "source.tar.gz")
	err = xmount_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xmount_cmd_direct := exec.Command("./binary")
	err = xmount_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: afflib")
exec.Command("latte", "install", "afflib")
	fmt.Println("Instalando dependencia: libewf")
exec.Command("latte", "install", "libewf")
	fmt.Println("Instalando dependencia: libfuse@2")
exec.Command("latte", "install", "libfuse@2")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
