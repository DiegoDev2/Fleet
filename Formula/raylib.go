package main

// Raylib - Simple and easy-to-use library to learn videogames programming
// Homepage: https://www.raylib.com/

import (
	"fmt"
	
	"os/exec"
)

func installRaylib() {
	// Método 1: Descargar y extraer .tar.gz
	raylib_tar_url := "https://github.com/raysan5/raylib/archive/refs/tags/5.0.tar.gz"
	raylib_cmd_tar := exec.Command("curl", "-L", raylib_tar_url, "-o", "package.tar.gz")
	err := raylib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	raylib_zip_url := "https://github.com/raysan5/raylib/archive/refs/tags/5.0.zip"
	raylib_cmd_zip := exec.Command("curl", "-L", raylib_zip_url, "-o", "package.zip")
	err = raylib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	raylib_bin_url := "https://github.com/raysan5/raylib/archive/refs/tags/5.0.bin"
	raylib_cmd_bin := exec.Command("curl", "-L", raylib_bin_url, "-o", "binary.bin")
	err = raylib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	raylib_src_url := "https://github.com/raysan5/raylib/archive/refs/tags/5.0.src.tar.gz"
	raylib_cmd_src := exec.Command("curl", "-L", raylib_src_url, "-o", "source.tar.gz")
	err = raylib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	raylib_cmd_direct := exec.Command("./binary")
	err = raylib_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: libxcursor")
	exec.Command("latte", "install", "libxcursor").Run()
	fmt.Println("Instalando dependencia: libxi")
	exec.Command("latte", "install", "libxi").Run()
	fmt.Println("Instalando dependencia: libxinerama")
	exec.Command("latte", "install", "libxinerama").Run()
	fmt.Println("Instalando dependencia: libxrandr")
	exec.Command("latte", "install", "libxrandr").Run()
	fmt.Println("Instalando dependencia: mesa")
	exec.Command("latte", "install", "mesa").Run()
	fmt.Println("Instalando dependencia: mesa-glu")
	exec.Command("latte", "install", "mesa-glu").Run()
}
