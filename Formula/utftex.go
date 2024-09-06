package main

// Utftex - Pretty print math in monospace fonts, using a TeX-like syntax
// Homepage: https://github.com/bartp5/libtexprintf

import (
	"fmt"
	
	"os/exec"
)

func installUtftex() {
	// Método 1: Descargar y extraer .tar.gz
	utftex_tar_url := "https://github.com/bartp5/libtexprintf/archive/refs/tags/v1.25.tar.gz"
	utftex_cmd_tar := exec.Command("curl", "-L", utftex_tar_url, "-o", "package.tar.gz")
	err := utftex_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	utftex_zip_url := "https://github.com/bartp5/libtexprintf/archive/refs/tags/v1.25.zip"
	utftex_cmd_zip := exec.Command("curl", "-L", utftex_zip_url, "-o", "package.zip")
	err = utftex_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	utftex_bin_url := "https://github.com/bartp5/libtexprintf/archive/refs/tags/v1.25.bin"
	utftex_cmd_bin := exec.Command("curl", "-L", utftex_bin_url, "-o", "binary.bin")
	err = utftex_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	utftex_src_url := "https://github.com/bartp5/libtexprintf/archive/refs/tags/v1.25.src.tar.gz"
	utftex_cmd_src := exec.Command("curl", "-L", utftex_src_url, "-o", "source.tar.gz")
	err = utftex_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	utftex_cmd_direct := exec.Command("./binary")
	err = utftex_cmd_direct.Run()
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
