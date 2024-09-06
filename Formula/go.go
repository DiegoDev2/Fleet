package main

// Go - Open source programming language to build simple/reliable/efficient software
// Homepage: https://go.dev/

import (
	"fmt"
	
	"os/exec"
)

func installGo() {
	// Método 1: Descargar y extraer .tar.gz
	go_tar_url := "https://go.dev/dl/go1.23.1.src.tar.gz"
	go_cmd_tar := exec.Command("curl", "-L", go_tar_url, "-o", "package.tar.gz")
	err := go_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	go_zip_url := "https://go.dev/dl/go1.23.1.src.zip"
	go_cmd_zip := exec.Command("curl", "-L", go_zip_url, "-o", "package.zip")
	err = go_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	go_bin_url := "https://go.dev/dl/go1.23.1.src.bin"
	go_cmd_bin := exec.Command("curl", "-L", go_bin_url, "-o", "binary.bin")
	err = go_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	go_src_url := "https://go.dev/dl/go1.23.1.src.src.tar.gz"
	go_cmd_src := exec.Command("curl", "-L", go_src_url, "-o", "source.tar.gz")
	err = go_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	go_cmd_direct := exec.Command("./binary")
	err = go_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
