package main

// Pioneer - Game of lonely space adventure
// Homepage: https://pioneerspacesim.net/

import (
	"fmt"
	
	"os/exec"
)

func installPioneer() {
	// Método 1: Descargar y extraer .tar.gz
	pioneer_tar_url := "https://github.com/pioneerspacesim/pioneer/archive/refs/tags/20230203.tar.gz"
	pioneer_cmd_tar := exec.Command("curl", "-L", pioneer_tar_url, "-o", "package.tar.gz")
	err := pioneer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pioneer_zip_url := "https://github.com/pioneerspacesim/pioneer/archive/refs/tags/20230203.zip"
	pioneer_cmd_zip := exec.Command("curl", "-L", pioneer_zip_url, "-o", "package.zip")
	err = pioneer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pioneer_bin_url := "https://github.com/pioneerspacesim/pioneer/archive/refs/tags/20230203.bin"
	pioneer_cmd_bin := exec.Command("curl", "-L", pioneer_bin_url, "-o", "binary.bin")
	err = pioneer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pioneer_src_url := "https://github.com/pioneerspacesim/pioneer/archive/refs/tags/20230203.src.tar.gz"
	pioneer_cmd_src := exec.Command("curl", "-L", pioneer_src_url, "-o", "source.tar.gz")
	err = pioneer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pioneer_cmd_direct := exec.Command("./binary")
	err = pioneer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: assimp")
	exec.Command("latte", "install", "assimp").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: glew")
	exec.Command("latte", "install", "glew").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: libsigc++@2")
	exec.Command("latte", "install", "libsigc++@2").Run()
	fmt.Println("Instalando dependencia: libvorbis")
	exec.Command("latte", "install", "libvorbis").Run()
	fmt.Println("Instalando dependencia: sdl2")
	exec.Command("latte", "install", "sdl2").Run()
	fmt.Println("Instalando dependencia: sdl2_image")
	exec.Command("latte", "install", "sdl2_image").Run()
	fmt.Println("Instalando dependencia: mesa")
	exec.Command("latte", "install", "mesa").Run()
}
