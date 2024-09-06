package main

// Glbinding - C++ binding for the OpenGL API
// Homepage: https://github.com/cginternals/glbinding

import (
	"fmt"
	
	"os/exec"
)

func installGlbinding() {
	// Método 1: Descargar y extraer .tar.gz
	glbinding_tar_url := "https://github.com/cginternals/glbinding/archive/refs/tags/v3.3.0.tar.gz"
	glbinding_cmd_tar := exec.Command("curl", "-L", glbinding_tar_url, "-o", "package.tar.gz")
	err := glbinding_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	glbinding_zip_url := "https://github.com/cginternals/glbinding/archive/refs/tags/v3.3.0.zip"
	glbinding_cmd_zip := exec.Command("curl", "-L", glbinding_zip_url, "-o", "package.zip")
	err = glbinding_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	glbinding_bin_url := "https://github.com/cginternals/glbinding/archive/refs/tags/v3.3.0.bin"
	glbinding_cmd_bin := exec.Command("curl", "-L", glbinding_bin_url, "-o", "binary.bin")
	err = glbinding_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	glbinding_src_url := "https://github.com/cginternals/glbinding/archive/refs/tags/v3.3.0.src.tar.gz"
	glbinding_cmd_src := exec.Command("curl", "-L", glbinding_src_url, "-o", "source.tar.gz")
	err = glbinding_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	glbinding_cmd_direct := exec.Command("./binary")
	err = glbinding_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: glfw")
	exec.Command("latte", "install", "glfw").Run()
	fmt.Println("Instalando dependencia: mesa")
	exec.Command("latte", "install", "mesa").Run()
	fmt.Println("Instalando dependencia: mesa-glu")
	exec.Command("latte", "install", "mesa-glu").Run()
}
