package main

// Youtubeuploader - Scripted uploads to Youtube
// Homepage: https://github.com/porjo/youtubeuploader

import (
	"fmt"
	
	"os/exec"
)

func installYoutubeuploader() {
	// Método 1: Descargar y extraer .tar.gz
	youtubeuploader_tar_url := "https://github.com/porjo/youtubeuploader/archive/refs/tags/24.01.tar.gz"
	youtubeuploader_cmd_tar := exec.Command("curl", "-L", youtubeuploader_tar_url, "-o", "package.tar.gz")
	err := youtubeuploader_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	youtubeuploader_zip_url := "https://github.com/porjo/youtubeuploader/archive/refs/tags/24.01.zip"
	youtubeuploader_cmd_zip := exec.Command("curl", "-L", youtubeuploader_zip_url, "-o", "package.zip")
	err = youtubeuploader_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	youtubeuploader_bin_url := "https://github.com/porjo/youtubeuploader/archive/refs/tags/24.01.bin"
	youtubeuploader_cmd_bin := exec.Command("curl", "-L", youtubeuploader_bin_url, "-o", "binary.bin")
	err = youtubeuploader_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	youtubeuploader_src_url := "https://github.com/porjo/youtubeuploader/archive/refs/tags/24.01.src.tar.gz"
	youtubeuploader_cmd_src := exec.Command("curl", "-L", youtubeuploader_src_url, "-o", "source.tar.gz")
	err = youtubeuploader_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	youtubeuploader_cmd_direct := exec.Command("./binary")
	err = youtubeuploader_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
