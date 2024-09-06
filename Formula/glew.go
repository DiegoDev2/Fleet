package main

// Glew - OpenGL Extension Wrangler Library
// Homepage: https://glew.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installGlew() {
	// Método 1: Descargar y extraer .tar.gz
	glew_tar_url := "https://downloads.sourceforge.net/project/glew/glew/2.2.0/glew-2.2.0.tgz"
	glew_cmd_tar := exec.Command("curl", "-L", glew_tar_url, "-o", "package.tar.gz")
	err := glew_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	glew_zip_url := "https://downloads.sourceforge.net/project/glew/glew/2.2.0/glew-2.2.0.tgz"
	glew_cmd_zip := exec.Command("curl", "-L", glew_zip_url, "-o", "package.zip")
	err = glew_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	glew_bin_url := "https://downloads.sourceforge.net/project/glew/glew/2.2.0/glew-2.2.0.tgz"
	glew_cmd_bin := exec.Command("curl", "-L", glew_bin_url, "-o", "binary.bin")
	err = glew_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	glew_src_url := "https://downloads.sourceforge.net/project/glew/glew/2.2.0/glew-2.2.0.tgz"
	glew_cmd_src := exec.Command("curl", "-L", glew_src_url, "-o", "source.tar.gz")
	err = glew_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	glew_cmd_direct := exec.Command("./binary")
	err = glew_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: freeglut")
	exec.Command("latte", "install", "freeglut").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: mesa")
	exec.Command("latte", "install", "mesa").Run()
	fmt.Println("Instalando dependencia: mesa-glu")
	exec.Command("latte", "install", "mesa-glu").Run()
}
