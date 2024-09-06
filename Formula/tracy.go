package main

// Tracy - Real-time, nanosecond resolution frame profiler
// Homepage: https://github.com/wolfpld/tracy

import (
	"fmt"
	
	"os/exec"
)

func installTracy() {
	// Método 1: Descargar y extraer .tar.gz
	tracy_tar_url := "https://github.com/wolfpld/tracy/archive/refs/tags/v0.11.1.tar.gz"
	tracy_cmd_tar := exec.Command("curl", "-L", tracy_tar_url, "-o", "package.tar.gz")
	err := tracy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tracy_zip_url := "https://github.com/wolfpld/tracy/archive/refs/tags/v0.11.1.zip"
	tracy_cmd_zip := exec.Command("curl", "-L", tracy_zip_url, "-o", "package.zip")
	err = tracy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tracy_bin_url := "https://github.com/wolfpld/tracy/archive/refs/tags/v0.11.1.bin"
	tracy_cmd_bin := exec.Command("curl", "-L", tracy_bin_url, "-o", "binary.bin")
	err = tracy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tracy_src_url := "https://github.com/wolfpld/tracy/archive/refs/tags/v0.11.1.src.tar.gz"
	tracy_cmd_src := exec.Command("curl", "-L", tracy_src_url, "-o", "source.tar.gz")
	err = tracy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tracy_cmd_direct := exec.Command("./binary")
	err = tracy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: capstone")
	exec.Command("latte", "install", "capstone").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: glfw")
	exec.Command("latte", "install", "glfw").Run()
	fmt.Println("Instalando dependencia: dbus")
	exec.Command("latte", "install", "dbus").Run()
	fmt.Println("Instalando dependencia: libxkbcommon")
	exec.Command("latte", "install", "libxkbcommon").Run()
	fmt.Println("Instalando dependencia: mesa")
	exec.Command("latte", "install", "mesa").Run()
	fmt.Println("Instalando dependencia: tbb")
	exec.Command("latte", "install", "tbb").Run()
	fmt.Println("Instalando dependencia: wayland")
	exec.Command("latte", "install", "wayland").Run()
}
