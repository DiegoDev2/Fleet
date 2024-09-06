package main

// Libplctag - Portable and simple API for accessing AB PLC data over Ethernet
// Homepage: https://github.com/libplctag/libplctag

import (
	"fmt"
	
	"os/exec"
)

func installLibplctag() {
	// Método 1: Descargar y extraer .tar.gz
	libplctag_tar_url := "https://github.com/libplctag/libplctag/archive/refs/tags/v2.6.3.tar.gz"
	libplctag_cmd_tar := exec.Command("curl", "-L", libplctag_tar_url, "-o", "package.tar.gz")
	err := libplctag_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libplctag_zip_url := "https://github.com/libplctag/libplctag/archive/refs/tags/v2.6.3.zip"
	libplctag_cmd_zip := exec.Command("curl", "-L", libplctag_zip_url, "-o", "package.zip")
	err = libplctag_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libplctag_bin_url := "https://github.com/libplctag/libplctag/archive/refs/tags/v2.6.3.bin"
	libplctag_cmd_bin := exec.Command("curl", "-L", libplctag_bin_url, "-o", "binary.bin")
	err = libplctag_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libplctag_src_url := "https://github.com/libplctag/libplctag/archive/refs/tags/v2.6.3.src.tar.gz"
	libplctag_cmd_src := exec.Command("curl", "-L", libplctag_src_url, "-o", "source.tar.gz")
	err = libplctag_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libplctag_cmd_direct := exec.Command("./binary")
	err = libplctag_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
