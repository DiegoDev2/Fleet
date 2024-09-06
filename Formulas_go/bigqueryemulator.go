package main

// BigqueryEmulator - Emulate a GCP BigQuery server on your local machine
// Homepage: https://github.com/goccy/bigquery-emulator

import (
	"fmt"
	
	"os/exec"
)

func installBigqueryEmulator() {
	// Método 1: Descargar y extraer .tar.gz
	bigqueryemulator_tar_url := "https://github.com/goccy/bigquery-emulator/archive/refs/tags/v0.6.5.tar.gz"
	bigqueryemulator_cmd_tar := exec.Command("curl", "-L", bigqueryemulator_tar_url, "-o", "package.tar.gz")
	err := bigqueryemulator_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bigqueryemulator_zip_url := "https://github.com/goccy/bigquery-emulator/archive/refs/tags/v0.6.5.zip"
	bigqueryemulator_cmd_zip := exec.Command("curl", "-L", bigqueryemulator_zip_url, "-o", "package.zip")
	err = bigqueryemulator_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bigqueryemulator_bin_url := "https://github.com/goccy/bigquery-emulator/archive/refs/tags/v0.6.5.bin"
	bigqueryemulator_cmd_bin := exec.Command("curl", "-L", bigqueryemulator_bin_url, "-o", "binary.bin")
	err = bigqueryemulator_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bigqueryemulator_src_url := "https://github.com/goccy/bigquery-emulator/archive/refs/tags/v0.6.5.src.tar.gz"
	bigqueryemulator_cmd_src := exec.Command("curl", "-L", bigqueryemulator_src_url, "-o", "source.tar.gz")
	err = bigqueryemulator_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bigqueryemulator_cmd_direct := exec.Command("./binary")
	err = bigqueryemulator_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go@1.22")
exec.Command("latte", "install", "go@1.22")
}
