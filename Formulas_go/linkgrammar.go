package main

// LinkGrammar - Carnegie Mellon University's link grammar parser
// Homepage: https://github.com/opencog/link-grammar

import (
	"fmt"
	
	"os/exec"
)

func installLinkGrammar() {
	// Método 1: Descargar y extraer .tar.gz
	linkgrammar_tar_url := "https://github.com/opencog/link-grammar/archive/refs/tags/link-grammar-5.12.5.tar.gz"
	linkgrammar_cmd_tar := exec.Command("curl", "-L", linkgrammar_tar_url, "-o", "package.tar.gz")
	err := linkgrammar_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	linkgrammar_zip_url := "https://github.com/opencog/link-grammar/archive/refs/tags/link-grammar-5.12.5.zip"
	linkgrammar_cmd_zip := exec.Command("curl", "-L", linkgrammar_zip_url, "-o", "package.zip")
	err = linkgrammar_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	linkgrammar_bin_url := "https://github.com/opencog/link-grammar/archive/refs/tags/link-grammar-5.12.5.bin"
	linkgrammar_cmd_bin := exec.Command("curl", "-L", linkgrammar_bin_url, "-o", "binary.bin")
	err = linkgrammar_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	linkgrammar_src_url := "https://github.com/opencog/link-grammar/archive/refs/tags/link-grammar-5.12.5.src.tar.gz"
	linkgrammar_cmd_src := exec.Command("curl", "-L", linkgrammar_src_url, "-o", "source.tar.gz")
	err = linkgrammar_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	linkgrammar_cmd_direct := exec.Command("./binary")
	err = linkgrammar_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ant")
exec.Command("latte", "install", "ant")
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: autoconf-archive")
exec.Command("latte", "install", "autoconf-archive")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: swig")
exec.Command("latte", "install", "swig")
}
