package main

// YoutubeDl - Download YouTube videos from the command-line
// Homepage: https://ytdl-org.github.io/youtube-dl/

import (
	"fmt"
	
	"os/exec"
)

func installYoutubeDl() {
	// Método 1: Descargar y extraer .tar.gz
	youtubedl_tar_url := "https://files.pythonhosted.org/packages/01/4f/ab0d0806f4d818168d0ec833df14078c9d1ddddb5c42fa7bfb6f15ecbfa7/youtube_dl-2021.12.17.tar.gz"
	youtubedl_cmd_tar := exec.Command("curl", "-L", youtubedl_tar_url, "-o", "package.tar.gz")
	err := youtubedl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	youtubedl_zip_url := "https://files.pythonhosted.org/packages/01/4f/ab0d0806f4d818168d0ec833df14078c9d1ddddb5c42fa7bfb6f15ecbfa7/youtube_dl-2021.12.17.zip"
	youtubedl_cmd_zip := exec.Command("curl", "-L", youtubedl_zip_url, "-o", "package.zip")
	err = youtubedl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	youtubedl_bin_url := "https://files.pythonhosted.org/packages/01/4f/ab0d0806f4d818168d0ec833df14078c9d1ddddb5c42fa7bfb6f15ecbfa7/youtube_dl-2021.12.17.bin"
	youtubedl_cmd_bin := exec.Command("curl", "-L", youtubedl_bin_url, "-o", "binary.bin")
	err = youtubedl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	youtubedl_src_url := "https://files.pythonhosted.org/packages/01/4f/ab0d0806f4d818168d0ec833df14078c9d1ddddb5c42fa7bfb6f15ecbfa7/youtube_dl-2021.12.17.src.tar.gz"
	youtubedl_cmd_src := exec.Command("curl", "-L", youtubedl_src_url, "-o", "source.tar.gz")
	err = youtubedl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	youtubedl_cmd_direct := exec.Command("./binary")
	err = youtubedl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pandoc")
exec.Command("latte", "install", "pandoc")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
