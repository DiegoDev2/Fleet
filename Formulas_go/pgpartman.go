package main

// PgPartman - Partition management extension for PostgreSQL
// Homepage: https://github.com/pgpartman/pg_partman

import (
	"fmt"
	
	"os/exec"
)

func installPgPartman() {
	// Método 1: Descargar y extraer .tar.gz
	pgpartman_tar_url := "https://github.com/pgpartman/pg_partman/archive/refs/tags/v5.1.0.tar.gz"
	pgpartman_cmd_tar := exec.Command("curl", "-L", pgpartman_tar_url, "-o", "package.tar.gz")
	err := pgpartman_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pgpartman_zip_url := "https://github.com/pgpartman/pg_partman/archive/refs/tags/v5.1.0.zip"
	pgpartman_cmd_zip := exec.Command("curl", "-L", pgpartman_zip_url, "-o", "package.zip")
	err = pgpartman_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pgpartman_bin_url := "https://github.com/pgpartman/pg_partman/archive/refs/tags/v5.1.0.bin"
	pgpartman_cmd_bin := exec.Command("curl", "-L", pgpartman_bin_url, "-o", "binary.bin")
	err = pgpartman_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pgpartman_src_url := "https://github.com/pgpartman/pg_partman/archive/refs/tags/v5.1.0.src.tar.gz"
	pgpartman_cmd_src := exec.Command("curl", "-L", pgpartman_src_url, "-o", "source.tar.gz")
	err = pgpartman_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pgpartman_cmd_direct := exec.Command("./binary")
	err = pgpartman_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: postgresql@14")
exec.Command("latte", "install", "postgresql@14")
}
