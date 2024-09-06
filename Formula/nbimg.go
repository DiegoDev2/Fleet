package main

// Nbimg - Smartphone boot splash screen converter for Android and winCE
// Homepage: https://github.com/poliva/nbimg

import (
	"fmt"
	
	"os/exec"
)

func installNbimg() {
	// Método 1: Descargar y extraer .tar.gz
	nbimg_tar_url := "https://github.com/poliva/nbimg/archive/refs/tags/v1.2.1.tar.gz"
	nbimg_cmd_tar := exec.Command("curl", "-L", nbimg_tar_url, "-o", "package.tar.gz")
	err := nbimg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nbimg_zip_url := "https://github.com/poliva/nbimg/archive/refs/tags/v1.2.1.zip"
	nbimg_cmd_zip := exec.Command("curl", "-L", nbimg_zip_url, "-o", "package.zip")
	err = nbimg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nbimg_bin_url := "https://github.com/poliva/nbimg/archive/refs/tags/v1.2.1.bin"
	nbimg_cmd_bin := exec.Command("curl", "-L", nbimg_bin_url, "-o", "binary.bin")
	err = nbimg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nbimg_src_url := "https://github.com/poliva/nbimg/archive/refs/tags/v1.2.1.src.tar.gz"
	nbimg_cmd_src := exec.Command("curl", "-L", nbimg_src_url, "-o", "source.tar.gz")
	err = nbimg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nbimg_cmd_direct := exec.Command("./binary")
	err = nbimg_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
