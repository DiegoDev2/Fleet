package main

// Pgbadger - Log analyzer for PostgreSQL
// Homepage: https://pgbadger.darold.net/

import (
	"fmt"
	
	"os/exec"
)

func installPgbadger() {
	// Método 1: Descargar y extraer .tar.gz
	pgbadger_tar_url := "https://github.com/darold/pgbadger/archive/refs/tags/v12.4.tar.gz"
	pgbadger_cmd_tar := exec.Command("curl", "-L", pgbadger_tar_url, "-o", "package.tar.gz")
	err := pgbadger_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pgbadger_zip_url := "https://github.com/darold/pgbadger/archive/refs/tags/v12.4.zip"
	pgbadger_cmd_zip := exec.Command("curl", "-L", pgbadger_zip_url, "-o", "package.zip")
	err = pgbadger_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pgbadger_bin_url := "https://github.com/darold/pgbadger/archive/refs/tags/v12.4.bin"
	pgbadger_cmd_bin := exec.Command("curl", "-L", pgbadger_bin_url, "-o", "binary.bin")
	err = pgbadger_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pgbadger_src_url := "https://github.com/darold/pgbadger/archive/refs/tags/v12.4.src.tar.gz"
	pgbadger_cmd_src := exec.Command("curl", "-L", pgbadger_src_url, "-o", "source.tar.gz")
	err = pgbadger_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pgbadger_cmd_direct := exec.Command("./binary")
	err = pgbadger_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
