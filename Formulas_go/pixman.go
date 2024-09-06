package main

// Pixman - Low-level library for pixel manipulation
// Homepage: https://cairographics.org/

import (
	"fmt"
	
	"os/exec"
)

func installPixman() {
	// Método 1: Descargar y extraer .tar.gz
	pixman_tar_url := "https://cairographics.org/releases/pixman-0.42.2.tar.gz"
	pixman_cmd_tar := exec.Command("curl", "-L", pixman_tar_url, "-o", "package.tar.gz")
	err := pixman_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pixman_zip_url := "https://cairographics.org/releases/pixman-0.42.2.zip"
	pixman_cmd_zip := exec.Command("curl", "-L", pixman_zip_url, "-o", "package.zip")
	err = pixman_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pixman_bin_url := "https://cairographics.org/releases/pixman-0.42.2.bin"
	pixman_cmd_bin := exec.Command("curl", "-L", pixman_bin_url, "-o", "binary.bin")
	err = pixman_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pixman_src_url := "https://cairographics.org/releases/pixman-0.42.2.src.tar.gz"
	pixman_cmd_src := exec.Command("curl", "-L", pixman_src_url, "-o", "source.tar.gz")
	err = pixman_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pixman_cmd_direct := exec.Command("./binary")
	err = pixman_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
