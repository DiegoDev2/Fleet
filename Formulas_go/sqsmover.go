package main

// Sqsmover - AWS SQS Message mover
// Homepage: https://github.com/mercury2269/sqsmover

import (
	"fmt"
	
	"os/exec"
)

func installSqsmover() {
	// Método 1: Descargar y extraer .tar.gz
	sqsmover_tar_url := "https://github.com/mercury2269/sqsmover/archive/refs/tags/v0.4.0.tar.gz"
	sqsmover_cmd_tar := exec.Command("curl", "-L", sqsmover_tar_url, "-o", "package.tar.gz")
	err := sqsmover_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sqsmover_zip_url := "https://github.com/mercury2269/sqsmover/archive/refs/tags/v0.4.0.zip"
	sqsmover_cmd_zip := exec.Command("curl", "-L", sqsmover_zip_url, "-o", "package.zip")
	err = sqsmover_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sqsmover_bin_url := "https://github.com/mercury2269/sqsmover/archive/refs/tags/v0.4.0.bin"
	sqsmover_cmd_bin := exec.Command("curl", "-L", sqsmover_bin_url, "-o", "binary.bin")
	err = sqsmover_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sqsmover_src_url := "https://github.com/mercury2269/sqsmover/archive/refs/tags/v0.4.0.src.tar.gz"
	sqsmover_cmd_src := exec.Command("curl", "-L", sqsmover_src_url, "-o", "source.tar.gz")
	err = sqsmover_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sqsmover_cmd_direct := exec.Command("./binary")
	err = sqsmover_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
