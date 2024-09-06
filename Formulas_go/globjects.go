package main

// Globjects - C++ library strictly wrapping OpenGL objects
// Homepage: https://github.com/cginternals/globjects

import (
	"fmt"
	
	"os/exec"
)

func installGlobjects() {
	// Método 1: Descargar y extraer .tar.gz
	globjects_tar_url := "https://github.com/cginternals/globjects/archive/refs/tags/v1.1.0.tar.gz"
	globjects_cmd_tar := exec.Command("curl", "-L", globjects_tar_url, "-o", "package.tar.gz")
	err := globjects_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	globjects_zip_url := "https://github.com/cginternals/globjects/archive/refs/tags/v1.1.0.zip"
	globjects_cmd_zip := exec.Command("curl", "-L", globjects_zip_url, "-o", "package.zip")
	err = globjects_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	globjects_bin_url := "https://github.com/cginternals/globjects/archive/refs/tags/v1.1.0.bin"
	globjects_cmd_bin := exec.Command("curl", "-L", globjects_bin_url, "-o", "binary.bin")
	err = globjects_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	globjects_src_url := "https://github.com/cginternals/globjects/archive/refs/tags/v1.1.0.src.tar.gz"
	globjects_cmd_src := exec.Command("curl", "-L", globjects_src_url, "-o", "source.tar.gz")
	err = globjects_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	globjects_cmd_direct := exec.Command("./binary")
	err = globjects_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: glbinding@2")
exec.Command("latte", "install", "glbinding@2")
	fmt.Println("Instalando dependencia: glbinding")
exec.Command("latte", "install", "glbinding")
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: glm")
exec.Command("latte", "install", "glm")
}
