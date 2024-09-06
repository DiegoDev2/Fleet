package main

// PegMarkdown - Markdown implementation based on a PEG grammar
// Homepage: https://github.com/jgm/peg-markdown

import (
	"fmt"
	
	"os/exec"
)

func installPegMarkdown() {
	// Método 1: Descargar y extraer .tar.gz
	pegmarkdown_tar_url := "https://github.com/jgm/peg-markdown/archive/refs/tags/0.4.14.tar.gz"
	pegmarkdown_cmd_tar := exec.Command("curl", "-L", pegmarkdown_tar_url, "-o", "package.tar.gz")
	err := pegmarkdown_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pegmarkdown_zip_url := "https://github.com/jgm/peg-markdown/archive/refs/tags/0.4.14.zip"
	pegmarkdown_cmd_zip := exec.Command("curl", "-L", pegmarkdown_zip_url, "-o", "package.zip")
	err = pegmarkdown_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pegmarkdown_bin_url := "https://github.com/jgm/peg-markdown/archive/refs/tags/0.4.14.bin"
	pegmarkdown_cmd_bin := exec.Command("curl", "-L", pegmarkdown_bin_url, "-o", "binary.bin")
	err = pegmarkdown_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pegmarkdown_src_url := "https://github.com/jgm/peg-markdown/archive/refs/tags/0.4.14.src.tar.gz"
	pegmarkdown_cmd_src := exec.Command("curl", "-L", pegmarkdown_src_url, "-o", "source.tar.gz")
	err = pegmarkdown_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pegmarkdown_cmd_direct := exec.Command("./binary")
	err = pegmarkdown_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
}
