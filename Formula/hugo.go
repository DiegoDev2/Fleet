package main

// Hugo - Configurable static site generator
// Homepage: https://gohugo.io/

import (
	"fmt"
	
	"os/exec"
)

func installHugo() {
	// Método 1: Descargar y extraer .tar.gz
	hugo_tar_url := "https://github.com/gohugoio/hugo/archive/refs/tags/v0.134.1.tar.gz"
	hugo_cmd_tar := exec.Command("curl", "-L", hugo_tar_url, "-o", "package.tar.gz")
	err := hugo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hugo_zip_url := "https://github.com/gohugoio/hugo/archive/refs/tags/v0.134.1.zip"
	hugo_cmd_zip := exec.Command("curl", "-L", hugo_zip_url, "-o", "package.zip")
	err = hugo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hugo_bin_url := "https://github.com/gohugoio/hugo/archive/refs/tags/v0.134.1.bin"
	hugo_cmd_bin := exec.Command("curl", "-L", hugo_bin_url, "-o", "binary.bin")
	err = hugo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hugo_src_url := "https://github.com/gohugoio/hugo/archive/refs/tags/v0.134.1.src.tar.gz"
	hugo_cmd_src := exec.Command("curl", "-L", hugo_src_url, "-o", "source.tar.gz")
	err = hugo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hugo_cmd_direct := exec.Command("./binary")
	err = hugo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
