package main

// Loudmouth - Lightweight C library for the Jabber protocol
// Homepage: https://mcabber.com

import (
	"fmt"
	
	"os/exec"
)

func installLoudmouth() {
	// Método 1: Descargar y extraer .tar.gz
	loudmouth_tar_url := "https://mcabber.com/files/loudmouth/loudmouth-1.5.4.tar.bz2"
	loudmouth_cmd_tar := exec.Command("curl", "-L", loudmouth_tar_url, "-o", "package.tar.gz")
	err := loudmouth_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	loudmouth_zip_url := "https://mcabber.com/files/loudmouth/loudmouth-1.5.4.tar.bz2"
	loudmouth_cmd_zip := exec.Command("curl", "-L", loudmouth_zip_url, "-o", "package.zip")
	err = loudmouth_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	loudmouth_bin_url := "https://mcabber.com/files/loudmouth/loudmouth-1.5.4.tar.bz2"
	loudmouth_cmd_bin := exec.Command("curl", "-L", loudmouth_bin_url, "-o", "binary.bin")
	err = loudmouth_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	loudmouth_src_url := "https://mcabber.com/files/loudmouth/loudmouth-1.5.4.tar.bz2"
	loudmouth_cmd_src := exec.Command("curl", "-L", loudmouth_src_url, "-o", "source.tar.gz")
	err = loudmouth_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	loudmouth_cmd_direct := exec.Command("./binary")
	err = loudmouth_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gnutls")
exec.Command("latte", "install", "gnutls")
	fmt.Println("Instalando dependencia: libidn")
exec.Command("latte", "install", "libidn")
}
