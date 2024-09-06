package main

// Yewtube - Terminal based YouTube player and downloader
// Homepage: https://github.com/mps-youtube/yewtube

import (
	"fmt"
	
	"os/exec"
)

func installYewtube() {
	// Método 1: Descargar y extraer .tar.gz
	yewtube_tar_url := "https://github.com/mps-youtube/yewtube/archive/refs/tags/v2.10.5.tar.gz"
	yewtube_cmd_tar := exec.Command("curl", "-L", yewtube_tar_url, "-o", "package.tar.gz")
	err := yewtube_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	yewtube_zip_url := "https://github.com/mps-youtube/yewtube/archive/refs/tags/v2.10.5.zip"
	yewtube_cmd_zip := exec.Command("curl", "-L", yewtube_zip_url, "-o", "package.zip")
	err = yewtube_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	yewtube_bin_url := "https://github.com/mps-youtube/yewtube/archive/refs/tags/v2.10.5.bin"
	yewtube_cmd_bin := exec.Command("curl", "-L", yewtube_bin_url, "-o", "binary.bin")
	err = yewtube_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	yewtube_src_url := "https://github.com/mps-youtube/yewtube/archive/refs/tags/v2.10.5.src.tar.gz"
	yewtube_cmd_src := exec.Command("curl", "-L", yewtube_src_url, "-o", "source.tar.gz")
	err = yewtube_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	yewtube_cmd_direct := exec.Command("./binary")
	err = yewtube_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ffmpeg")
exec.Command("latte", "install", "ffmpeg")
	fmt.Println("Instalando dependencia: mplayer")
exec.Command("latte", "install", "mplayer")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
