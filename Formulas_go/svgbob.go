package main

// Svgbob - Convert your ascii diagram scribbles into happy little SVG
// Homepage: https://ivanceras.github.io/svgbob-editor/

import (
	"fmt"
	
	"os/exec"
)

func installSvgbob() {
	// Método 1: Descargar y extraer .tar.gz
	svgbob_tar_url := "https://github.com/ivanceras/svgbob/archive/refs/tags/0.7.2.tar.gz"
	svgbob_cmd_tar := exec.Command("curl", "-L", svgbob_tar_url, "-o", "package.tar.gz")
	err := svgbob_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	svgbob_zip_url := "https://github.com/ivanceras/svgbob/archive/refs/tags/0.7.2.zip"
	svgbob_cmd_zip := exec.Command("curl", "-L", svgbob_zip_url, "-o", "package.zip")
	err = svgbob_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	svgbob_bin_url := "https://github.com/ivanceras/svgbob/archive/refs/tags/0.7.2.bin"
	svgbob_cmd_bin := exec.Command("curl", "-L", svgbob_bin_url, "-o", "binary.bin")
	err = svgbob_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	svgbob_src_url := "https://github.com/ivanceras/svgbob/archive/refs/tags/0.7.2.src.tar.gz"
	svgbob_cmd_src := exec.Command("curl", "-L", svgbob_src_url, "-o", "source.tar.gz")
	err = svgbob_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	svgbob_cmd_direct := exec.Command("./binary")
	err = svgbob_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
