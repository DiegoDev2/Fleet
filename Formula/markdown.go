package main

// Markdown - Text-to-HTML conversion tool
// Homepage: https://daringfireball.net/projects/markdown/

import (
	"fmt"
	
	"os/exec"
)

func installMarkdown() {
	// Método 1: Descargar y extraer .tar.gz
	markdown_tar_url := "https://daringfireball.net/projects/downloads/Markdown_1.0.1.zip"
	markdown_cmd_tar := exec.Command("curl", "-L", markdown_tar_url, "-o", "package.tar.gz")
	err := markdown_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	markdown_zip_url := "https://daringfireball.net/projects/downloads/Markdown_1.0.1.zip"
	markdown_cmd_zip := exec.Command("curl", "-L", markdown_zip_url, "-o", "package.zip")
	err = markdown_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	markdown_bin_url := "https://daringfireball.net/projects/downloads/Markdown_1.0.1.zip"
	markdown_cmd_bin := exec.Command("curl", "-L", markdown_bin_url, "-o", "binary.bin")
	err = markdown_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	markdown_src_url := "https://daringfireball.net/projects/downloads/Markdown_1.0.1.zip"
	markdown_cmd_src := exec.Command("curl", "-L", markdown_src_url, "-o", "source.tar.gz")
	err = markdown_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	markdown_cmd_direct := exec.Command("./binary")
	err = markdown_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
