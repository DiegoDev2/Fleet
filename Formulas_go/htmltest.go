package main

// Htmltest - HTML validator written in Go
// Homepage: https://github.com/wjdp/htmltest

import (
	"fmt"
	
	"os/exec"
)

func installHtmltest() {
	// Método 1: Descargar y extraer .tar.gz
	htmltest_tar_url := "https://github.com/wjdp/htmltest/archive/refs/tags/v0.17.0.tar.gz"
	htmltest_cmd_tar := exec.Command("curl", "-L", htmltest_tar_url, "-o", "package.tar.gz")
	err := htmltest_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	htmltest_zip_url := "https://github.com/wjdp/htmltest/archive/refs/tags/v0.17.0.zip"
	htmltest_cmd_zip := exec.Command("curl", "-L", htmltest_zip_url, "-o", "package.zip")
	err = htmltest_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	htmltest_bin_url := "https://github.com/wjdp/htmltest/archive/refs/tags/v0.17.0.bin"
	htmltest_cmd_bin := exec.Command("curl", "-L", htmltest_bin_url, "-o", "binary.bin")
	err = htmltest_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	htmltest_src_url := "https://github.com/wjdp/htmltest/archive/refs/tags/v0.17.0.src.tar.gz"
	htmltest_cmd_src := exec.Command("curl", "-L", htmltest_src_url, "-o", "source.tar.gz")
	err = htmltest_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	htmltest_cmd_direct := exec.Command("./binary")
	err = htmltest_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
