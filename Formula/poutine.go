package main

// Poutine - Security scanner that detects vulnerabilities in build pipelines
// Homepage: https://boostsecurityio.github.io/poutine/

import (
	"fmt"
	
	"os/exec"
)

func installPoutine() {
	// Método 1: Descargar y extraer .tar.gz
	poutine_tar_url := "https://github.com/boostsecurityio/poutine/archive/refs/tags/v0.15.0.tar.gz"
	poutine_cmd_tar := exec.Command("curl", "-L", poutine_tar_url, "-o", "package.tar.gz")
	err := poutine_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	poutine_zip_url := "https://github.com/boostsecurityio/poutine/archive/refs/tags/v0.15.0.zip"
	poutine_cmd_zip := exec.Command("curl", "-L", poutine_zip_url, "-o", "package.zip")
	err = poutine_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	poutine_bin_url := "https://github.com/boostsecurityio/poutine/archive/refs/tags/v0.15.0.bin"
	poutine_cmd_bin := exec.Command("curl", "-L", poutine_bin_url, "-o", "binary.bin")
	err = poutine_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	poutine_src_url := "https://github.com/boostsecurityio/poutine/archive/refs/tags/v0.15.0.src.tar.gz"
	poutine_cmd_src := exec.Command("curl", "-L", poutine_src_url, "-o", "source.tar.gz")
	err = poutine_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	poutine_cmd_direct := exec.Command("./binary")
	err = poutine_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
