package main

// Textidote - Spelling, grammar and style checking on LaTeX documents
// Homepage: https://sylvainhalle.github.io/textidote

import (
	"fmt"
	
	"os/exec"
)

func installTextidote() {
	// Método 1: Descargar y extraer .tar.gz
	textidote_tar_url := "https://github.com/sylvainhalle/textidote/archive/refs/tags/v0.8.3.tar.gz"
	textidote_cmd_tar := exec.Command("curl", "-L", textidote_tar_url, "-o", "package.tar.gz")
	err := textidote_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	textidote_zip_url := "https://github.com/sylvainhalle/textidote/archive/refs/tags/v0.8.3.zip"
	textidote_cmd_zip := exec.Command("curl", "-L", textidote_zip_url, "-o", "package.zip")
	err = textidote_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	textidote_bin_url := "https://github.com/sylvainhalle/textidote/archive/refs/tags/v0.8.3.bin"
	textidote_cmd_bin := exec.Command("curl", "-L", textidote_bin_url, "-o", "binary.bin")
	err = textidote_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	textidote_src_url := "https://github.com/sylvainhalle/textidote/archive/refs/tags/v0.8.3.src.tar.gz"
	textidote_cmd_src := exec.Command("curl", "-L", textidote_src_url, "-o", "source.tar.gz")
	err = textidote_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	textidote_cmd_direct := exec.Command("./binary")
	err = textidote_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ant")
exec.Command("latte", "install", "ant")
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
