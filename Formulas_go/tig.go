package main

// Tig - Text interface for Git repositories
// Homepage: https://jonas.github.io/tig/

import (
	"fmt"
	
	"os/exec"
)

func installTig() {
	// Método 1: Descargar y extraer .tar.gz
	tig_tar_url := "https://github.com/jonas/tig/releases/download/tig-2.5.10/tig-2.5.10.tar.gz"
	tig_cmd_tar := exec.Command("curl", "-L", tig_tar_url, "-o", "package.tar.gz")
	err := tig_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tig_zip_url := "https://github.com/jonas/tig/releases/download/tig-2.5.10/tig-2.5.10.zip"
	tig_cmd_zip := exec.Command("curl", "-L", tig_zip_url, "-o", "package.zip")
	err = tig_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tig_bin_url := "https://github.com/jonas/tig/releases/download/tig-2.5.10/tig-2.5.10.bin"
	tig_cmd_bin := exec.Command("curl", "-L", tig_bin_url, "-o", "binary.bin")
	err = tig_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tig_src_url := "https://github.com/jonas/tig/releases/download/tig-2.5.10/tig-2.5.10.src.tar.gz"
	tig_cmd_src := exec.Command("curl", "-L", tig_src_url, "-o", "source.tar.gz")
	err = tig_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tig_cmd_direct := exec.Command("./binary")
	err = tig_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: asciidoc")
exec.Command("latte", "install", "asciidoc")
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: xmlto")
exec.Command("latte", "install", "xmlto")
	fmt.Println("Instalando dependencia: ncurses")
exec.Command("latte", "install", "ncurses")
	fmt.Println("Instalando dependencia: pcre2")
exec.Command("latte", "install", "pcre2")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
