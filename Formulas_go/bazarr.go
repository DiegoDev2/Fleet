package main

// Bazarr - Companion to Sonarr and Radarr for managing and downloading subtitles
// Homepage: https://www.bazarr.media

import (
	"fmt"
	
	"os/exec"
)

func installBazarr() {
	// Método 1: Descargar y extraer .tar.gz
	bazarr_tar_url := "https://github.com/morpheus65535/bazarr/releases/download/v1.4.3/bazarr.zip"
	bazarr_cmd_tar := exec.Command("curl", "-L", bazarr_tar_url, "-o", "package.tar.gz")
	err := bazarr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bazarr_zip_url := "https://github.com/morpheus65535/bazarr/releases/download/v1.4.3/bazarr.zip"
	bazarr_cmd_zip := exec.Command("curl", "-L", bazarr_zip_url, "-o", "package.zip")
	err = bazarr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bazarr_bin_url := "https://github.com/morpheus65535/bazarr/releases/download/v1.4.3/bazarr.zip"
	bazarr_cmd_bin := exec.Command("curl", "-L", bazarr_bin_url, "-o", "binary.bin")
	err = bazarr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bazarr_src_url := "https://github.com/morpheus65535/bazarr/releases/download/v1.4.3/bazarr.zip"
	bazarr_cmd_src := exec.Command("curl", "-L", bazarr_src_url, "-o", "source.tar.gz")
	err = bazarr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bazarr_cmd_direct := exec.Command("./binary")
	err = bazarr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
	fmt.Println("Instalando dependencia: ffmpeg")
exec.Command("latte", "install", "ffmpeg")
	fmt.Println("Instalando dependencia: gcc")
exec.Command("latte", "install", "gcc")
	fmt.Println("Instalando dependencia: numpy")
exec.Command("latte", "install", "numpy")
	fmt.Println("Instalando dependencia: pillow")
exec.Command("latte", "install", "pillow")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: unar")
exec.Command("latte", "install", "unar")
}
