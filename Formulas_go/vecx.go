package main

// Vecx - Vectrex emulator
// Homepage: https://github.com/jhawthorn/vecx

import (
	"fmt"
	
	"os/exec"
)

func installVecx() {
	// Método 1: Descargar y extraer .tar.gz
	vecx_tar_url := "https://github.com/jhawthorn/vecx/archive/refs/tags/v1.1.tar.gz"
	vecx_cmd_tar := exec.Command("curl", "-L", vecx_tar_url, "-o", "package.tar.gz")
	err := vecx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vecx_zip_url := "https://github.com/jhawthorn/vecx/archive/refs/tags/v1.1.zip"
	vecx_cmd_zip := exec.Command("curl", "-L", vecx_zip_url, "-o", "package.zip")
	err = vecx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vecx_bin_url := "https://github.com/jhawthorn/vecx/archive/refs/tags/v1.1.bin"
	vecx_cmd_bin := exec.Command("curl", "-L", vecx_bin_url, "-o", "binary.bin")
	err = vecx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vecx_src_url := "https://github.com/jhawthorn/vecx/archive/refs/tags/v1.1.src.tar.gz"
	vecx_cmd_src := exec.Command("curl", "-L", vecx_src_url, "-o", "source.tar.gz")
	err = vecx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vecx_cmd_direct := exec.Command("./binary")
	err = vecx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: sdl12-compat")
exec.Command("latte", "install", "sdl12-compat")
	fmt.Println("Instalando dependencia: sdl_gfx")
exec.Command("latte", "install", "sdl_gfx")
	fmt.Println("Instalando dependencia: sdl_image")
exec.Command("latte", "install", "sdl_image")
}
