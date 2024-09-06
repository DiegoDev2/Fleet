package main

// GoAT120 - Open source programming language to build simple/reliable/efficient software
// Homepage: https://go.dev/

import (
	"fmt"
	
	"os/exec"
)

func installGoAT120() {
	// Método 1: Descargar y extraer .tar.gz
	goat120_tar_url := "https://go.dev/dl/go1.20.14.src.tar.gz"
	goat120_cmd_tar := exec.Command("curl", "-L", goat120_tar_url, "-o", "package.tar.gz")
	err := goat120_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	goat120_zip_url := "https://go.dev/dl/go1.20.14.src.zip"
	goat120_cmd_zip := exec.Command("curl", "-L", goat120_zip_url, "-o", "package.zip")
	err = goat120_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	goat120_bin_url := "https://go.dev/dl/go1.20.14.src.bin"
	goat120_cmd_bin := exec.Command("curl", "-L", goat120_bin_url, "-o", "binary.bin")
	err = goat120_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	goat120_src_url := "https://go.dev/dl/go1.20.14.src.src.tar.gz"
	goat120_cmd_src := exec.Command("curl", "-L", goat120_src_url, "-o", "source.tar.gz")
	err = goat120_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	goat120_cmd_direct := exec.Command("./binary")
	err = goat120_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
