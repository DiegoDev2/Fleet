package main

// GlbindingAT2 - C++ binding for the OpenGL API
// Homepage: https://github.com/cginternals/glbinding

import (
	"fmt"
	
	"os/exec"
)

func installGlbindingAT2() {
	// Método 1: Descargar y extraer .tar.gz
	glbindingat2_tar_url := "https://github.com/cginternals/glbinding/archive/refs/tags/v2.1.4.tar.gz"
	glbindingat2_cmd_tar := exec.Command("curl", "-L", glbindingat2_tar_url, "-o", "package.tar.gz")
	err := glbindingat2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	glbindingat2_zip_url := "https://github.com/cginternals/glbinding/archive/refs/tags/v2.1.4.zip"
	glbindingat2_cmd_zip := exec.Command("curl", "-L", glbindingat2_zip_url, "-o", "package.zip")
	err = glbindingat2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	glbindingat2_bin_url := "https://github.com/cginternals/glbinding/archive/refs/tags/v2.1.4.bin"
	glbindingat2_cmd_bin := exec.Command("curl", "-L", glbindingat2_bin_url, "-o", "binary.bin")
	err = glbindingat2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	glbindingat2_src_url := "https://github.com/cginternals/glbinding/archive/refs/tags/v2.1.4.src.tar.gz"
	glbindingat2_cmd_src := exec.Command("curl", "-L", glbindingat2_src_url, "-o", "source.tar.gz")
	err = glbindingat2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	glbindingat2_cmd_direct := exec.Command("./binary")
	err = glbindingat2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: mesa")
	exec.Command("latte", "install", "mesa").Run()
	fmt.Println("Instalando dependencia: mesa-glu")
	exec.Command("latte", "install", "mesa-glu").Run()
}
