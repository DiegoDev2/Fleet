package main

// AnimeDownloader - Download your favourite anime
// Homepage: https://github.com/anime-dl/anime-downloader

import (
	"fmt"
	
	"os/exec"
)

func installAnimeDownloader() {
	// Método 1: Descargar y extraer .tar.gz
	animedownloader_tar_url := "https://files.pythonhosted.org/packages/00/8b/2f354c0c2e56f1fe45e805698bd6a81c472473a48b814c44aaed2d41016d/anime-downloader-5.0.9.tar.gz"
	animedownloader_cmd_tar := exec.Command("curl", "-L", animedownloader_tar_url, "-o", "package.tar.gz")
	err := animedownloader_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	animedownloader_zip_url := "https://files.pythonhosted.org/packages/00/8b/2f354c0c2e56f1fe45e805698bd6a81c472473a48b814c44aaed2d41016d/anime-downloader-5.0.9.zip"
	animedownloader_cmd_zip := exec.Command("curl", "-L", animedownloader_zip_url, "-o", "package.zip")
	err = animedownloader_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	animedownloader_bin_url := "https://files.pythonhosted.org/packages/00/8b/2f354c0c2e56f1fe45e805698bd6a81c472473a48b814c44aaed2d41016d/anime-downloader-5.0.9.bin"
	animedownloader_cmd_bin := exec.Command("curl", "-L", animedownloader_bin_url, "-o", "binary.bin")
	err = animedownloader_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	animedownloader_src_url := "https://files.pythonhosted.org/packages/00/8b/2f354c0c2e56f1fe45e805698bd6a81c472473a48b814c44aaed2d41016d/anime-downloader-5.0.9.src.tar.gz"
	animedownloader_cmd_src := exec.Command("curl", "-L", animedownloader_src_url, "-o", "source.tar.gz")
	err = animedownloader_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	animedownloader_cmd_direct := exec.Command("./binary")
	err = animedownloader_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: aria2")
exec.Command("latte", "install", "aria2")
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
	fmt.Println("Instalando dependencia: python@3.11")
exec.Command("latte", "install", "python@3.11")
}
