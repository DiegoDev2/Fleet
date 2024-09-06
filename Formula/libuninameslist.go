package main

// Libuninameslist - Library of Unicode names and annotation data
// Homepage: https://github.com/fontforge/libuninameslist

import (
	"fmt"
	
	"os/exec"
)

func installLibuninameslist() {
	// Método 1: Descargar y extraer .tar.gz
	libuninameslist_tar_url := "https://github.com/fontforge/libuninameslist/releases/download/20240524/libuninameslist-dist-20240524.tar.gz"
	libuninameslist_cmd_tar := exec.Command("curl", "-L", libuninameslist_tar_url, "-o", "package.tar.gz")
	err := libuninameslist_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libuninameslist_zip_url := "https://github.com/fontforge/libuninameslist/releases/download/20240524/libuninameslist-dist-20240524.zip"
	libuninameslist_cmd_zip := exec.Command("curl", "-L", libuninameslist_zip_url, "-o", "package.zip")
	err = libuninameslist_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libuninameslist_bin_url := "https://github.com/fontforge/libuninameslist/releases/download/20240524/libuninameslist-dist-20240524.bin"
	libuninameslist_cmd_bin := exec.Command("curl", "-L", libuninameslist_bin_url, "-o", "binary.bin")
	err = libuninameslist_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libuninameslist_src_url := "https://github.com/fontforge/libuninameslist/releases/download/20240524/libuninameslist-dist-20240524.src.tar.gz"
	libuninameslist_cmd_src := exec.Command("curl", "-L", libuninameslist_src_url, "-o", "source.tar.gz")
	err = libuninameslist_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libuninameslist_cmd_direct := exec.Command("./binary")
	err = libuninameslist_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
}
