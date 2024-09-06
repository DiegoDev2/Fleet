package main

// Cmark - Strongly specified, highly compatible implementation of Markdown
// Homepage: https://commonmark.org/

import (
	"fmt"
	
	"os/exec"
)

func installCmark() {
	// Método 1: Descargar y extraer .tar.gz
	cmark_tar_url := "https://github.com/commonmark/cmark/archive/refs/tags/0.31.1.tar.gz"
	cmark_cmd_tar := exec.Command("curl", "-L", cmark_tar_url, "-o", "package.tar.gz")
	err := cmark_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cmark_zip_url := "https://github.com/commonmark/cmark/archive/refs/tags/0.31.1.zip"
	cmark_cmd_zip := exec.Command("curl", "-L", cmark_zip_url, "-o", "package.zip")
	err = cmark_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cmark_bin_url := "https://github.com/commonmark/cmark/archive/refs/tags/0.31.1.bin"
	cmark_cmd_bin := exec.Command("curl", "-L", cmark_bin_url, "-o", "binary.bin")
	err = cmark_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cmark_src_url := "https://github.com/commonmark/cmark/archive/refs/tags/0.31.1.src.tar.gz"
	cmark_cmd_src := exec.Command("curl", "-L", cmark_src_url, "-o", "source.tar.gz")
	err = cmark_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cmark_cmd_direct := exec.Command("./binary")
	err = cmark_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
