package main

// GerbilScheme - Opinionated dialect of Scheme designed for Systems Programming
// Homepage: https://cons.io

import (
	"fmt"
	
	"os/exec"
)

func installGerbilScheme() {
	// Método 1: Descargar y extraer .tar.gz
	gerbilscheme_tar_url := "https://github.com/vyzo/gerbil/archive/refs/tags/v0.17.tar.gz"
	gerbilscheme_cmd_tar := exec.Command("curl", "-L", gerbilscheme_tar_url, "-o", "package.tar.gz")
	err := gerbilscheme_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gerbilscheme_zip_url := "https://github.com/vyzo/gerbil/archive/refs/tags/v0.17.zip"
	gerbilscheme_cmd_zip := exec.Command("curl", "-L", gerbilscheme_zip_url, "-o", "package.zip")
	err = gerbilscheme_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gerbilscheme_bin_url := "https://github.com/vyzo/gerbil/archive/refs/tags/v0.17.bin"
	gerbilscheme_cmd_bin := exec.Command("curl", "-L", gerbilscheme_bin_url, "-o", "binary.bin")
	err = gerbilscheme_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gerbilscheme_src_url := "https://github.com/vyzo/gerbil/archive/refs/tags/v0.17.src.tar.gz"
	gerbilscheme_cmd_src := exec.Command("curl", "-L", gerbilscheme_src_url, "-o", "source.tar.gz")
	err = gerbilscheme_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gerbilscheme_cmd_direct := exec.Command("./binary")
	err = gerbilscheme_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gambit-scheme")
exec.Command("latte", "install", "gambit-scheme")
	fmt.Println("Instalando dependencia: leveldb")
exec.Command("latte", "install", "leveldb")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: lmdb")
exec.Command("latte", "install", "lmdb")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: gcc")
exec.Command("latte", "install", "gcc")
}
