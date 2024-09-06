package main

// Libdpp - C++ Discord API Bot Library
// Homepage: https://github.com/brainboxdotcc/DPP

import (
	"fmt"
	
	"os/exec"
)

func installLibdpp() {
	// Método 1: Descargar y extraer .tar.gz
	libdpp_tar_url := "https://github.com/brainboxdotcc/DPP/releases/download/v10.0.30/DPP-10.0.30.tar.gz"
	libdpp_cmd_tar := exec.Command("curl", "-L", libdpp_tar_url, "-o", "package.tar.gz")
	err := libdpp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libdpp_zip_url := "https://github.com/brainboxdotcc/DPP/releases/download/v10.0.30/DPP-10.0.30.zip"
	libdpp_cmd_zip := exec.Command("curl", "-L", libdpp_zip_url, "-o", "package.zip")
	err = libdpp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libdpp_bin_url := "https://github.com/brainboxdotcc/DPP/releases/download/v10.0.30/DPP-10.0.30.bin"
	libdpp_cmd_bin := exec.Command("curl", "-L", libdpp_bin_url, "-o", "binary.bin")
	err = libdpp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libdpp_src_url := "https://github.com/brainboxdotcc/DPP/releases/download/v10.0.30/DPP-10.0.30.src.tar.gz"
	libdpp_cmd_src := exec.Command("curl", "-L", libdpp_src_url, "-o", "source.tar.gz")
	err = libdpp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libdpp_cmd_direct := exec.Command("./binary")
	err = libdpp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: libsodium")
	exec.Command("latte", "install", "libsodium").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: opus")
	exec.Command("latte", "install", "opus").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
}
