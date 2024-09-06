package main

// Tenyr - 32-bit computing environment (including simulated CPU)
// Homepage: https://kulp.github.io/tenyr/

import (
	"fmt"
	
	"os/exec"
)

func installTenyr() {
	// Método 1: Descargar y extraer .tar.gz
	tenyr_tar_url := "https://github.com/kulp/tenyr/archive/refs/tags/v0.9.9.tar.gz"
	tenyr_cmd_tar := exec.Command("curl", "-L", tenyr_tar_url, "-o", "package.tar.gz")
	err := tenyr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tenyr_zip_url := "https://github.com/kulp/tenyr/archive/refs/tags/v0.9.9.zip"
	tenyr_cmd_zip := exec.Command("curl", "-L", tenyr_zip_url, "-o", "package.zip")
	err = tenyr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tenyr_bin_url := "https://github.com/kulp/tenyr/archive/refs/tags/v0.9.9.bin"
	tenyr_cmd_bin := exec.Command("curl", "-L", tenyr_bin_url, "-o", "binary.bin")
	err = tenyr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tenyr_src_url := "https://github.com/kulp/tenyr/archive/refs/tags/v0.9.9.src.tar.gz"
	tenyr_cmd_src := exec.Command("curl", "-L", tenyr_src_url, "-o", "source.tar.gz")
	err = tenyr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tenyr_cmd_direct := exec.Command("./binary")
	err = tenyr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bison")
exec.Command("latte", "install", "bison")
	fmt.Println("Instalando dependencia: sdl2")
exec.Command("latte", "install", "sdl2")
	fmt.Println("Instalando dependencia: sdl2_image")
exec.Command("latte", "install", "sdl2_image")
}
