package main

// Glui - C++ user interface library
// Homepage: https://github.com/libglui/glui

import (
	"fmt"
	
	"os/exec"
)

func installGlui() {
	// Método 1: Descargar y extraer .tar.gz
	glui_tar_url := "https://github.com/libglui/glui/archive/refs/tags/2.37.tar.gz"
	glui_cmd_tar := exec.Command("curl", "-L", glui_tar_url, "-o", "package.tar.gz")
	err := glui_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	glui_zip_url := "https://github.com/libglui/glui/archive/refs/tags/2.37.zip"
	glui_cmd_zip := exec.Command("curl", "-L", glui_zip_url, "-o", "package.zip")
	err = glui_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	glui_bin_url := "https://github.com/libglui/glui/archive/refs/tags/2.37.bin"
	glui_cmd_bin := exec.Command("curl", "-L", glui_bin_url, "-o", "binary.bin")
	err = glui_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	glui_src_url := "https://github.com/libglui/glui/archive/refs/tags/2.37.src.tar.gz"
	glui_cmd_src := exec.Command("curl", "-L", glui_src_url, "-o", "source.tar.gz")
	err = glui_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	glui_cmd_direct := exec.Command("./binary")
	err = glui_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: freeglut")
exec.Command("latte", "install", "freeglut")
	fmt.Println("Instalando dependencia: mesa")
exec.Command("latte", "install", "mesa")
	fmt.Println("Instalando dependencia: mesa-glu")
exec.Command("latte", "install", "mesa-glu")
}
