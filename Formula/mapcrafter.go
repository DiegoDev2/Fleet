package main

// Mapcrafter - Minecraft map renderer
// Homepage: https://mapcrafter.org

import (
	"fmt"
	
	"os/exec"
)

func installMapcrafter() {
	// Método 1: Descargar y extraer .tar.gz
	mapcrafter_tar_url := "https://github.com/mapcrafter/mapcrafter/archive/refs/tags/v.2.4.tar.gz"
	mapcrafter_cmd_tar := exec.Command("curl", "-L", mapcrafter_tar_url, "-o", "package.tar.gz")
	err := mapcrafter_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mapcrafter_zip_url := "https://github.com/mapcrafter/mapcrafter/archive/refs/tags/v.2.4.zip"
	mapcrafter_cmd_zip := exec.Command("curl", "-L", mapcrafter_zip_url, "-o", "package.zip")
	err = mapcrafter_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mapcrafter_bin_url := "https://github.com/mapcrafter/mapcrafter/archive/refs/tags/v.2.4.bin"
	mapcrafter_cmd_bin := exec.Command("curl", "-L", mapcrafter_bin_url, "-o", "binary.bin")
	err = mapcrafter_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mapcrafter_src_url := "https://github.com/mapcrafter/mapcrafter/archive/refs/tags/v.2.4.src.tar.gz"
	mapcrafter_cmd_src := exec.Command("curl", "-L", mapcrafter_src_url, "-o", "source.tar.gz")
	err = mapcrafter_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mapcrafter_cmd_direct := exec.Command("./binary")
	err = mapcrafter_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
}
