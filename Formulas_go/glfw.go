package main

// Glfw - Multi-platform library for OpenGL applications
// Homepage: https://www.glfw.org/

import (
	"fmt"
	
	"os/exec"
)

func installGlfw() {
	// Método 1: Descargar y extraer .tar.gz
	glfw_tar_url := "https://github.com/glfw/glfw/archive/refs/tags/3.4.tar.gz"
	glfw_cmd_tar := exec.Command("curl", "-L", glfw_tar_url, "-o", "package.tar.gz")
	err := glfw_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	glfw_zip_url := "https://github.com/glfw/glfw/archive/refs/tags/3.4.zip"
	glfw_cmd_zip := exec.Command("curl", "-L", glfw_zip_url, "-o", "package.zip")
	err = glfw_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	glfw_bin_url := "https://github.com/glfw/glfw/archive/refs/tags/3.4.bin"
	glfw_cmd_bin := exec.Command("curl", "-L", glfw_bin_url, "-o", "binary.bin")
	err = glfw_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	glfw_src_url := "https://github.com/glfw/glfw/archive/refs/tags/3.4.src.tar.gz"
	glfw_cmd_src := exec.Command("curl", "-L", glfw_src_url, "-o", "source.tar.gz")
	err = glfw_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	glfw_cmd_direct := exec.Command("./binary")
	err = glfw_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: freeglut")
exec.Command("latte", "install", "freeglut")
	fmt.Println("Instalando dependencia: libxcursor")
exec.Command("latte", "install", "libxcursor")
	fmt.Println("Instalando dependencia: libxkbcommon")
exec.Command("latte", "install", "libxkbcommon")
	fmt.Println("Instalando dependencia: mesa")
exec.Command("latte", "install", "mesa")
}
