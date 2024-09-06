package main

// Nut - Network UPS Tools: Support for various power devices
// Homepage: https://networkupstools.org/

import (
	"fmt"
	
	"os/exec"
)

func installNut() {
	// Método 1: Descargar y extraer .tar.gz
	nut_tar_url := "https://github.com/networkupstools/nut/releases/download/v2.8.2/nut-2.8.2.tar.gz"
	nut_cmd_tar := exec.Command("curl", "-L", nut_tar_url, "-o", "package.tar.gz")
	err := nut_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nut_zip_url := "https://github.com/networkupstools/nut/releases/download/v2.8.2/nut-2.8.2.zip"
	nut_cmd_zip := exec.Command("curl", "-L", nut_zip_url, "-o", "package.zip")
	err = nut_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nut_bin_url := "https://github.com/networkupstools/nut/releases/download/v2.8.2/nut-2.8.2.bin"
	nut_cmd_bin := exec.Command("curl", "-L", nut_bin_url, "-o", "binary.bin")
	err = nut_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nut_src_url := "https://github.com/networkupstools/nut/releases/download/v2.8.2/nut-2.8.2.src.tar.gz"
	nut_cmd_src := exec.Command("curl", "-L", nut_src_url, "-o", "source.tar.gz")
	err = nut_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nut_cmd_direct := exec.Command("./binary")
	err = nut_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: asciidoc")
exec.Command("latte", "install", "asciidoc")
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libusb")
exec.Command("latte", "install", "libusb")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
