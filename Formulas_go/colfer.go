package main

// Colfer - Schema compiler for binary data exchange
// Homepage: https://github.com/pascaldekloe/colfer

import (
	"fmt"
	
	"os/exec"
)

func installColfer() {
	// Método 1: Descargar y extraer .tar.gz
	colfer_tar_url := "https://github.com/pascaldekloe/colfer/archive/refs/tags/v1.8.1.tar.gz"
	colfer_cmd_tar := exec.Command("curl", "-L", colfer_tar_url, "-o", "package.tar.gz")
	err := colfer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	colfer_zip_url := "https://github.com/pascaldekloe/colfer/archive/refs/tags/v1.8.1.zip"
	colfer_cmd_zip := exec.Command("curl", "-L", colfer_zip_url, "-o", "package.zip")
	err = colfer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	colfer_bin_url := "https://github.com/pascaldekloe/colfer/archive/refs/tags/v1.8.1.bin"
	colfer_cmd_bin := exec.Command("curl", "-L", colfer_bin_url, "-o", "binary.bin")
	err = colfer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	colfer_src_url := "https://github.com/pascaldekloe/colfer/archive/refs/tags/v1.8.1.src.tar.gz"
	colfer_cmd_src := exec.Command("curl", "-L", colfer_src_url, "-o", "source.tar.gz")
	err = colfer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	colfer_cmd_direct := exec.Command("./binary")
	err = colfer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
