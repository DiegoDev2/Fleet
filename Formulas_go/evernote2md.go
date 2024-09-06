package main

// Evernote2md - Convert Evernote .enex file to Markdown
// Homepage: https://github.com/wormi4ok/evernote2md

import (
	"fmt"
	
	"os/exec"
)

func installEvernote2md() {
	// Método 1: Descargar y extraer .tar.gz
	evernote2md_tar_url := "https://github.com/wormi4ok/evernote2md/archive/refs/tags/v0.21.0.tar.gz"
	evernote2md_cmd_tar := exec.Command("curl", "-L", evernote2md_tar_url, "-o", "package.tar.gz")
	err := evernote2md_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	evernote2md_zip_url := "https://github.com/wormi4ok/evernote2md/archive/refs/tags/v0.21.0.zip"
	evernote2md_cmd_zip := exec.Command("curl", "-L", evernote2md_zip_url, "-o", "package.zip")
	err = evernote2md_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	evernote2md_bin_url := "https://github.com/wormi4ok/evernote2md/archive/refs/tags/v0.21.0.bin"
	evernote2md_cmd_bin := exec.Command("curl", "-L", evernote2md_bin_url, "-o", "binary.bin")
	err = evernote2md_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	evernote2md_src_url := "https://github.com/wormi4ok/evernote2md/archive/refs/tags/v0.21.0.src.tar.gz"
	evernote2md_cmd_src := exec.Command("curl", "-L", evernote2md_src_url, "-o", "source.tar.gz")
	err = evernote2md_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	evernote2md_cmd_direct := exec.Command("./binary")
	err = evernote2md_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
