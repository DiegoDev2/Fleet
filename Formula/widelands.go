package main

// Widelands - Free real-time strategy game like Settlers II
// Homepage: https://www.widelands.org/

import (
	"fmt"
	
	"os/exec"
)

func installWidelands() {
	// Método 1: Descargar y extraer .tar.gz
	widelands_tar_url := "https://github.com/widelands/widelands/archive/refs/tags/v1.2.tar.gz"
	widelands_cmd_tar := exec.Command("curl", "-L", widelands_tar_url, "-o", "package.tar.gz")
	err := widelands_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	widelands_zip_url := "https://github.com/widelands/widelands/archive/refs/tags/v1.2.zip"
	widelands_cmd_zip := exec.Command("curl", "-L", widelands_zip_url, "-o", "package.zip")
	err = widelands_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	widelands_bin_url := "https://github.com/widelands/widelands/archive/refs/tags/v1.2.bin"
	widelands_cmd_bin := exec.Command("curl", "-L", widelands_bin_url, "-o", "binary.bin")
	err = widelands_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	widelands_src_url := "https://github.com/widelands/widelands/archive/refs/tags/v1.2.src.tar.gz"
	widelands_cmd_src := exec.Command("curl", "-L", widelands_src_url, "-o", "source.tar.gz")
	err = widelands_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	widelands_cmd_direct := exec.Command("./binary")
	err = widelands_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: asio")
	exec.Command("latte", "install", "asio").Run()
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: doxygen")
	exec.Command("latte", "install", "doxygen").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: glew")
	exec.Command("latte", "install", "glew").Run()
	fmt.Println("Instalando dependencia: icu4c")
	exec.Command("latte", "install", "icu4c").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: lua")
	exec.Command("latte", "install", "lua").Run()
	fmt.Println("Instalando dependencia: minizip")
	exec.Command("latte", "install", "minizip").Run()
	fmt.Println("Instalando dependencia: sdl2")
	exec.Command("latte", "install", "sdl2").Run()
	fmt.Println("Instalando dependencia: sdl2_image")
	exec.Command("latte", "install", "sdl2_image").Run()
	fmt.Println("Instalando dependencia: sdl2_mixer")
	exec.Command("latte", "install", "sdl2_mixer").Run()
	fmt.Println("Instalando dependencia: sdl2_ttf")
	exec.Command("latte", "install", "sdl2_ttf").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: mesa")
	exec.Command("latte", "install", "mesa").Run()
}
