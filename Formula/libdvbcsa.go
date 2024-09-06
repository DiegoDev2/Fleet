package main

// Libdvbcsa - Free implementation of the DVB Common Scrambling Algorithm
// Homepage: https://www.videolan.org/developers/libdvbcsa.html

import (
	"fmt"
	
	"os/exec"
)

func installLibdvbcsa() {
	// Método 1: Descargar y extraer .tar.gz
	libdvbcsa_tar_url := "https://get.videolan.org/libdvbcsa/1.1.0/libdvbcsa-1.1.0.tar.gz"
	libdvbcsa_cmd_tar := exec.Command("curl", "-L", libdvbcsa_tar_url, "-o", "package.tar.gz")
	err := libdvbcsa_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libdvbcsa_zip_url := "https://get.videolan.org/libdvbcsa/1.1.0/libdvbcsa-1.1.0.zip"
	libdvbcsa_cmd_zip := exec.Command("curl", "-L", libdvbcsa_zip_url, "-o", "package.zip")
	err = libdvbcsa_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libdvbcsa_bin_url := "https://get.videolan.org/libdvbcsa/1.1.0/libdvbcsa-1.1.0.bin"
	libdvbcsa_cmd_bin := exec.Command("curl", "-L", libdvbcsa_bin_url, "-o", "binary.bin")
	err = libdvbcsa_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libdvbcsa_src_url := "https://get.videolan.org/libdvbcsa/1.1.0/libdvbcsa-1.1.0.src.tar.gz"
	libdvbcsa_cmd_src := exec.Command("curl", "-L", libdvbcsa_src_url, "-o", "source.tar.gz")
	err = libdvbcsa_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libdvbcsa_cmd_direct := exec.Command("./binary")
	err = libdvbcsa_cmd_direct.Run()
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
