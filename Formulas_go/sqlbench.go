package main

// Sqlbench - Measures and compares the execution time of one or more SQL queries
// Homepage: https://github.com/felixge/sqlbench

import (
	"fmt"
	
	"os/exec"
)

func installSqlbench() {
	// Método 1: Descargar y extraer .tar.gz
	sqlbench_tar_url := "https://github.com/felixge/sqlbench/archive/refs/tags/v1.1.0.tar.gz"
	sqlbench_cmd_tar := exec.Command("curl", "-L", sqlbench_tar_url, "-o", "package.tar.gz")
	err := sqlbench_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sqlbench_zip_url := "https://github.com/felixge/sqlbench/archive/refs/tags/v1.1.0.zip"
	sqlbench_cmd_zip := exec.Command("curl", "-L", sqlbench_zip_url, "-o", "package.zip")
	err = sqlbench_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sqlbench_bin_url := "https://github.com/felixge/sqlbench/archive/refs/tags/v1.1.0.bin"
	sqlbench_cmd_bin := exec.Command("curl", "-L", sqlbench_bin_url, "-o", "binary.bin")
	err = sqlbench_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sqlbench_src_url := "https://github.com/felixge/sqlbench/archive/refs/tags/v1.1.0.src.tar.gz"
	sqlbench_cmd_src := exec.Command("curl", "-L", sqlbench_src_url, "-o", "source.tar.gz")
	err = sqlbench_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sqlbench_cmd_direct := exec.Command("./binary")
	err = sqlbench_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
