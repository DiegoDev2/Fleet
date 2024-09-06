package main

// Ttygif - Converts a ttyrec file into gif files
// Homepage: https://github.com/icholy/ttygif

import (
	"fmt"
	
	"os/exec"
)

func installTtygif() {
	// Método 1: Descargar y extraer .tar.gz
	ttygif_tar_url := "https://github.com/icholy/ttygif/archive/refs/tags/1.6.0.tar.gz"
	ttygif_cmd_tar := exec.Command("curl", "-L", ttygif_tar_url, "-o", "package.tar.gz")
	err := ttygif_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ttygif_zip_url := "https://github.com/icholy/ttygif/archive/refs/tags/1.6.0.zip"
	ttygif_cmd_zip := exec.Command("curl", "-L", ttygif_zip_url, "-o", "package.zip")
	err = ttygif_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ttygif_bin_url := "https://github.com/icholy/ttygif/archive/refs/tags/1.6.0.bin"
	ttygif_cmd_bin := exec.Command("curl", "-L", ttygif_bin_url, "-o", "binary.bin")
	err = ttygif_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ttygif_src_url := "https://github.com/icholy/ttygif/archive/refs/tags/1.6.0.src.tar.gz"
	ttygif_cmd_src := exec.Command("curl", "-L", ttygif_src_url, "-o", "source.tar.gz")
	err = ttygif_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ttygif_cmd_direct := exec.Command("./binary")
	err = ttygif_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: imagemagick")
exec.Command("latte", "install", "imagemagick")
	fmt.Println("Instalando dependencia: ttyrec")
exec.Command("latte", "install", "ttyrec")
}
