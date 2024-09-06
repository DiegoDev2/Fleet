package main

// Gmime - MIME mail utilities
// Homepage: https://github.com/jstedfast/gmime

import (
	"fmt"
	
	"os/exec"
)

func installGmime() {
	// Método 1: Descargar y extraer .tar.gz
	gmime_tar_url := "https://github.com/jstedfast/gmime/releases/download/3.2.15/gmime-3.2.15.tar.xz"
	gmime_cmd_tar := exec.Command("curl", "-L", gmime_tar_url, "-o", "package.tar.gz")
	err := gmime_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gmime_zip_url := "https://github.com/jstedfast/gmime/releases/download/3.2.15/gmime-3.2.15.tar.xz"
	gmime_cmd_zip := exec.Command("curl", "-L", gmime_zip_url, "-o", "package.zip")
	err = gmime_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gmime_bin_url := "https://github.com/jstedfast/gmime/releases/download/3.2.15/gmime-3.2.15.tar.xz"
	gmime_cmd_bin := exec.Command("curl", "-L", gmime_bin_url, "-o", "binary.bin")
	err = gmime_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gmime_src_url := "https://github.com/jstedfast/gmime/releases/download/3.2.15/gmime-3.2.15.tar.xz"
	gmime_cmd_src := exec.Command("curl", "-L", gmime_src_url, "-o", "source.tar.gz")
	err = gmime_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gmime_cmd_direct := exec.Command("./binary")
	err = gmime_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: gtk-doc")
	exec.Command("latte", "install", "gtk-doc").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: gobject-introspection")
	exec.Command("latte", "install", "gobject-introspection").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: gpgme")
	exec.Command("latte", "install", "gpgme").Run()
	fmt.Println("Instalando dependencia: libidn2")
	exec.Command("latte", "install", "libidn2").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: libgpg-error")
	exec.Command("latte", "install", "libgpg-error").Run()
}
