package main

// Libmodbus - Portable modbus library
// Homepage: https://libmodbus.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibmodbus() {
	// Método 1: Descargar y extraer .tar.gz
	libmodbus_tar_url := "https://github.com/stephane/libmodbus/archive/refs/tags/v3.1.10.tar.gz"
	libmodbus_cmd_tar := exec.Command("curl", "-L", libmodbus_tar_url, "-o", "package.tar.gz")
	err := libmodbus_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libmodbus_zip_url := "https://github.com/stephane/libmodbus/archive/refs/tags/v3.1.10.zip"
	libmodbus_cmd_zip := exec.Command("curl", "-L", libmodbus_zip_url, "-o", "package.zip")
	err = libmodbus_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libmodbus_bin_url := "https://github.com/stephane/libmodbus/archive/refs/tags/v3.1.10.bin"
	libmodbus_cmd_bin := exec.Command("curl", "-L", libmodbus_bin_url, "-o", "binary.bin")
	err = libmodbus_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libmodbus_src_url := "https://github.com/stephane/libmodbus/archive/refs/tags/v3.1.10.src.tar.gz"
	libmodbus_cmd_src := exec.Command("curl", "-L", libmodbus_src_url, "-o", "source.tar.gz")
	err = libmodbus_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libmodbus_cmd_direct := exec.Command("./binary")
	err = libmodbus_cmd_direct.Run()
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
