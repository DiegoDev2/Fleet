package main

// Neofetch - Fast, highly customisable system info script
// Homepage: https://github.com/dylanaraps/neofetch

import (
	"fmt"
	
	"os/exec"
)

func installNeofetch() {
	// Método 1: Descargar y extraer .tar.gz
	neofetch_tar_url := "https://github.com/dylanaraps/neofetch/archive/refs/tags/7.1.0.tar.gz"
	neofetch_cmd_tar := exec.Command("curl", "-L", neofetch_tar_url, "-o", "package.tar.gz")
	err := neofetch_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	neofetch_zip_url := "https://github.com/dylanaraps/neofetch/archive/refs/tags/7.1.0.zip"
	neofetch_cmd_zip := exec.Command("curl", "-L", neofetch_zip_url, "-o", "package.zip")
	err = neofetch_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	neofetch_bin_url := "https://github.com/dylanaraps/neofetch/archive/refs/tags/7.1.0.bin"
	neofetch_cmd_bin := exec.Command("curl", "-L", neofetch_bin_url, "-o", "binary.bin")
	err = neofetch_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	neofetch_src_url := "https://github.com/dylanaraps/neofetch/archive/refs/tags/7.1.0.src.tar.gz"
	neofetch_cmd_src := exec.Command("curl", "-L", neofetch_src_url, "-o", "source.tar.gz")
	err = neofetch_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	neofetch_cmd_direct := exec.Command("./binary")
	err = neofetch_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: screenresolution")
exec.Command("latte", "install", "screenresolution")
}
