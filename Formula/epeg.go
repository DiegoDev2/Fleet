package main

// Epeg - JPEG/JPG thumbnail scaling
// Homepage: https://github.com/mattes/epeg

import (
	"fmt"
	
	"os/exec"
)

func installEpeg() {
	// Método 1: Descargar y extraer .tar.gz
	epeg_tar_url := "https://github.com/mattes/epeg/archive/refs/tags/v0.9.3.tar.gz"
	epeg_cmd_tar := exec.Command("curl", "-L", epeg_tar_url, "-o", "package.tar.gz")
	err := epeg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	epeg_zip_url := "https://github.com/mattes/epeg/archive/refs/tags/v0.9.3.zip"
	epeg_cmd_zip := exec.Command("curl", "-L", epeg_zip_url, "-o", "package.zip")
	err = epeg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	epeg_bin_url := "https://github.com/mattes/epeg/archive/refs/tags/v0.9.3.bin"
	epeg_cmd_bin := exec.Command("curl", "-L", epeg_bin_url, "-o", "binary.bin")
	err = epeg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	epeg_src_url := "https://github.com/mattes/epeg/archive/refs/tags/v0.9.3.src.tar.gz"
	epeg_cmd_src := exec.Command("curl", "-L", epeg_src_url, "-o", "source.tar.gz")
	err = epeg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	epeg_cmd_direct := exec.Command("./binary")
	err = epeg_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: libexif")
	exec.Command("latte", "install", "libexif").Run()
}
