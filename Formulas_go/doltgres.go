package main

// Doltgres - Dolt for Postgres
// Homepage: https://github.com/dolthub/doltgresql

import (
	"fmt"
	
	"os/exec"
)

func installDoltgres() {
	// Método 1: Descargar y extraer .tar.gz
	doltgres_tar_url := "https://github.com/dolthub/doltgresql/archive/refs/tags/v0.11.1.tar.gz"
	doltgres_cmd_tar := exec.Command("curl", "-L", doltgres_tar_url, "-o", "package.tar.gz")
	err := doltgres_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	doltgres_zip_url := "https://github.com/dolthub/doltgresql/archive/refs/tags/v0.11.1.zip"
	doltgres_cmd_zip := exec.Command("curl", "-L", doltgres_zip_url, "-o", "package.zip")
	err = doltgres_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	doltgres_bin_url := "https://github.com/dolthub/doltgresql/archive/refs/tags/v0.11.1.bin"
	doltgres_cmd_bin := exec.Command("curl", "-L", doltgres_bin_url, "-o", "binary.bin")
	err = doltgres_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	doltgres_src_url := "https://github.com/dolthub/doltgresql/archive/refs/tags/v0.11.1.src.tar.gz"
	doltgres_cmd_src := exec.Command("curl", "-L", doltgres_src_url, "-o", "source.tar.gz")
	err = doltgres_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	doltgres_cmd_direct := exec.Command("./binary")
	err = doltgres_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: libpq")
exec.Command("latte", "install", "libpq")
}
