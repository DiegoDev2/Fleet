package main

// Picoc - C interpreter for scripting
// Homepage: https://gitlab.com/zsaleeba/picoc

import (
	"fmt"
	
	"os/exec"
)

func installPicoc() {
	// Método 1: Descargar y extraer .tar.gz
	picoc_tar_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/picoc/picoc-2.1.tar.bz2"
	picoc_cmd_tar := exec.Command("curl", "-L", picoc_tar_url, "-o", "package.tar.gz")
	err := picoc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	picoc_zip_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/picoc/picoc-2.1.tar.bz2"
	picoc_cmd_zip := exec.Command("curl", "-L", picoc_zip_url, "-o", "package.zip")
	err = picoc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	picoc_bin_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/picoc/picoc-2.1.tar.bz2"
	picoc_cmd_bin := exec.Command("curl", "-L", picoc_bin_url, "-o", "binary.bin")
	err = picoc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	picoc_src_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/picoc/picoc-2.1.tar.bz2"
	picoc_cmd_src := exec.Command("curl", "-L", picoc_src_url, "-o", "source.tar.gz")
	err = picoc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	picoc_cmd_direct := exec.Command("./binary")
	err = picoc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
}
