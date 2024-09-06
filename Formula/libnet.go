package main

// Libnet - C library for creating IP packets
// Homepage: https://github.com/libnet/libnet

import (
	"fmt"
	
	"os/exec"
)

func installLibnet() {
	// Método 1: Descargar y extraer .tar.gz
	libnet_tar_url := "https://github.com/libnet/libnet/releases/download/v1.3/libnet-1.3.tar.gz"
	libnet_cmd_tar := exec.Command("curl", "-L", libnet_tar_url, "-o", "package.tar.gz")
	err := libnet_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libnet_zip_url := "https://github.com/libnet/libnet/releases/download/v1.3/libnet-1.3.zip"
	libnet_cmd_zip := exec.Command("curl", "-L", libnet_zip_url, "-o", "package.zip")
	err = libnet_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libnet_bin_url := "https://github.com/libnet/libnet/releases/download/v1.3/libnet-1.3.bin"
	libnet_cmd_bin := exec.Command("curl", "-L", libnet_bin_url, "-o", "binary.bin")
	err = libnet_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libnet_src_url := "https://github.com/libnet/libnet/releases/download/v1.3/libnet-1.3.src.tar.gz"
	libnet_cmd_src := exec.Command("curl", "-L", libnet_src_url, "-o", "source.tar.gz")
	err = libnet_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libnet_cmd_direct := exec.Command("./binary")
	err = libnet_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: doxygen")
	exec.Command("latte", "install", "doxygen").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
}
