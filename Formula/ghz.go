package main

// Ghz - Simple gRPC benchmarking and load testing tool
// Homepage: https://ghz.sh

import (
	"fmt"
	
	"os/exec"
)

func installGhz() {
	// Método 1: Descargar y extraer .tar.gz
	ghz_tar_url := "https://github.com/bojand/ghz/archive/refs/tags/v0.120.0.tar.gz"
	ghz_cmd_tar := exec.Command("curl", "-L", ghz_tar_url, "-o", "package.tar.gz")
	err := ghz_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ghz_zip_url := "https://github.com/bojand/ghz/archive/refs/tags/v0.120.0.zip"
	ghz_cmd_zip := exec.Command("curl", "-L", ghz_zip_url, "-o", "package.zip")
	err = ghz_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ghz_bin_url := "https://github.com/bojand/ghz/archive/refs/tags/v0.120.0.bin"
	ghz_cmd_bin := exec.Command("curl", "-L", ghz_bin_url, "-o", "binary.bin")
	err = ghz_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ghz_src_url := "https://github.com/bojand/ghz/archive/refs/tags/v0.120.0.src.tar.gz"
	ghz_cmd_src := exec.Command("curl", "-L", ghz_src_url, "-o", "source.tar.gz")
	err = ghz_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ghz_cmd_direct := exec.Command("./binary")
	err = ghz_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
