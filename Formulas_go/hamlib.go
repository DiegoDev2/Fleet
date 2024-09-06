package main

// Hamlib - Ham radio control libraries
// Homepage: http://www.hamlib.org/

import (
	"fmt"
	
	"os/exec"
)

func installHamlib() {
	// Método 1: Descargar y extraer .tar.gz
	hamlib_tar_url := "https://github.com/Hamlib/Hamlib/releases/download/4.5.5/hamlib-4.5.5.tar.gz"
	hamlib_cmd_tar := exec.Command("curl", "-L", hamlib_tar_url, "-o", "package.tar.gz")
	err := hamlib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hamlib_zip_url := "https://github.com/Hamlib/Hamlib/releases/download/4.5.5/hamlib-4.5.5.zip"
	hamlib_cmd_zip := exec.Command("curl", "-L", hamlib_zip_url, "-o", "package.zip")
	err = hamlib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hamlib_bin_url := "https://github.com/Hamlib/Hamlib/releases/download/4.5.5/hamlib-4.5.5.bin"
	hamlib_cmd_bin := exec.Command("curl", "-L", hamlib_bin_url, "-o", "binary.bin")
	err = hamlib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hamlib_src_url := "https://github.com/Hamlib/Hamlib/releases/download/4.5.5/hamlib-4.5.5.src.tar.gz"
	hamlib_cmd_src := exec.Command("curl", "-L", hamlib_src_url, "-o", "source.tar.gz")
	err = hamlib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hamlib_cmd_direct := exec.Command("./binary")
	err = hamlib_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: libusb")
exec.Command("latte", "install", "libusb")
	fmt.Println("Instalando dependencia: libusb-compat")
exec.Command("latte", "install", "libusb-compat")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
