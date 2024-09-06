package main

// Guichan - Small, efficient C++ GUI library designed for games
// Homepage: https://guichan.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installGuichan() {
	// Método 1: Descargar y extraer .tar.gz
	guichan_tar_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/guichan/guichan-0.8.2.tar.gz"
	guichan_cmd_tar := exec.Command("curl", "-L", guichan_tar_url, "-o", "package.tar.gz")
	err := guichan_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	guichan_zip_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/guichan/guichan-0.8.2.zip"
	guichan_cmd_zip := exec.Command("curl", "-L", guichan_zip_url, "-o", "package.zip")
	err = guichan_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	guichan_bin_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/guichan/guichan-0.8.2.bin"
	guichan_cmd_bin := exec.Command("curl", "-L", guichan_bin_url, "-o", "binary.bin")
	err = guichan_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	guichan_src_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/guichan/guichan-0.8.2.src.tar.gz"
	guichan_cmd_src := exec.Command("curl", "-L", guichan_src_url, "-o", "source.tar.gz")
	err = guichan_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	guichan_cmd_direct := exec.Command("./binary")
	err = guichan_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: sdl_image")
	exec.Command("latte", "install", "sdl_image").Run()
	fmt.Println("Instalando dependencia: mesa")
	exec.Command("latte", "install", "mesa").Run()
	fmt.Println("Instalando dependencia: mesa-glu")
	exec.Command("latte", "install", "mesa-glu").Run()
}
