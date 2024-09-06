package main

// Libexosip - Toolkit for eXosip2
// Homepage: https://savannah.nongnu.org/projects/exosip

import (
	"fmt"
	
	"os/exec"
)

func installLibexosip() {
	// Método 1: Descargar y extraer .tar.gz
	libexosip_tar_url := "https://download.savannah.gnu.org/releases/exosip/libexosip2-5.3.0.tar.gz"
	libexosip_cmd_tar := exec.Command("curl", "-L", libexosip_tar_url, "-o", "package.tar.gz")
	err := libexosip_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libexosip_zip_url := "https://download.savannah.gnu.org/releases/exosip/libexosip2-5.3.0.zip"
	libexosip_cmd_zip := exec.Command("curl", "-L", libexosip_zip_url, "-o", "package.zip")
	err = libexosip_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libexosip_bin_url := "https://download.savannah.gnu.org/releases/exosip/libexosip2-5.3.0.bin"
	libexosip_cmd_bin := exec.Command("curl", "-L", libexosip_bin_url, "-o", "binary.bin")
	err = libexosip_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libexosip_src_url := "https://download.savannah.gnu.org/releases/exosip/libexosip2-5.3.0.src.tar.gz"
	libexosip_cmd_src := exec.Command("curl", "-L", libexosip_src_url, "-o", "source.tar.gz")
	err = libexosip_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libexosip_cmd_direct := exec.Command("./binary")
	err = libexosip_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: c-ares")
exec.Command("latte", "install", "c-ares")
	fmt.Println("Instalando dependencia: libosip")
exec.Command("latte", "install", "libosip")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
