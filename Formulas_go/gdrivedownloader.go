package main

// GdriveDownloader - Download a gdrive folder or file easily, shell ftw
// Homepage: https://github.com/Akianonymus/gdrive-downloader

import (
	"fmt"
	
	"os/exec"
)

func installGdriveDownloader() {
	// Método 1: Descargar y extraer .tar.gz
	gdrivedownloader_tar_url := "https://github.com/Akianonymus/gdrive-downloader/archive/refs/tags/v1.1.tar.gz"
	gdrivedownloader_cmd_tar := exec.Command("curl", "-L", gdrivedownloader_tar_url, "-o", "package.tar.gz")
	err := gdrivedownloader_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gdrivedownloader_zip_url := "https://github.com/Akianonymus/gdrive-downloader/archive/refs/tags/v1.1.zip"
	gdrivedownloader_cmd_zip := exec.Command("curl", "-L", gdrivedownloader_zip_url, "-o", "package.zip")
	err = gdrivedownloader_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gdrivedownloader_bin_url := "https://github.com/Akianonymus/gdrive-downloader/archive/refs/tags/v1.1.bin"
	gdrivedownloader_cmd_bin := exec.Command("curl", "-L", gdrivedownloader_bin_url, "-o", "binary.bin")
	err = gdrivedownloader_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gdrivedownloader_src_url := "https://github.com/Akianonymus/gdrive-downloader/archive/refs/tags/v1.1.src.tar.gz"
	gdrivedownloader_cmd_src := exec.Command("curl", "-L", gdrivedownloader_src_url, "-o", "source.tar.gz")
	err = gdrivedownloader_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gdrivedownloader_cmd_direct := exec.Command("./binary")
	err = gdrivedownloader_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
