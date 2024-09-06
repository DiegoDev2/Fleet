package main

// Libconfig - Configuration file processing library
// Homepage: https://hyperrealm.github.io/libconfig/

import (
	"fmt"
	
	"os/exec"
)

func installLibconfig() {
	// Método 1: Descargar y extraer .tar.gz
	libconfig_tar_url := "https://github.com/hyperrealm/libconfig/releases/download/v1.7.3/libconfig-1.7.3.tar.gz"
	libconfig_cmd_tar := exec.Command("curl", "-L", libconfig_tar_url, "-o", "package.tar.gz")
	err := libconfig_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libconfig_zip_url := "https://github.com/hyperrealm/libconfig/releases/download/v1.7.3/libconfig-1.7.3.zip"
	libconfig_cmd_zip := exec.Command("curl", "-L", libconfig_zip_url, "-o", "package.zip")
	err = libconfig_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libconfig_bin_url := "https://github.com/hyperrealm/libconfig/releases/download/v1.7.3/libconfig-1.7.3.bin"
	libconfig_cmd_bin := exec.Command("curl", "-L", libconfig_bin_url, "-o", "binary.bin")
	err = libconfig_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libconfig_src_url := "https://github.com/hyperrealm/libconfig/releases/download/v1.7.3/libconfig-1.7.3.src.tar.gz"
	libconfig_cmd_src := exec.Command("curl", "-L", libconfig_src_url, "-o", "source.tar.gz")
	err = libconfig_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libconfig_cmd_direct := exec.Command("./binary")
	err = libconfig_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: texinfo")
exec.Command("latte", "install", "texinfo")
}
