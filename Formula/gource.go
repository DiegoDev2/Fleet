package main

// Gource - Version Control Visualization Tool
// Homepage: https://github.com/acaudwell/Gource

import (
	"fmt"
	
	"os/exec"
)

func installGource() {
	// Método 1: Descargar y extraer .tar.gz
	gource_tar_url := "https://github.com/acaudwell/Gource/releases/download/gource-0.55/gource-0.55.tar.gz"
	gource_cmd_tar := exec.Command("curl", "-L", gource_tar_url, "-o", "package.tar.gz")
	err := gource_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gource_zip_url := "https://github.com/acaudwell/Gource/releases/download/gource-0.55/gource-0.55.zip"
	gource_cmd_zip := exec.Command("curl", "-L", gource_zip_url, "-o", "package.zip")
	err = gource_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gource_bin_url := "https://github.com/acaudwell/Gource/releases/download/gource-0.55/gource-0.55.bin"
	gource_cmd_bin := exec.Command("curl", "-L", gource_bin_url, "-o", "binary.bin")
	err = gource_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gource_src_url := "https://github.com/acaudwell/Gource/releases/download/gource-0.55/gource-0.55.src.tar.gz"
	gource_cmd_src := exec.Command("curl", "-L", gource_src_url, "-o", "source.tar.gz")
	err = gource_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gource_cmd_direct := exec.Command("./binary")
	err = gource_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: glm")
	exec.Command("latte", "install", "glm").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: glew")
	exec.Command("latte", "install", "glew").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: pcre2")
	exec.Command("latte", "install", "pcre2").Run()
	fmt.Println("Instalando dependencia: sdl2")
	exec.Command("latte", "install", "sdl2").Run()
	fmt.Println("Instalando dependencia: sdl2_image")
	exec.Command("latte", "install", "sdl2_image").Run()
	fmt.Println("Instalando dependencia: mesa")
	exec.Command("latte", "install", "mesa").Run()
	fmt.Println("Instalando dependencia: mesa-glu")
	exec.Command("latte", "install", "mesa-glu").Run()
}
