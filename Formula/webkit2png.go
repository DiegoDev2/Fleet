package main

// Webkit2png - Create screenshots of webpages from the terminal
// Homepage: https://www.paulhammond.org/webkit2png/

import (
	"fmt"
	
	"os/exec"
)

func installWebkit2png() {
	// Método 1: Descargar y extraer .tar.gz
	webkit2png_tar_url := "https://github.com/paulhammond/webkit2png/archive/refs/tags/v0.7.tar.gz"
	webkit2png_cmd_tar := exec.Command("curl", "-L", webkit2png_tar_url, "-o", "package.tar.gz")
	err := webkit2png_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	webkit2png_zip_url := "https://github.com/paulhammond/webkit2png/archive/refs/tags/v0.7.zip"
	webkit2png_cmd_zip := exec.Command("curl", "-L", webkit2png_zip_url, "-o", "package.zip")
	err = webkit2png_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	webkit2png_bin_url := "https://github.com/paulhammond/webkit2png/archive/refs/tags/v0.7.bin"
	webkit2png_cmd_bin := exec.Command("curl", "-L", webkit2png_bin_url, "-o", "binary.bin")
	err = webkit2png_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	webkit2png_src_url := "https://github.com/paulhammond/webkit2png/archive/refs/tags/v0.7.src.tar.gz"
	webkit2png_cmd_src := exec.Command("curl", "-L", webkit2png_src_url, "-o", "source.tar.gz")
	err = webkit2png_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	webkit2png_cmd_direct := exec.Command("./binary")
	err = webkit2png_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
