package main

// Ola - Open Lighting Architecture for lighting control information
// Homepage: https://www.openlighting.org/ola/

import (
	"fmt"
	
	"os/exec"
)

func installOla() {
	// Método 1: Descargar y extraer .tar.gz
	ola_tar_url := "https://github.com/OpenLightingProject/ola/releases/download/0.10.9/ola-0.10.9.tar.gz"
	ola_cmd_tar := exec.Command("curl", "-L", ola_tar_url, "-o", "package.tar.gz")
	err := ola_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ola_zip_url := "https://github.com/OpenLightingProject/ola/releases/download/0.10.9/ola-0.10.9.zip"
	ola_cmd_zip := exec.Command("curl", "-L", ola_zip_url, "-o", "package.zip")
	err = ola_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ola_bin_url := "https://github.com/OpenLightingProject/ola/releases/download/0.10.9/ola-0.10.9.bin"
	ola_cmd_bin := exec.Command("curl", "-L", ola_bin_url, "-o", "binary.bin")
	err = ola_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ola_src_url := "https://github.com/OpenLightingProject/ola/releases/download/0.10.9/ola-0.10.9.src.tar.gz"
	ola_cmd_src := exec.Command("curl", "-L", ola_src_url, "-o", "source.tar.gz")
	err = ola_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ola_cmd_direct := exec.Command("./binary")
	err = ola_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: liblo")
	exec.Command("latte", "install", "liblo").Run()
	fmt.Println("Instalando dependencia: libmicrohttpd")
	exec.Command("latte", "install", "libmicrohttpd").Run()
	fmt.Println("Instalando dependencia: libusb")
	exec.Command("latte", "install", "libusb").Run()
	fmt.Println("Instalando dependencia: numpy")
	exec.Command("latte", "install", "numpy").Run()
	fmt.Println("Instalando dependencia: protobuf@21")
	exec.Command("latte", "install", "protobuf@21").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: util-linux")
	exec.Command("latte", "install", "util-linux").Run()
}
