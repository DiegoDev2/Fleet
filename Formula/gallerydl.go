package main

// GalleryDl - Command-line downloader for image-hosting site galleries and collections
// Homepage: https://github.com/mikf/gallery-dl

import (
	"fmt"
	
	"os/exec"
)

func installGalleryDl() {
	// Método 1: Descargar y extraer .tar.gz
	gallerydl_tar_url := "https://files.pythonhosted.org/packages/a9/9e/c5bd12aed4f1bb46bbcc2279ca235b2f4e84b17f58f574400bb4bb8c1798/gallery_dl-1.27.3.tar.gz"
	gallerydl_cmd_tar := exec.Command("curl", "-L", gallerydl_tar_url, "-o", "package.tar.gz")
	err := gallerydl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gallerydl_zip_url := "https://files.pythonhosted.org/packages/a9/9e/c5bd12aed4f1bb46bbcc2279ca235b2f4e84b17f58f574400bb4bb8c1798/gallery_dl-1.27.3.zip"
	gallerydl_cmd_zip := exec.Command("curl", "-L", gallerydl_zip_url, "-o", "package.zip")
	err = gallerydl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gallerydl_bin_url := "https://files.pythonhosted.org/packages/a9/9e/c5bd12aed4f1bb46bbcc2279ca235b2f4e84b17f58f574400bb4bb8c1798/gallery_dl-1.27.3.bin"
	gallerydl_cmd_bin := exec.Command("curl", "-L", gallerydl_bin_url, "-o", "binary.bin")
	err = gallerydl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gallerydl_src_url := "https://files.pythonhosted.org/packages/a9/9e/c5bd12aed4f1bb46bbcc2279ca235b2f4e84b17f58f574400bb4bb8c1798/gallery_dl-1.27.3.src.tar.gz"
	gallerydl_cmd_src := exec.Command("curl", "-L", gallerydl_src_url, "-o", "source.tar.gz")
	err = gallerydl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gallerydl_cmd_direct := exec.Command("./binary")
	err = gallerydl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
