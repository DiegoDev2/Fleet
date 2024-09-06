package main

// Sqliteodbc - ODBC driver for SQLite
// Homepage: https://ch-werner.homepage.t-online.de/sqliteodbc/

import (
	"fmt"
	
	"os/exec"
)

func installSqliteodbc() {
	// Método 1: Descargar y extraer .tar.gz
	sqliteodbc_tar_url := "https://ch-werner.homepage.t-online.de/sqliteodbc/sqliteodbc-0.99991.tar.gz"
	sqliteodbc_cmd_tar := exec.Command("curl", "-L", sqliteodbc_tar_url, "-o", "package.tar.gz")
	err := sqliteodbc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sqliteodbc_zip_url := "https://ch-werner.homepage.t-online.de/sqliteodbc/sqliteodbc-0.99991.zip"
	sqliteodbc_cmd_zip := exec.Command("curl", "-L", sqliteodbc_zip_url, "-o", "package.zip")
	err = sqliteodbc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sqliteodbc_bin_url := "https://ch-werner.homepage.t-online.de/sqliteodbc/sqliteodbc-0.99991.bin"
	sqliteodbc_cmd_bin := exec.Command("curl", "-L", sqliteodbc_bin_url, "-o", "binary.bin")
	err = sqliteodbc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sqliteodbc_src_url := "https://ch-werner.homepage.t-online.de/sqliteodbc/sqliteodbc-0.99991.src.tar.gz"
	sqliteodbc_cmd_src := exec.Command("curl", "-L", sqliteodbc_src_url, "-o", "source.tar.gz")
	err = sqliteodbc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sqliteodbc_cmd_direct := exec.Command("./binary")
	err = sqliteodbc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: sqlite")
exec.Command("latte", "install", "sqlite")
	fmt.Println("Instalando dependencia: unixodbc")
exec.Command("latte", "install", "unixodbc")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
}
