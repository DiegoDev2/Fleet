package main

// Glslang - OpenGL and OpenGL ES reference compiler for shading languages
// Homepage: https://www.khronos.org/opengles/sdk/tools/Reference-Compiler/

import (
	"fmt"
	
	"os/exec"
)

func installGlslang() {
	// Método 1: Descargar y extraer .tar.gz
	glslang_tar_url := "https://github.com/KhronosGroup/glslang/archive/refs/tags/14.3.0.tar.gz"
	glslang_cmd_tar := exec.Command("curl", "-L", glslang_tar_url, "-o", "package.tar.gz")
	err := glslang_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	glslang_zip_url := "https://github.com/KhronosGroup/glslang/archive/refs/tags/14.3.0.zip"
	glslang_cmd_zip := exec.Command("curl", "-L", glslang_zip_url, "-o", "package.zip")
	err = glslang_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	glslang_bin_url := "https://github.com/KhronosGroup/glslang/archive/refs/tags/14.3.0.bin"
	glslang_cmd_bin := exec.Command("curl", "-L", glslang_bin_url, "-o", "binary.bin")
	err = glslang_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	glslang_src_url := "https://github.com/KhronosGroup/glslang/archive/refs/tags/14.3.0.src.tar.gz"
	glslang_cmd_src := exec.Command("curl", "-L", glslang_src_url, "-o", "source.tar.gz")
	err = glslang_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	glslang_cmd_direct := exec.Command("./binary")
	err = glslang_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: spirv-tools")
exec.Command("latte", "install", "spirv-tools")
}
