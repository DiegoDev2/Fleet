package main

// Mutt - Mongrel of mail user agents (part elm, pine, mush, mh, etc.)
// Homepage: http://www.mutt.org/

import (
	"fmt"
	
	"os/exec"
)

func installMutt() {
	// Método 1: Descargar y extraer .tar.gz
	mutt_tar_url := "https://bitbucket.org/mutt/mutt/downloads/mutt-2.2.13.tar.gz"
	mutt_cmd_tar := exec.Command("curl", "-L", mutt_tar_url, "-o", "package.tar.gz")
	err := mutt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mutt_zip_url := "https://bitbucket.org/mutt/mutt/downloads/mutt-2.2.13.zip"
	mutt_cmd_zip := exec.Command("curl", "-L", mutt_zip_url, "-o", "package.zip")
	err = mutt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mutt_bin_url := "https://bitbucket.org/mutt/mutt/downloads/mutt-2.2.13.bin"
	mutt_cmd_bin := exec.Command("curl", "-L", mutt_bin_url, "-o", "binary.bin")
	err = mutt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mutt_src_url := "https://bitbucket.org/mutt/mutt/downloads/mutt-2.2.13.src.tar.gz"
	mutt_cmd_src := exec.Command("curl", "-L", mutt_src_url, "-o", "source.tar.gz")
	err = mutt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mutt_cmd_direct := exec.Command("./binary")
	err = mutt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: gpgme")
exec.Command("latte", "install", "gpgme")
	fmt.Println("Instalando dependencia: libgpg-error")
exec.Command("latte", "install", "libgpg-error")
	fmt.Println("Instalando dependencia: libidn2")
exec.Command("latte", "install", "libidn2")
	fmt.Println("Instalando dependencia: ncurses")
exec.Command("latte", "install", "ncurses")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: tokyo-cabinet")
exec.Command("latte", "install", "tokyo-cabinet")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: libunistring")
exec.Command("latte", "install", "libunistring")
}
