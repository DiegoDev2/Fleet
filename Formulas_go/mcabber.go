package main

// Mcabber - Console Jabber client
// Homepage: https://mcabber.com/

import (
	"fmt"
	
	"os/exec"
)

func installMcabber() {
	// Método 1: Descargar y extraer .tar.gz
	mcabber_tar_url := "https://mcabber.com/files/mcabber-1.1.2.tar.bz2"
	mcabber_cmd_tar := exec.Command("curl", "-L", mcabber_tar_url, "-o", "package.tar.gz")
	err := mcabber_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mcabber_zip_url := "https://mcabber.com/files/mcabber-1.1.2.tar.bz2"
	mcabber_cmd_zip := exec.Command("curl", "-L", mcabber_zip_url, "-o", "package.zip")
	err = mcabber_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mcabber_bin_url := "https://mcabber.com/files/mcabber-1.1.2.tar.bz2"
	mcabber_cmd_bin := exec.Command("curl", "-L", mcabber_bin_url, "-o", "binary.bin")
	err = mcabber_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mcabber_src_url := "https://mcabber.com/files/mcabber-1.1.2.tar.bz2"
	mcabber_cmd_src := exec.Command("curl", "-L", mcabber_src_url, "-o", "source.tar.gz")
	err = mcabber_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mcabber_cmd_direct := exec.Command("./binary")
	err = mcabber_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gpgme")
exec.Command("latte", "install", "gpgme")
	fmt.Println("Instalando dependencia: libgcrypt")
exec.Command("latte", "install", "libgcrypt")
	fmt.Println("Instalando dependencia: libidn")
exec.Command("latte", "install", "libidn")
	fmt.Println("Instalando dependencia: libotr")
exec.Command("latte", "install", "libotr")
	fmt.Println("Instalando dependencia: loudmouth")
exec.Command("latte", "install", "loudmouth")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: libassuan")
exec.Command("latte", "install", "libassuan")
	fmt.Println("Instalando dependencia: libgpg-error")
exec.Command("latte", "install", "libgpg-error")
}
