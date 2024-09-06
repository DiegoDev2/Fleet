package main

// Projectm - Milkdrop-compatible music visualizer
// Homepage: https://github.com/projectM-visualizer/projectm

import (
	"fmt"
	
	"os/exec"
)

func installProjectm() {
	// Método 1: Descargar y extraer .tar.gz
	projectm_tar_url := "https://github.com/projectM-visualizer/projectm/releases/download/v3.1.12/projectM-3.1.12.tar.gz"
	projectm_cmd_tar := exec.Command("curl", "-L", projectm_tar_url, "-o", "package.tar.gz")
	err := projectm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	projectm_zip_url := "https://github.com/projectM-visualizer/projectm/releases/download/v3.1.12/projectM-3.1.12.zip"
	projectm_cmd_zip := exec.Command("curl", "-L", projectm_zip_url, "-o", "package.zip")
	err = projectm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	projectm_bin_url := "https://github.com/projectM-visualizer/projectm/releases/download/v3.1.12/projectM-3.1.12.bin"
	projectm_cmd_bin := exec.Command("curl", "-L", projectm_bin_url, "-o", "binary.bin")
	err = projectm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	projectm_src_url := "https://github.com/projectM-visualizer/projectm/releases/download/v3.1.12/projectM-3.1.12.src.tar.gz"
	projectm_cmd_src := exec.Command("curl", "-L", projectm_src_url, "-o", "source.tar.gz")
	err = projectm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	projectm_cmd_direct := exec.Command("./binary")
	err = projectm_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: sdl2")
	exec.Command("latte", "install", "sdl2").Run()
	fmt.Println("Instalando dependencia: mesa")
	exec.Command("latte", "install", "mesa").Run()
}
