package main

// SpacemanDiff - Diff images from the command-line
// Homepage: https://github.com/holman/spaceman-diff

import (
	"fmt"
	
	"os/exec"
)

func installSpacemanDiff() {
	// Método 1: Descargar y extraer .tar.gz
	spacemandiff_tar_url := "https://github.com/holman/spaceman-diff/archive/refs/tags/v1.0.3.tar.gz"
	spacemandiff_cmd_tar := exec.Command("curl", "-L", spacemandiff_tar_url, "-o", "package.tar.gz")
	err := spacemandiff_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	spacemandiff_zip_url := "https://github.com/holman/spaceman-diff/archive/refs/tags/v1.0.3.zip"
	spacemandiff_cmd_zip := exec.Command("curl", "-L", spacemandiff_zip_url, "-o", "package.zip")
	err = spacemandiff_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	spacemandiff_bin_url := "https://github.com/holman/spaceman-diff/archive/refs/tags/v1.0.3.bin"
	spacemandiff_cmd_bin := exec.Command("curl", "-L", spacemandiff_bin_url, "-o", "binary.bin")
	err = spacemandiff_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	spacemandiff_src_url := "https://github.com/holman/spaceman-diff/archive/refs/tags/v1.0.3.src.tar.gz"
	spacemandiff_cmd_src := exec.Command("curl", "-L", spacemandiff_src_url, "-o", "source.tar.gz")
	err = spacemandiff_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	spacemandiff_cmd_direct := exec.Command("./binary")
	err = spacemandiff_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: imagemagick")
	exec.Command("latte", "install", "imagemagick").Run()
	fmt.Println("Instalando dependencia: jp2a")
	exec.Command("latte", "install", "jp2a").Run()
}
