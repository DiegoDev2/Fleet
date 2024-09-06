package main

// Libcaca - Convert pixel information into colored ASCII art
// Homepage: http://caca.zoy.org/wiki/libcaca

import (
	"fmt"
	
	"os/exec"
)

func installLibcaca() {
	// Método 1: Descargar y extraer .tar.gz
	libcaca_tar_url := "https://github.com/cacalabs/libcaca/releases/download/v0.99.beta20/libcaca-0.99.beta20.tar.bz2"
	libcaca_cmd_tar := exec.Command("curl", "-L", libcaca_tar_url, "-o", "package.tar.gz")
	err := libcaca_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libcaca_zip_url := "https://github.com/cacalabs/libcaca/releases/download/v0.99.beta20/libcaca-0.99.beta20.tar.bz2"
	libcaca_cmd_zip := exec.Command("curl", "-L", libcaca_zip_url, "-o", "package.zip")
	err = libcaca_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libcaca_bin_url := "https://github.com/cacalabs/libcaca/releases/download/v0.99.beta20/libcaca-0.99.beta20.tar.bz2"
	libcaca_cmd_bin := exec.Command("curl", "-L", libcaca_bin_url, "-o", "binary.bin")
	err = libcaca_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libcaca_src_url := "https://github.com/cacalabs/libcaca/releases/download/v0.99.beta20/libcaca-0.99.beta20.tar.bz2"
	libcaca_cmd_src := exec.Command("curl", "-L", libcaca_src_url, "-o", "source.tar.gz")
	err = libcaca_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libcaca_cmd_direct := exec.Command("./binary")
	err = libcaca_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: imlib2")
	exec.Command("latte", "install", "imlib2").Run()
}
