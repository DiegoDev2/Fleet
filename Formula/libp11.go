package main

// Libp11 - PKCS#11 wrapper library in C
// Homepage: https://github.com/OpenSC/libp11/wiki

import (
	"fmt"
	
	"os/exec"
)

func installLibp11() {
	// Método 1: Descargar y extraer .tar.gz
	libp11_tar_url := "https://github.com/OpenSC/libp11/releases/download/libp11-0.4.12/libp11-0.4.12.tar.gz"
	libp11_cmd_tar := exec.Command("curl", "-L", libp11_tar_url, "-o", "package.tar.gz")
	err := libp11_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libp11_zip_url := "https://github.com/OpenSC/libp11/releases/download/libp11-0.4.12/libp11-0.4.12.zip"
	libp11_cmd_zip := exec.Command("curl", "-L", libp11_zip_url, "-o", "package.zip")
	err = libp11_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libp11_bin_url := "https://github.com/OpenSC/libp11/releases/download/libp11-0.4.12/libp11-0.4.12.bin"
	libp11_cmd_bin := exec.Command("curl", "-L", libp11_bin_url, "-o", "binary.bin")
	err = libp11_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libp11_src_url := "https://github.com/OpenSC/libp11/releases/download/libp11-0.4.12/libp11-0.4.12.src.tar.gz"
	libp11_cmd_src := exec.Command("curl", "-L", libp11_src_url, "-o", "source.tar.gz")
	err = libp11_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libp11_cmd_direct := exec.Command("./binary")
	err = libp11_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
