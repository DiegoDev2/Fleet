package main

// Pixz - Parallel, indexed, xz compressor
// Homepage: https://github.com/vasi/pixz

import (
	"fmt"
	
	"os/exec"
)

func installPixz() {
	// Método 1: Descargar y extraer .tar.gz
	pixz_tar_url := "https://github.com/vasi/pixz/releases/download/v1.0.7/pixz-1.0.7.tar.gz"
	pixz_cmd_tar := exec.Command("curl", "-L", pixz_tar_url, "-o", "package.tar.gz")
	err := pixz_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pixz_zip_url := "https://github.com/vasi/pixz/releases/download/v1.0.7/pixz-1.0.7.zip"
	pixz_cmd_zip := exec.Command("curl", "-L", pixz_zip_url, "-o", "package.zip")
	err = pixz_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pixz_bin_url := "https://github.com/vasi/pixz/releases/download/v1.0.7/pixz-1.0.7.bin"
	pixz_cmd_bin := exec.Command("curl", "-L", pixz_bin_url, "-o", "binary.bin")
	err = pixz_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pixz_src_url := "https://github.com/vasi/pixz/releases/download/v1.0.7/pixz-1.0.7.src.tar.gz"
	pixz_cmd_src := exec.Command("curl", "-L", pixz_src_url, "-o", "source.tar.gz")
	err = pixz_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pixz_cmd_direct := exec.Command("./binary")
	err = pixz_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: asciidoc")
exec.Command("latte", "install", "asciidoc")
	fmt.Println("Instalando dependencia: docbook-xsl")
exec.Command("latte", "install", "docbook-xsl")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libarchive")
exec.Command("latte", "install", "libarchive")
	fmt.Println("Instalando dependencia: xz")
exec.Command("latte", "install", "xz")
}
