package main

// SqlMigrate - SQL schema migration tool for Go
// Homepage: https://github.com/rubenv/sql-migrate

import (
	"fmt"
	
	"os/exec"
)

func installSqlMigrate() {
	// Método 1: Descargar y extraer .tar.gz
	sqlmigrate_tar_url := "https://github.com/rubenv/sql-migrate/archive/refs/tags/v1.7.0.tar.gz"
	sqlmigrate_cmd_tar := exec.Command("curl", "-L", sqlmigrate_tar_url, "-o", "package.tar.gz")
	err := sqlmigrate_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sqlmigrate_zip_url := "https://github.com/rubenv/sql-migrate/archive/refs/tags/v1.7.0.zip"
	sqlmigrate_cmd_zip := exec.Command("curl", "-L", sqlmigrate_zip_url, "-o", "package.zip")
	err = sqlmigrate_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sqlmigrate_bin_url := "https://github.com/rubenv/sql-migrate/archive/refs/tags/v1.7.0.bin"
	sqlmigrate_cmd_bin := exec.Command("curl", "-L", sqlmigrate_bin_url, "-o", "binary.bin")
	err = sqlmigrate_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sqlmigrate_src_url := "https://github.com/rubenv/sql-migrate/archive/refs/tags/v1.7.0.src.tar.gz"
	sqlmigrate_cmd_src := exec.Command("curl", "-L", sqlmigrate_src_url, "-o", "source.tar.gz")
	err = sqlmigrate_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sqlmigrate_cmd_direct := exec.Command("./binary")
	err = sqlmigrate_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
