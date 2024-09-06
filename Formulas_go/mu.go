package main

// Mu - Tool for searching e-mail messages stored in the maildir-format
// Homepage: https://www.djcbsoftware.nl/code/mu/

import (
	"fmt"
	
	"os/exec"
)

func installMu() {
	// Método 1: Descargar y extraer .tar.gz
	mu_tar_url := "https://github.com/djcb/mu/releases/download/v1.12.6/mu-1.12.6.tar.xz"
	mu_cmd_tar := exec.Command("curl", "-L", mu_tar_url, "-o", "package.tar.gz")
	err := mu_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mu_zip_url := "https://github.com/djcb/mu/releases/download/v1.12.6/mu-1.12.6.tar.xz"
	mu_cmd_zip := exec.Command("curl", "-L", mu_zip_url, "-o", "package.zip")
	err = mu_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mu_bin_url := "https://github.com/djcb/mu/releases/download/v1.12.6/mu-1.12.6.tar.xz"
	mu_cmd_bin := exec.Command("curl", "-L", mu_bin_url, "-o", "binary.bin")
	err = mu_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mu_src_url := "https://github.com/djcb/mu/releases/download/v1.12.6/mu-1.12.6.tar.xz"
	mu_cmd_src := exec.Command("curl", "-L", mu_src_url, "-o", "source.tar.gz")
	err = mu_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mu_cmd_direct := exec.Command("./binary")
	err = mu_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: emacs")
exec.Command("latte", "install", "emacs")
	fmt.Println("Instalando dependencia: libgpg-error")
exec.Command("latte", "install", "libgpg-error")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: texinfo")
exec.Command("latte", "install", "texinfo")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gmime")
exec.Command("latte", "install", "gmime")
	fmt.Println("Instalando dependencia: xapian")
exec.Command("latte", "install", "xapian")
}
