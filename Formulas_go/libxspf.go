package main

// Libxspf - C++ library for XSPF playlist reading and writing
// Homepage: https://libspiff.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installLibxspf() {
	// Método 1: Descargar y extraer .tar.gz
	libxspf_tar_url := "https://downloads.sourceforge.net/project/libspiff/Sources/1.2.1/libxspf-1.2.1.tar.bz2"
	libxspf_cmd_tar := exec.Command("curl", "-L", libxspf_tar_url, "-o", "package.tar.gz")
	err := libxspf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libxspf_zip_url := "https://downloads.sourceforge.net/project/libspiff/Sources/1.2.1/libxspf-1.2.1.tar.bz2"
	libxspf_cmd_zip := exec.Command("curl", "-L", libxspf_zip_url, "-o", "package.zip")
	err = libxspf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libxspf_bin_url := "https://downloads.sourceforge.net/project/libspiff/Sources/1.2.1/libxspf-1.2.1.tar.bz2"
	libxspf_cmd_bin := exec.Command("curl", "-L", libxspf_bin_url, "-o", "binary.bin")
	err = libxspf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libxspf_src_url := "https://downloads.sourceforge.net/project/libspiff/Sources/1.2.1/libxspf-1.2.1.tar.bz2"
	libxspf_cmd_src := exec.Command("curl", "-L", libxspf_src_url, "-o", "source.tar.gz")
	err = libxspf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libxspf_cmd_direct := exec.Command("./binary")
	err = libxspf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cpptest")
exec.Command("latte", "install", "cpptest")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: uriparser")
exec.Command("latte", "install", "uriparser")
}
