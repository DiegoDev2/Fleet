package main

// Immudb - Lightweight, high-speed immutable database
// Homepage: https://www.codenotary.io

import (
	"fmt"
	
	"os/exec"
)

func installImmudb() {
	// Método 1: Descargar y extraer .tar.gz
	immudb_tar_url := "https://github.com/codenotary/immudb/archive/refs/tags/v1.9.4.tar.gz"
	immudb_cmd_tar := exec.Command("curl", "-L", immudb_tar_url, "-o", "package.tar.gz")
	err := immudb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	immudb_zip_url := "https://github.com/codenotary/immudb/archive/refs/tags/v1.9.4.zip"
	immudb_cmd_zip := exec.Command("curl", "-L", immudb_zip_url, "-o", "package.zip")
	err = immudb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	immudb_bin_url := "https://github.com/codenotary/immudb/archive/refs/tags/v1.9.4.bin"
	immudb_cmd_bin := exec.Command("curl", "-L", immudb_bin_url, "-o", "binary.bin")
	err = immudb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	immudb_src_url := "https://github.com/codenotary/immudb/archive/refs/tags/v1.9.4.src.tar.gz"
	immudb_cmd_src := exec.Command("curl", "-L", immudb_src_url, "-o", "source.tar.gz")
	err = immudb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	immudb_cmd_direct := exec.Command("./binary")
	err = immudb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
