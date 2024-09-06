package main

// Pgcopydb - Copy a Postgres database to a target Postgres server
// Homepage: https://github.com/dimitri/pgcopydb

import (
	"fmt"
	
	"os/exec"
)

func installPgcopydb() {
	// Método 1: Descargar y extraer .tar.gz
	pgcopydb_tar_url := "https://github.com/dimitri/pgcopydb/archive/refs/tags/v0.17.tar.gz"
	pgcopydb_cmd_tar := exec.Command("curl", "-L", pgcopydb_tar_url, "-o", "package.tar.gz")
	err := pgcopydb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pgcopydb_zip_url := "https://github.com/dimitri/pgcopydb/archive/refs/tags/v0.17.zip"
	pgcopydb_cmd_zip := exec.Command("curl", "-L", pgcopydb_zip_url, "-o", "package.zip")
	err = pgcopydb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pgcopydb_bin_url := "https://github.com/dimitri/pgcopydb/archive/refs/tags/v0.17.bin"
	pgcopydb_cmd_bin := exec.Command("curl", "-L", pgcopydb_bin_url, "-o", "binary.bin")
	err = pgcopydb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pgcopydb_src_url := "https://github.com/dimitri/pgcopydb/archive/refs/tags/v0.17.src.tar.gz"
	pgcopydb_cmd_src := exec.Command("curl", "-L", pgcopydb_src_url, "-o", "source.tar.gz")
	err = pgcopydb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pgcopydb_cmd_direct := exec.Command("./binary")
	err = pgcopydb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: sphinx-doc")
exec.Command("latte", "install", "sphinx-doc")
	fmt.Println("Instalando dependencia: bdw-gc")
exec.Command("latte", "install", "bdw-gc")
	fmt.Println("Instalando dependencia: libpq")
exec.Command("latte", "install", "libpq")
}
