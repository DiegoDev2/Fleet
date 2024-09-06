package main

// Emacs - GNU Emacs text editor
// Homepage: https://www.gnu.org/software/emacs/

import (
	"fmt"
	
	"os/exec"
)

func installEmacs() {
	// Método 1: Descargar y extraer .tar.gz
	emacs_tar_url := "https://ftp.gnu.org/gnu/emacs/emacs-29.4.tar.xz"
	emacs_cmd_tar := exec.Command("curl", "-L", emacs_tar_url, "-o", "package.tar.gz")
	err := emacs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	emacs_zip_url := "https://ftp.gnu.org/gnu/emacs/emacs-29.4.tar.xz"
	emacs_cmd_zip := exec.Command("curl", "-L", emacs_zip_url, "-o", "package.zip")
	err = emacs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	emacs_bin_url := "https://ftp.gnu.org/gnu/emacs/emacs-29.4.tar.xz"
	emacs_cmd_bin := exec.Command("curl", "-L", emacs_bin_url, "-o", "binary.bin")
	err = emacs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	emacs_src_url := "https://ftp.gnu.org/gnu/emacs/emacs-29.4.tar.xz"
	emacs_cmd_src := exec.Command("curl", "-L", emacs_src_url, "-o", "source.tar.gz")
	err = emacs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	emacs_cmd_direct := exec.Command("./binary")
	err = emacs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: gnu-sed")
exec.Command("latte", "install", "gnu-sed")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: texinfo")
exec.Command("latte", "install", "texinfo")
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
	fmt.Println("Instalando dependencia: gnutls")
exec.Command("latte", "install", "gnutls")
	fmt.Println("Instalando dependencia: jansson")
exec.Command("latte", "install", "jansson")
	fmt.Println("Instalando dependencia: tree-sitter")
exec.Command("latte", "install", "tree-sitter")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
}
