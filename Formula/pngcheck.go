package main

// Pngcheck - Print info and check PNG, JNG, and MNG files
// Homepage: http://www.libpng.org/pub/png/apps/pngcheck.html

import (
	"fmt"
	
	"os/exec"
)

func installPngcheck() {
	// Método 1: Descargar y extraer .tar.gz
	pngcheck_tar_url := "http://www.libpng.org/pub/png/src/pngcheck-3.0.3.tar.gz"
	pngcheck_cmd_tar := exec.Command("curl", "-L", pngcheck_tar_url, "-o", "package.tar.gz")
	err := pngcheck_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pngcheck_zip_url := "http://www.libpng.org/pub/png/src/pngcheck-3.0.3.zip"
	pngcheck_cmd_zip := exec.Command("curl", "-L", pngcheck_zip_url, "-o", "package.zip")
	err = pngcheck_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pngcheck_bin_url := "http://www.libpng.org/pub/png/src/pngcheck-3.0.3.bin"
	pngcheck_cmd_bin := exec.Command("curl", "-L", pngcheck_bin_url, "-o", "binary.bin")
	err = pngcheck_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pngcheck_src_url := "http://www.libpng.org/pub/png/src/pngcheck-3.0.3.src.tar.gz"
	pngcheck_cmd_src := exec.Command("curl", "-L", pngcheck_src_url, "-o", "source.tar.gz")
	err = pngcheck_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pngcheck_cmd_direct := exec.Command("./binary")
	err = pngcheck_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
