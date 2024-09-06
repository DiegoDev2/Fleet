package main

// Pdfgrep - Search PDFs for strings matching a regular expression
// Homepage: https://pdfgrep.org/

import (
	"fmt"
	
	"os/exec"
)

func installPdfgrep() {
	// Método 1: Descargar y extraer .tar.gz
	pdfgrep_tar_url := "https://pdfgrep.org/download/pdfgrep-2.2.0.tar.gz"
	pdfgrep_cmd_tar := exec.Command("curl", "-L", pdfgrep_tar_url, "-o", "package.tar.gz")
	err := pdfgrep_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pdfgrep_zip_url := "https://pdfgrep.org/download/pdfgrep-2.2.0.zip"
	pdfgrep_cmd_zip := exec.Command("curl", "-L", pdfgrep_zip_url, "-o", "package.zip")
	err = pdfgrep_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pdfgrep_bin_url := "https://pdfgrep.org/download/pdfgrep-2.2.0.bin"
	pdfgrep_cmd_bin := exec.Command("curl", "-L", pdfgrep_bin_url, "-o", "binary.bin")
	err = pdfgrep_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pdfgrep_src_url := "https://pdfgrep.org/download/pdfgrep-2.2.0.src.tar.gz"
	pdfgrep_cmd_src := exec.Command("curl", "-L", pdfgrep_src_url, "-o", "source.tar.gz")
	err = pdfgrep_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pdfgrep_cmd_direct := exec.Command("./binary")
	err = pdfgrep_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libgcrypt")
exec.Command("latte", "install", "libgcrypt")
	fmt.Println("Instalando dependencia: pcre2")
exec.Command("latte", "install", "pcre2")
	fmt.Println("Instalando dependencia: poppler")
exec.Command("latte", "install", "poppler")
	fmt.Println("Instalando dependencia: libgpg-error")
exec.Command("latte", "install", "libgpg-error")
}
