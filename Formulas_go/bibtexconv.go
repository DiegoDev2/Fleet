package main

// Bibtexconv - BibTeX file converter
// Homepage: https://github.com/dreibh/bibtexconv

import (
	"fmt"
	
	"os/exec"
)

func installBibtexconv() {
	// Método 1: Descargar y extraer .tar.gz
	bibtexconv_tar_url := "https://github.com/dreibh/bibtexconv/archive/refs/tags/bibtexconv-1.3.6.tar.gz"
	bibtexconv_cmd_tar := exec.Command("curl", "-L", bibtexconv_tar_url, "-o", "package.tar.gz")
	err := bibtexconv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bibtexconv_zip_url := "https://github.com/dreibh/bibtexconv/archive/refs/tags/bibtexconv-1.3.6.zip"
	bibtexconv_cmd_zip := exec.Command("curl", "-L", bibtexconv_zip_url, "-o", "package.zip")
	err = bibtexconv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bibtexconv_bin_url := "https://github.com/dreibh/bibtexconv/archive/refs/tags/bibtexconv-1.3.6.bin"
	bibtexconv_cmd_bin := exec.Command("curl", "-L", bibtexconv_bin_url, "-o", "binary.bin")
	err = bibtexconv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bibtexconv_src_url := "https://github.com/dreibh/bibtexconv/archive/refs/tags/bibtexconv-1.3.6.src.tar.gz"
	bibtexconv_cmd_src := exec.Command("curl", "-L", bibtexconv_src_url, "-o", "source.tar.gz")
	err = bibtexconv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bibtexconv_cmd_direct := exec.Command("./binary")
	err = bibtexconv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bison")
exec.Command("latte", "install", "bison")
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
