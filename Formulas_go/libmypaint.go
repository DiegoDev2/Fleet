package main

// Libmypaint - MyPaint brush engine library
// Homepage: https://github.com/mypaint/libmypaint/wiki

import (
	"fmt"
	
	"os/exec"
)

func installLibmypaint() {
	// Método 1: Descargar y extraer .tar.gz
	libmypaint_tar_url := "https://github.com/mypaint/libmypaint/releases/download/v1.6.1/libmypaint-1.6.1.tar.xz"
	libmypaint_cmd_tar := exec.Command("curl", "-L", libmypaint_tar_url, "-o", "package.tar.gz")
	err := libmypaint_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libmypaint_zip_url := "https://github.com/mypaint/libmypaint/releases/download/v1.6.1/libmypaint-1.6.1.tar.xz"
	libmypaint_cmd_zip := exec.Command("curl", "-L", libmypaint_zip_url, "-o", "package.zip")
	err = libmypaint_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libmypaint_bin_url := "https://github.com/mypaint/libmypaint/releases/download/v1.6.1/libmypaint-1.6.1.tar.xz"
	libmypaint_cmd_bin := exec.Command("curl", "-L", libmypaint_bin_url, "-o", "binary.bin")
	err = libmypaint_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libmypaint_src_url := "https://github.com/mypaint/libmypaint/releases/download/v1.6.1/libmypaint-1.6.1.tar.xz"
	libmypaint_cmd_src := exec.Command("curl", "-L", libmypaint_src_url, "-o", "source.tar.gz")
	err = libmypaint_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libmypaint_cmd_direct := exec.Command("./binary")
	err = libmypaint_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: intltool")
exec.Command("latte", "install", "intltool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: json-c")
exec.Command("latte", "install", "json-c")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: perl-xml-parser")
exec.Command("latte", "install", "perl-xml-parser")
}
