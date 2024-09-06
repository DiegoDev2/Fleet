package main

// Youtubedr - Download Youtube Video in Golang
// Homepage: https://github.com/kkdai/youtube

import (
	"fmt"
	
	"os/exec"
)

func installYoutubedr() {
	// Método 1: Descargar y extraer .tar.gz
	youtubedr_tar_url := "https://github.com/kkdai/youtube/archive/refs/tags/v2.10.1.tar.gz"
	youtubedr_cmd_tar := exec.Command("curl", "-L", youtubedr_tar_url, "-o", "package.tar.gz")
	err := youtubedr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	youtubedr_zip_url := "https://github.com/kkdai/youtube/archive/refs/tags/v2.10.1.zip"
	youtubedr_cmd_zip := exec.Command("curl", "-L", youtubedr_zip_url, "-o", "package.zip")
	err = youtubedr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	youtubedr_bin_url := "https://github.com/kkdai/youtube/archive/refs/tags/v2.10.1.bin"
	youtubedr_cmd_bin := exec.Command("curl", "-L", youtubedr_bin_url, "-o", "binary.bin")
	err = youtubedr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	youtubedr_src_url := "https://github.com/kkdai/youtube/archive/refs/tags/v2.10.1.src.tar.gz"
	youtubedr_cmd_src := exec.Command("curl", "-L", youtubedr_src_url, "-o", "source.tar.gz")
	err = youtubedr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	youtubedr_cmd_direct := exec.Command("./binary")
	err = youtubedr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
