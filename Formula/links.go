package main

// Links - Lynx-like WWW browser that supports tables, menus, etc.
// Homepage: http://links.twibright.com/

import (
	"fmt"
	
	"os/exec"
)

func installLinks() {
	// Método 1: Descargar y extraer .tar.gz
	links_tar_url := "http://links.twibright.com/download/links-2.30.tar.bz2"
	links_cmd_tar := exec.Command("curl", "-L", links_tar_url, "-o", "package.tar.gz")
	err := links_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	links_zip_url := "http://links.twibright.com/download/links-2.30.tar.bz2"
	links_cmd_zip := exec.Command("curl", "-L", links_zip_url, "-o", "package.zip")
	err = links_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	links_bin_url := "http://links.twibright.com/download/links-2.30.tar.bz2"
	links_cmd_bin := exec.Command("curl", "-L", links_bin_url, "-o", "binary.bin")
	err = links_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	links_src_url := "http://links.twibright.com/download/links-2.30.tar.bz2"
	links_cmd_src := exec.Command("curl", "-L", links_src_url, "-o", "source.tar.gz")
	err = links_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	links_cmd_direct := exec.Command("./binary")
	err = links_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
