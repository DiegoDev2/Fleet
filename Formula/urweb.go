package main

// Urweb - Ur/Web programming language
// Homepage: http://www.impredicative.com/ur/

import (
	"fmt"
	
	"os/exec"
)

func installUrweb() {
	// Método 1: Descargar y extraer .tar.gz
	urweb_tar_url := "https://github.com/urweb/urweb/releases/download/20200209/urweb-20200209.tar.gz"
	urweb_cmd_tar := exec.Command("curl", "-L", urweb_tar_url, "-o", "package.tar.gz")
	err := urweb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	urweb_zip_url := "https://github.com/urweb/urweb/releases/download/20200209/urweb-20200209.zip"
	urweb_cmd_zip := exec.Command("curl", "-L", urweb_zip_url, "-o", "package.zip")
	err = urweb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	urweb_bin_url := "https://github.com/urweb/urweb/releases/download/20200209/urweb-20200209.bin"
	urweb_cmd_bin := exec.Command("curl", "-L", urweb_bin_url, "-o", "binary.bin")
	err = urweb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	urweb_src_url := "https://github.com/urweb/urweb/releases/download/20200209/urweb-20200209.src.tar.gz"
	urweb_cmd_src := exec.Command("curl", "-L", urweb_src_url, "-o", "source.tar.gz")
	err = urweb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	urweb_cmd_direct := exec.Command("./binary")
	err = urweb_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: mlton")
	exec.Command("latte", "install", "mlton").Run()
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
	fmt.Println("Instalando dependencia: icu4c")
	exec.Command("latte", "install", "icu4c").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
