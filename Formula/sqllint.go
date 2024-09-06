package main

// SqlLint - SQL linter to do sanity checks on your queries and bring errors back from the DB
// Homepage: https://github.com/joereynolds/sql-lint

import (
	"fmt"
	
	"os/exec"
)

func installSqlLint() {
	// Método 1: Descargar y extraer .tar.gz
	sqllint_tar_url := "https://registry.npmjs.org/sql-lint/-/sql-lint-1.0.0.tgz"
	sqllint_cmd_tar := exec.Command("curl", "-L", sqllint_tar_url, "-o", "package.tar.gz")
	err := sqllint_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sqllint_zip_url := "https://registry.npmjs.org/sql-lint/-/sql-lint-1.0.0.tgz"
	sqllint_cmd_zip := exec.Command("curl", "-L", sqllint_zip_url, "-o", "package.zip")
	err = sqllint_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sqllint_bin_url := "https://registry.npmjs.org/sql-lint/-/sql-lint-1.0.0.tgz"
	sqllint_cmd_bin := exec.Command("curl", "-L", sqllint_bin_url, "-o", "binary.bin")
	err = sqllint_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sqllint_src_url := "https://registry.npmjs.org/sql-lint/-/sql-lint-1.0.0.tgz"
	sqllint_cmd_src := exec.Command("curl", "-L", sqllint_src_url, "-o", "source.tar.gz")
	err = sqllint_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sqllint_cmd_direct := exec.Command("./binary")
	err = sqllint_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
}
