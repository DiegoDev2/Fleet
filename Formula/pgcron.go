package main

// PgCron - Run periodic jobs in PostgreSQL
// Homepage: https://github.com/citusdata/pg_cron

import (
	"fmt"
	
	"os/exec"
)

func installPgCron() {
	// Método 1: Descargar y extraer .tar.gz
	pgcron_tar_url := "https://github.com/citusdata/pg_cron/archive/refs/tags/v1.6.4.tar.gz"
	pgcron_cmd_tar := exec.Command("curl", "-L", pgcron_tar_url, "-o", "package.tar.gz")
	err := pgcron_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pgcron_zip_url := "https://github.com/citusdata/pg_cron/archive/refs/tags/v1.6.4.zip"
	pgcron_cmd_zip := exec.Command("curl", "-L", pgcron_zip_url, "-o", "package.zip")
	err = pgcron_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pgcron_bin_url := "https://github.com/citusdata/pg_cron/archive/refs/tags/v1.6.4.bin"
	pgcron_cmd_bin := exec.Command("curl", "-L", pgcron_bin_url, "-o", "binary.bin")
	err = pgcron_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pgcron_src_url := "https://github.com/citusdata/pg_cron/archive/refs/tags/v1.6.4.src.tar.gz"
	pgcron_cmd_src := exec.Command("curl", "-L", pgcron_src_url, "-o", "source.tar.gz")
	err = pgcron_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pgcron_cmd_direct := exec.Command("./binary")
	err = pgcron_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: postgresql@14")
	exec.Command("latte", "install", "postgresql@14").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
