package main

// Mmark - Powerful markdown processor in Go geared towards the IETF
// Homepage: https://mmark.miek.nl/

import (
	"fmt"
	
	"os/exec"
)

func installMmark() {
	// Método 1: Descargar y extraer .tar.gz
	mmark_tar_url := "https://github.com/mmarkdown/mmark/archive/refs/tags/v2.2.45.tar.gz"
	mmark_cmd_tar := exec.Command("curl", "-L", mmark_tar_url, "-o", "package.tar.gz")
	err := mmark_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mmark_zip_url := "https://github.com/mmarkdown/mmark/archive/refs/tags/v2.2.45.zip"
	mmark_cmd_zip := exec.Command("curl", "-L", mmark_zip_url, "-o", "package.zip")
	err = mmark_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mmark_bin_url := "https://github.com/mmarkdown/mmark/archive/refs/tags/v2.2.45.bin"
	mmark_cmd_bin := exec.Command("curl", "-L", mmark_bin_url, "-o", "binary.bin")
	err = mmark_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mmark_src_url := "https://github.com/mmarkdown/mmark/archive/refs/tags/v2.2.45.src.tar.gz"
	mmark_cmd_src := exec.Command("curl", "-L", mmark_src_url, "-o", "source.tar.gz")
	err = mmark_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mmark_cmd_direct := exec.Command("./binary")
	err = mmark_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
