package main

// MariadbConnectorOdbc - Database driver using the industry standard ODBC API
// Homepage: https://mariadb.org/download/?tab=connector&prod=connector-odbc

import (
	"fmt"
	
	"os/exec"
)

func installMariadbConnectorOdbc() {
	// Método 1: Descargar y extraer .tar.gz
	mariadbconnectorodbc_tar_url := "https://archive.mariadb.org/connector-odbc-3.2.3/mariadb-connector-odbc-3.2.3-src.tar.gz"
	mariadbconnectorodbc_cmd_tar := exec.Command("curl", "-L", mariadbconnectorodbc_tar_url, "-o", "package.tar.gz")
	err := mariadbconnectorodbc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mariadbconnectorodbc_zip_url := "https://archive.mariadb.org/connector-odbc-3.2.3/mariadb-connector-odbc-3.2.3-src.zip"
	mariadbconnectorodbc_cmd_zip := exec.Command("curl", "-L", mariadbconnectorodbc_zip_url, "-o", "package.zip")
	err = mariadbconnectorodbc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mariadbconnectorodbc_bin_url := "https://archive.mariadb.org/connector-odbc-3.2.3/mariadb-connector-odbc-3.2.3-src.bin"
	mariadbconnectorodbc_cmd_bin := exec.Command("curl", "-L", mariadbconnectorodbc_bin_url, "-o", "binary.bin")
	err = mariadbconnectorodbc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mariadbconnectorodbc_src_url := "https://archive.mariadb.org/connector-odbc-3.2.3/mariadb-connector-odbc-3.2.3-src.src.tar.gz"
	mariadbconnectorodbc_cmd_src := exec.Command("curl", "-L", mariadbconnectorodbc_src_url, "-o", "source.tar.gz")
	err = mariadbconnectorodbc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mariadbconnectorodbc_cmd_direct := exec.Command("./binary")
	err = mariadbconnectorodbc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: mariadb-connector-c")
	exec.Command("latte", "install", "mariadb-connector-c").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: unixodbc")
	exec.Command("latte", "install", "unixodbc").Run()
}
