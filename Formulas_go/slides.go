package main

// Slides - Terminal based presentation tool
// Homepage: https://github.com/maaslalani/slides

import (
	"fmt"
	
	"os/exec"
)

func installSlides() {
	// Método 1: Descargar y extraer .tar.gz
	slides_tar_url := "https://github.com/maaslalani/slides/archive/refs/tags/v0.9.0.tar.gz"
	slides_cmd_tar := exec.Command("curl", "-L", slides_tar_url, "-o", "package.tar.gz")
	err := slides_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	slides_zip_url := "https://github.com/maaslalani/slides/archive/refs/tags/v0.9.0.zip"
	slides_cmd_zip := exec.Command("curl", "-L", slides_zip_url, "-o", "package.zip")
	err = slides_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	slides_bin_url := "https://github.com/maaslalani/slides/archive/refs/tags/v0.9.0.bin"
	slides_cmd_bin := exec.Command("curl", "-L", slides_bin_url, "-o", "binary.bin")
	err = slides_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	slides_src_url := "https://github.com/maaslalani/slides/archive/refs/tags/v0.9.0.src.tar.gz"
	slides_cmd_src := exec.Command("curl", "-L", slides_src_url, "-o", "source.tar.gz")
	err = slides_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	slides_cmd_direct := exec.Command("./binary")
	err = slides_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
