package main

// Htmlcompressor - Minify HTML or XML
// Homepage: https://code.google.com/archive/p/htmlcompressor/

import (
	"fmt"
	
	"os/exec"
)

func installHtmlcompressor() {
	// Método 1: Descargar y extraer .tar.gz
	htmlcompressor_tar_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/htmlcompressor/htmlcompressor-1.5.3.jar"
	htmlcompressor_cmd_tar := exec.Command("curl", "-L", htmlcompressor_tar_url, "-o", "package.tar.gz")
	err := htmlcompressor_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	htmlcompressor_zip_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/htmlcompressor/htmlcompressor-1.5.3.jar"
	htmlcompressor_cmd_zip := exec.Command("curl", "-L", htmlcompressor_zip_url, "-o", "package.zip")
	err = htmlcompressor_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	htmlcompressor_bin_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/htmlcompressor/htmlcompressor-1.5.3.jar"
	htmlcompressor_cmd_bin := exec.Command("curl", "-L", htmlcompressor_bin_url, "-o", "binary.bin")
	err = htmlcompressor_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	htmlcompressor_src_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/htmlcompressor/htmlcompressor-1.5.3.jar"
	htmlcompressor_cmd_src := exec.Command("curl", "-L", htmlcompressor_src_url, "-o", "source.tar.gz")
	err = htmlcompressor_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	htmlcompressor_cmd_direct := exec.Command("./binary")
	err = htmlcompressor_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
