package main

// Libxo - Allows an application to generate text, XML, JSON, and HTML output
// Homepage: https://juniper.github.io/libxo/libxo-manual.html

import (
	"fmt"
	
	"os/exec"
)

func installLibxo() {
	// Método 1: Descargar y extraer .tar.gz
	libxo_tar_url := "https://github.com/Juniper/libxo/releases/download/1.7.5/libxo-1.7.5.tar.gz"
	libxo_cmd_tar := exec.Command("curl", "-L", libxo_tar_url, "-o", "package.tar.gz")
	err := libxo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libxo_zip_url := "https://github.com/Juniper/libxo/releases/download/1.7.5/libxo-1.7.5.zip"
	libxo_cmd_zip := exec.Command("curl", "-L", libxo_zip_url, "-o", "package.zip")
	err = libxo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libxo_bin_url := "https://github.com/Juniper/libxo/releases/download/1.7.5/libxo-1.7.5.bin"
	libxo_cmd_bin := exec.Command("curl", "-L", libxo_bin_url, "-o", "binary.bin")
	err = libxo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libxo_src_url := "https://github.com/Juniper/libxo/releases/download/1.7.5/libxo-1.7.5.src.tar.gz"
	libxo_cmd_src := exec.Command("curl", "-L", libxo_src_url, "-o", "source.tar.gz")
	err = libxo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libxo_cmd_direct := exec.Command("./binary")
	err = libxo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
}
