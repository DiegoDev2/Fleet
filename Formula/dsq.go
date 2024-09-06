package main

// Dsq - CLI tool for running SQL queries against JSON, CSV, Excel, Parquet, and more
// Homepage: https://github.com/multiprocessio/dsq

import (
	"fmt"
	
	"os/exec"
)

func installDsq() {
	// Método 1: Descargar y extraer .tar.gz
	dsq_tar_url := "https://github.com/multiprocessio/dsq/archive/refs/tags/v0.23.0.tar.gz"
	dsq_cmd_tar := exec.Command("curl", "-L", dsq_tar_url, "-o", "package.tar.gz")
	err := dsq_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dsq_zip_url := "https://github.com/multiprocessio/dsq/archive/refs/tags/v0.23.0.zip"
	dsq_cmd_zip := exec.Command("curl", "-L", dsq_zip_url, "-o", "package.zip")
	err = dsq_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dsq_bin_url := "https://github.com/multiprocessio/dsq/archive/refs/tags/v0.23.0.bin"
	dsq_cmd_bin := exec.Command("curl", "-L", dsq_bin_url, "-o", "binary.bin")
	err = dsq_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dsq_src_url := "https://github.com/multiprocessio/dsq/archive/refs/tags/v0.23.0.src.tar.gz"
	dsq_cmd_src := exec.Command("curl", "-L", dsq_src_url, "-o", "source.tar.gz")
	err = dsq_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dsq_cmd_direct := exec.Command("./binary")
	err = dsq_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
