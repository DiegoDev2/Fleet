package main

// Esbuild - Extremely fast JavaScript bundler and minifier
// Homepage: https://esbuild.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installEsbuild() {
	// Método 1: Descargar y extraer .tar.gz
	esbuild_tar_url := "https://github.com/evanw/esbuild/archive/refs/tags/v0.23.1.tar.gz"
	esbuild_cmd_tar := exec.Command("curl", "-L", esbuild_tar_url, "-o", "package.tar.gz")
	err := esbuild_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	esbuild_zip_url := "https://github.com/evanw/esbuild/archive/refs/tags/v0.23.1.zip"
	esbuild_cmd_zip := exec.Command("curl", "-L", esbuild_zip_url, "-o", "package.zip")
	err = esbuild_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	esbuild_bin_url := "https://github.com/evanw/esbuild/archive/refs/tags/v0.23.1.bin"
	esbuild_cmd_bin := exec.Command("curl", "-L", esbuild_bin_url, "-o", "binary.bin")
	err = esbuild_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	esbuild_src_url := "https://github.com/evanw/esbuild/archive/refs/tags/v0.23.1.src.tar.gz"
	esbuild_cmd_src := exec.Command("curl", "-L", esbuild_src_url, "-o", "source.tar.gz")
	err = esbuild_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	esbuild_cmd_direct := exec.Command("./binary")
	err = esbuild_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
