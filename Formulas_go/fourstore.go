package main

// Fourstore - Efficient, stable RDF database
// Homepage: https://github.com/4store/4store

import (
	"fmt"
	
	"os/exec"
)

func installFourstore() {
	// Método 1: Descargar y extraer .tar.gz
	fourstore_tar_url := "https://github.com/4store/4store/archive/refs/tags/v1.1.7.tar.gz"
	fourstore_cmd_tar := exec.Command("curl", "-L", fourstore_tar_url, "-o", "package.tar.gz")
	err := fourstore_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fourstore_zip_url := "https://github.com/4store/4store/archive/refs/tags/v1.1.7.zip"
	fourstore_cmd_zip := exec.Command("curl", "-L", fourstore_zip_url, "-o", "package.zip")
	err = fourstore_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fourstore_bin_url := "https://github.com/4store/4store/archive/refs/tags/v1.1.7.bin"
	fourstore_cmd_bin := exec.Command("curl", "-L", fourstore_bin_url, "-o", "binary.bin")
	err = fourstore_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fourstore_src_url := "https://github.com/4store/4store/archive/refs/tags/v1.1.7.src.tar.gz"
	fourstore_cmd_src := exec.Command("curl", "-L", fourstore_src_url, "-o", "source.tar.gz")
	err = fourstore_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fourstore_cmd_direct := exec.Command("./binary")
	err = fourstore_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: dbus")
exec.Command("latte", "install", "dbus")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: pcre")
exec.Command("latte", "install", "pcre")
	fmt.Println("Instalando dependencia: raptor")
exec.Command("latte", "install", "raptor")
	fmt.Println("Instalando dependencia: rasqal")
exec.Command("latte", "install", "rasqal")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
	fmt.Println("Instalando dependencia: util-linux")
exec.Command("latte", "install", "util-linux")
}
