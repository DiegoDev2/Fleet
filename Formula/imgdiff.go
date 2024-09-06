package main

// Imgdiff - Pixel-by-pixel image difference tool
// Homepage: https://github.com/n7olkachev/imgdiff

import (
	"fmt"
	
	"os/exec"
)

func installImgdiff() {
	// Método 1: Descargar y extraer .tar.gz
	imgdiff_tar_url := "https://github.com/n7olkachev/imgdiff/archive/refs/tags/v1.0.2.tar.gz"
	imgdiff_cmd_tar := exec.Command("curl", "-L", imgdiff_tar_url, "-o", "package.tar.gz")
	err := imgdiff_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	imgdiff_zip_url := "https://github.com/n7olkachev/imgdiff/archive/refs/tags/v1.0.2.zip"
	imgdiff_cmd_zip := exec.Command("curl", "-L", imgdiff_zip_url, "-o", "package.zip")
	err = imgdiff_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	imgdiff_bin_url := "https://github.com/n7olkachev/imgdiff/archive/refs/tags/v1.0.2.bin"
	imgdiff_cmd_bin := exec.Command("curl", "-L", imgdiff_bin_url, "-o", "binary.bin")
	err = imgdiff_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	imgdiff_src_url := "https://github.com/n7olkachev/imgdiff/archive/refs/tags/v1.0.2.src.tar.gz"
	imgdiff_cmd_src := exec.Command("curl", "-L", imgdiff_src_url, "-o", "source.tar.gz")
	err = imgdiff_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	imgdiff_cmd_direct := exec.Command("./binary")
	err = imgdiff_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
