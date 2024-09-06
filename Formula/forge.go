package main

// Forge - High Performance Visualization
// Homepage: https://github.com/arrayfire/forge

import (
	"fmt"
	
	"os/exec"
)

func installForge() {
	// Método 1: Descargar y extraer .tar.gz
	forge_tar_url := "https://github.com/arrayfire/forge/archive/refs/tags/v1.0.8.tar.gz"
	forge_cmd_tar := exec.Command("curl", "-L", forge_tar_url, "-o", "package.tar.gz")
	err := forge_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	forge_zip_url := "https://github.com/arrayfire/forge/archive/refs/tags/v1.0.8.zip"
	forge_cmd_zip := exec.Command("curl", "-L", forge_zip_url, "-o", "package.zip")
	err = forge_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	forge_bin_url := "https://github.com/arrayfire/forge/archive/refs/tags/v1.0.8.bin"
	forge_cmd_bin := exec.Command("curl", "-L", forge_bin_url, "-o", "binary.bin")
	err = forge_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	forge_src_url := "https://github.com/arrayfire/forge/archive/refs/tags/v1.0.8.src.tar.gz"
	forge_cmd_src := exec.Command("curl", "-L", forge_src_url, "-o", "source.tar.gz")
	err = forge_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	forge_cmd_direct := exec.Command("./binary")
	err = forge_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: doxygen")
	exec.Command("latte", "install", "doxygen").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: fontconfig")
	exec.Command("latte", "install", "fontconfig").Run()
	fmt.Println("Instalando dependencia: freeimage")
	exec.Command("latte", "install", "freeimage").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: glfw")
	exec.Command("latte", "install", "glfw").Run()
	fmt.Println("Instalando dependencia: glm")
	exec.Command("latte", "install", "glm").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: mesa")
	exec.Command("latte", "install", "mesa").Run()
}
