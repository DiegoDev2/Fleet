package main

// ScIm - Spreadsheet program for the terminal, using ncurses
// Homepage: https://github.com/andmarti1424/sc-im

import (
	"fmt"
	
	"os/exec"
)

func installScIm() {
	// Método 1: Descargar y extraer .tar.gz
	scim_tar_url := "https://github.com/andmarti1424/sc-im/archive/refs/tags/v0.8.4.tar.gz"
	scim_cmd_tar := exec.Command("curl", "-L", scim_tar_url, "-o", "package.tar.gz")
	err := scim_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	scim_zip_url := "https://github.com/andmarti1424/sc-im/archive/refs/tags/v0.8.4.zip"
	scim_cmd_zip := exec.Command("curl", "-L", scim_zip_url, "-o", "package.zip")
	err = scim_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	scim_bin_url := "https://github.com/andmarti1424/sc-im/archive/refs/tags/v0.8.4.bin"
	scim_cmd_bin := exec.Command("curl", "-L", scim_bin_url, "-o", "binary.bin")
	err = scim_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	scim_src_url := "https://github.com/andmarti1424/sc-im/archive/refs/tags/v0.8.4.src.tar.gz"
	scim_cmd_src := exec.Command("curl", "-L", scim_src_url, "-o", "source.tar.gz")
	err = scim_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	scim_cmd_direct := exec.Command("./binary")
	err = scim_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libxls")
	exec.Command("latte", "install", "libxls").Run()
	fmt.Println("Instalando dependencia: libxlsxwriter")
	exec.Command("latte", "install", "libxlsxwriter").Run()
	fmt.Println("Instalando dependencia: libxml2")
	exec.Command("latte", "install", "libxml2").Run()
	fmt.Println("Instalando dependencia: libzip")
	exec.Command("latte", "install", "libzip").Run()
	fmt.Println("Instalando dependencia: lua")
	exec.Command("latte", "install", "lua").Run()
	fmt.Println("Instalando dependencia: ncurses")
	exec.Command("latte", "install", "ncurses").Run()
}
