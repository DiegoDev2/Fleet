package main

// Liblo - Lightweight Open Sound Control implementation
// Homepage: https://liblo.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installLiblo() {
	// Método 1: Descargar y extraer .tar.gz
	liblo_tar_url := "https://downloads.sourceforge.net/project/liblo/liblo/0.32/liblo-0.32.tar.gz"
	liblo_cmd_tar := exec.Command("curl", "-L", liblo_tar_url, "-o", "package.tar.gz")
	err := liblo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	liblo_zip_url := "https://downloads.sourceforge.net/project/liblo/liblo/0.32/liblo-0.32.zip"
	liblo_cmd_zip := exec.Command("curl", "-L", liblo_zip_url, "-o", "package.zip")
	err = liblo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	liblo_bin_url := "https://downloads.sourceforge.net/project/liblo/liblo/0.32/liblo-0.32.bin"
	liblo_cmd_bin := exec.Command("curl", "-L", liblo_bin_url, "-o", "binary.bin")
	err = liblo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	liblo_src_url := "https://downloads.sourceforge.net/project/liblo/liblo/0.32/liblo-0.32.src.tar.gz"
	liblo_cmd_src := exec.Command("curl", "-L", liblo_src_url, "-o", "source.tar.gz")
	err = liblo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	liblo_cmd_direct := exec.Command("./binary")
	err = liblo_cmd_direct.Run()
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
