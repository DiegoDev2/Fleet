package main

// Neomutt - E-mail reader with support for Notmuch, NNTP and much more
// Homepage: https://neomutt.org/

import (
	"fmt"
	
	"os/exec"
)

func installNeomutt() {
	// Método 1: Descargar y extraer .tar.gz
	neomutt_tar_url := "https://github.com/neomutt/neomutt/archive/refs/tags/20240425.tar.gz"
	neomutt_cmd_tar := exec.Command("curl", "-L", neomutt_tar_url, "-o", "package.tar.gz")
	err := neomutt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	neomutt_zip_url := "https://github.com/neomutt/neomutt/archive/refs/tags/20240425.zip"
	neomutt_cmd_zip := exec.Command("curl", "-L", neomutt_zip_url, "-o", "package.zip")
	err = neomutt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	neomutt_bin_url := "https://github.com/neomutt/neomutt/archive/refs/tags/20240425.bin"
	neomutt_cmd_bin := exec.Command("curl", "-L", neomutt_bin_url, "-o", "binary.bin")
	err = neomutt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	neomutt_src_url := "https://github.com/neomutt/neomutt/archive/refs/tags/20240425.src.tar.gz"
	neomutt_cmd_src := exec.Command("curl", "-L", neomutt_src_url, "-o", "source.tar.gz")
	err = neomutt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	neomutt_cmd_direct := exec.Command("./binary")
	err = neomutt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: docbook-xsl")
	exec.Command("latte", "install", "docbook-xsl").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: tcl-tk")
	exec.Command("latte", "install", "tcl-tk").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: gpgme")
	exec.Command("latte", "install", "gpgme").Run()
	fmt.Println("Instalando dependencia: libidn2")
	exec.Command("latte", "install", "libidn2").Run()
	fmt.Println("Instalando dependencia: lmdb")
	exec.Command("latte", "install", "lmdb").Run()
	fmt.Println("Instalando dependencia: lua")
	exec.Command("latte", "install", "lua").Run()
	fmt.Println("Instalando dependencia: ncurses")
	exec.Command("latte", "install", "ncurses").Run()
	fmt.Println("Instalando dependencia: notmuch")
	exec.Command("latte", "install", "notmuch").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: pcre2")
	exec.Command("latte", "install", "pcre2").Run()
	fmt.Println("Instalando dependencia: sqlite")
	exec.Command("latte", "install", "sqlite").Run()
	fmt.Println("Instalando dependencia: tokyo-cabinet")
	exec.Command("latte", "install", "tokyo-cabinet").Run()
	fmt.Println("Instalando dependencia: libgpg-error")
	exec.Command("latte", "install", "libgpg-error").Run()
	fmt.Println("Instalando dependencia: libiconv")
	exec.Command("latte", "install", "libiconv").Run()
}
