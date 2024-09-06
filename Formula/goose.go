package main

// Goose - Go Language's command-line interface for database migrations
// Homepage: https://pressly.github.io/goose/

import (
	"fmt"
	
	"os/exec"
)

func installGoose() {
	// Método 1: Descargar y extraer .tar.gz
	goose_tar_url := "https://github.com/pressly/goose/archive/refs/tags/v3.22.0.tar.gz"
	goose_cmd_tar := exec.Command("curl", "-L", goose_tar_url, "-o", "package.tar.gz")
	err := goose_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	goose_zip_url := "https://github.com/pressly/goose/archive/refs/tags/v3.22.0.zip"
	goose_cmd_zip := exec.Command("curl", "-L", goose_zip_url, "-o", "package.zip")
	err = goose_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	goose_bin_url := "https://github.com/pressly/goose/archive/refs/tags/v3.22.0.bin"
	goose_cmd_bin := exec.Command("curl", "-L", goose_bin_url, "-o", "binary.bin")
	err = goose_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	goose_src_url := "https://github.com/pressly/goose/archive/refs/tags/v3.22.0.src.tar.gz"
	goose_cmd_src := exec.Command("curl", "-L", goose_src_url, "-o", "source.tar.gz")
	err = goose_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	goose_cmd_direct := exec.Command("./binary")
	err = goose_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
