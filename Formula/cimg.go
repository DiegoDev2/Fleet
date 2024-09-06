package main

// Cimg - C++ toolkit for image processing
// Homepage: https://cimg.eu/

import (
	"fmt"
	
	"os/exec"
)

func installCimg() {
	// Método 1: Descargar y extraer .tar.gz
	cimg_tar_url := "https://cimg.eu/files/CImg_3.4.2.zip"
	cimg_cmd_tar := exec.Command("curl", "-L", cimg_tar_url, "-o", "package.tar.gz")
	err := cimg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cimg_zip_url := "https://cimg.eu/files/CImg_3.4.2.zip"
	cimg_cmd_zip := exec.Command("curl", "-L", cimg_zip_url, "-o", "package.zip")
	err = cimg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cimg_bin_url := "https://cimg.eu/files/CImg_3.4.2.zip"
	cimg_cmd_bin := exec.Command("curl", "-L", cimg_bin_url, "-o", "binary.bin")
	err = cimg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cimg_src_url := "https://cimg.eu/files/CImg_3.4.2.zip"
	cimg_cmd_src := exec.Command("curl", "-L", cimg_src_url, "-o", "source.tar.gz")
	err = cimg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cimg_cmd_direct := exec.Command("./binary")
	err = cimg_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
