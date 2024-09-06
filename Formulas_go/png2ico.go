package main

// Png2ico - PNG to icon converter
// Homepage: https://www.winterdrache.de/freeware/png2ico/

import (
	"fmt"
	
	"os/exec"
)

func installPng2ico() {
	// Método 1: Descargar y extraer .tar.gz
	png2ico_tar_url := "https://www.winterdrache.de/freeware/png2ico/data/png2ico-src-2002-12-08.tar.gz"
	png2ico_cmd_tar := exec.Command("curl", "-L", png2ico_tar_url, "-o", "package.tar.gz")
	err := png2ico_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	png2ico_zip_url := "https://www.winterdrache.de/freeware/png2ico/data/png2ico-src-2002-12-08.zip"
	png2ico_cmd_zip := exec.Command("curl", "-L", png2ico_zip_url, "-o", "package.zip")
	err = png2ico_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	png2ico_bin_url := "https://www.winterdrache.de/freeware/png2ico/data/png2ico-src-2002-12-08.bin"
	png2ico_cmd_bin := exec.Command("curl", "-L", png2ico_bin_url, "-o", "binary.bin")
	err = png2ico_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	png2ico_src_url := "https://www.winterdrache.de/freeware/png2ico/data/png2ico-src-2002-12-08.src.tar.gz"
	png2ico_cmd_src := exec.Command("curl", "-L", png2ico_src_url, "-o", "source.tar.gz")
	err = png2ico_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	png2ico_cmd_direct := exec.Command("./binary")
	err = png2ico_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
}
