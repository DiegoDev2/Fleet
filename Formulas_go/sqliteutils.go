package main

// SqliteUtils - CLI utility for manipulating SQLite databases
// Homepage: https://sqlite-utils.datasette.io/

import (
	"fmt"
	
	"os/exec"
)

func installSqliteUtils() {
	// Método 1: Descargar y extraer .tar.gz
	sqliteutils_tar_url := "https://files.pythonhosted.org/packages/65/c5/a16a5d3f5f64e700a77de3df427ce1fcf5029e38db3352e12a0696448569/sqlite_utils-3.37.tar.gz"
	sqliteutils_cmd_tar := exec.Command("curl", "-L", sqliteutils_tar_url, "-o", "package.tar.gz")
	err := sqliteutils_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sqliteutils_zip_url := "https://files.pythonhosted.org/packages/65/c5/a16a5d3f5f64e700a77de3df427ce1fcf5029e38db3352e12a0696448569/sqlite_utils-3.37.zip"
	sqliteutils_cmd_zip := exec.Command("curl", "-L", sqliteutils_zip_url, "-o", "package.zip")
	err = sqliteutils_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sqliteutils_bin_url := "https://files.pythonhosted.org/packages/65/c5/a16a5d3f5f64e700a77de3df427ce1fcf5029e38db3352e12a0696448569/sqlite_utils-3.37.bin"
	sqliteutils_cmd_bin := exec.Command("curl", "-L", sqliteutils_bin_url, "-o", "binary.bin")
	err = sqliteutils_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sqliteutils_src_url := "https://files.pythonhosted.org/packages/65/c5/a16a5d3f5f64e700a77de3df427ce1fcf5029e38db3352e12a0696448569/sqlite_utils-3.37.src.tar.gz"
	sqliteutils_cmd_src := exec.Command("curl", "-L", sqliteutils_src_url, "-o", "source.tar.gz")
	err = sqliteutils_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sqliteutils_cmd_direct := exec.Command("./binary")
	err = sqliteutils_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
