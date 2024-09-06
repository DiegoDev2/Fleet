package main

// GetFlashVideos - Download or play videos from various Flash-based websites
// Homepage: https://github.com/monsieurvideo/get-flash-videos

import (
	"fmt"
	
	"os/exec"
)

func installGetFlashVideos() {
	// Método 1: Descargar y extraer .tar.gz
	getflashvideos_tar_url := "https://github.com/monsieurvideo/get-flash-videos/archive/refs/tags/1.25.99.03.tar.gz"
	getflashvideos_cmd_tar := exec.Command("curl", "-L", getflashvideos_tar_url, "-o", "package.tar.gz")
	err := getflashvideos_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	getflashvideos_zip_url := "https://github.com/monsieurvideo/get-flash-videos/archive/refs/tags/1.25.99.03.zip"
	getflashvideos_cmd_zip := exec.Command("curl", "-L", getflashvideos_zip_url, "-o", "package.zip")
	err = getflashvideos_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	getflashvideos_bin_url := "https://github.com/monsieurvideo/get-flash-videos/archive/refs/tags/1.25.99.03.bin"
	getflashvideos_cmd_bin := exec.Command("curl", "-L", getflashvideos_bin_url, "-o", "binary.bin")
	err = getflashvideos_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	getflashvideos_src_url := "https://github.com/monsieurvideo/get-flash-videos/archive/refs/tags/1.25.99.03.src.tar.gz"
	getflashvideos_cmd_src := exec.Command("curl", "-L", getflashvideos_src_url, "-o", "source.tar.gz")
	err = getflashvideos_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	getflashvideos_cmd_direct := exec.Command("./binary")
	err = getflashvideos_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rtmpdump")
exec.Command("latte", "install", "rtmpdump")
}
