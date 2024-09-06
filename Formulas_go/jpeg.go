package main

// Jpeg - Image manipulation library
// Homepage: https://www.ijg.org/

import (
	"fmt"
	
	"os/exec"
)

func installJpeg() {
	// Método 1: Descargar y extraer .tar.gz
	jpeg_tar_url := "https://www.ijg.org/files/jpegsrc.v9f.tar.gz"
	jpeg_cmd_tar := exec.Command("curl", "-L", jpeg_tar_url, "-o", "package.tar.gz")
	err := jpeg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jpeg_zip_url := "https://www.ijg.org/files/jpegsrc.v9f.zip"
	jpeg_cmd_zip := exec.Command("curl", "-L", jpeg_zip_url, "-o", "package.zip")
	err = jpeg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jpeg_bin_url := "https://www.ijg.org/files/jpegsrc.v9f.bin"
	jpeg_cmd_bin := exec.Command("curl", "-L", jpeg_bin_url, "-o", "binary.bin")
	err = jpeg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jpeg_src_url := "https://www.ijg.org/files/jpegsrc.v9f.src.tar.gz"
	jpeg_cmd_src := exec.Command("curl", "-L", jpeg_src_url, "-o", "source.tar.gz")
	err = jpeg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jpeg_cmd_direct := exec.Command("./binary")
	err = jpeg_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
