package main

// MarkdownToc - Generate a markdown TOC (table of contents) with Remarkable
// Homepage: https://github.com/jonschlinkert/markdown-toc

import (
	"fmt"
	
	"os/exec"
)

func installMarkdownToc() {
	// Método 1: Descargar y extraer .tar.gz
	markdowntoc_tar_url := "https://registry.npmjs.org/markdown-toc/-/markdown-toc-1.2.0.tgz"
	markdowntoc_cmd_tar := exec.Command("curl", "-L", markdowntoc_tar_url, "-o", "package.tar.gz")
	err := markdowntoc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	markdowntoc_zip_url := "https://registry.npmjs.org/markdown-toc/-/markdown-toc-1.2.0.tgz"
	markdowntoc_cmd_zip := exec.Command("curl", "-L", markdowntoc_zip_url, "-o", "package.zip")
	err = markdowntoc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	markdowntoc_bin_url := "https://registry.npmjs.org/markdown-toc/-/markdown-toc-1.2.0.tgz"
	markdowntoc_cmd_bin := exec.Command("curl", "-L", markdowntoc_bin_url, "-o", "binary.bin")
	err = markdowntoc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	markdowntoc_src_url := "https://registry.npmjs.org/markdown-toc/-/markdown-toc-1.2.0.tgz"
	markdowntoc_cmd_src := exec.Command("curl", "-L", markdowntoc_src_url, "-o", "source.tar.gz")
	err = markdowntoc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	markdowntoc_cmd_direct := exec.Command("./binary")
	err = markdowntoc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
}
