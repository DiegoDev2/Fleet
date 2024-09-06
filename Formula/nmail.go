package main

// Nmail - Terminal-based email client for Linux and macOS
// Homepage: https://github.com/d99kris/nmail

import (
	"fmt"
	
	"os/exec"
)

func installNmail() {
	// Método 1: Descargar y extraer .tar.gz
	nmail_tar_url := "https://github.com/d99kris/nmail/archive/refs/tags/v4.67.tar.gz"
	nmail_cmd_tar := exec.Command("curl", "-L", nmail_tar_url, "-o", "package.tar.gz")
	err := nmail_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nmail_zip_url := "https://github.com/d99kris/nmail/archive/refs/tags/v4.67.zip"
	nmail_cmd_zip := exec.Command("curl", "-L", nmail_zip_url, "-o", "package.zip")
	err = nmail_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nmail_bin_url := "https://github.com/d99kris/nmail/archive/refs/tags/v4.67.bin"
	nmail_cmd_bin := exec.Command("curl", "-L", nmail_bin_url, "-o", "binary.bin")
	err = nmail_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nmail_src_url := "https://github.com/d99kris/nmail/archive/refs/tags/v4.67.src.tar.gz"
	nmail_cmd_src := exec.Command("curl", "-L", nmail_src_url, "-o", "source.tar.gz")
	err = nmail_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nmail_cmd_direct := exec.Command("./binary")
	err = nmail_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: libmagic")
	exec.Command("latte", "install", "libmagic").Run()
	fmt.Println("Instalando dependencia: ncurses")
	exec.Command("latte", "install", "ncurses").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: util-linux")
	exec.Command("latte", "install", "util-linux").Run()
	fmt.Println("Instalando dependencia: xapian")
	exec.Command("latte", "install", "xapian").Run()
}
