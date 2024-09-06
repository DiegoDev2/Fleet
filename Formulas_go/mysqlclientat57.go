package main

// MysqlClientAT57 - Open source relational database management system
// Homepage: https://dev.mysql.com/doc/refman/5.7/en/

import (
	"fmt"
	
	"os/exec"
)

func installMysqlClientAT57() {
	// Método 1: Descargar y extraer .tar.gz
	mysqlclientat57_tar_url := "https://dev.mysql.com/get/Downloads/MySQL-5.7/mysql-boost-5.7.42.tar.gz"
	mysqlclientat57_cmd_tar := exec.Command("curl", "-L", mysqlclientat57_tar_url, "-o", "package.tar.gz")
	err := mysqlclientat57_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mysqlclientat57_zip_url := "https://dev.mysql.com/get/Downloads/MySQL-5.7/mysql-boost-5.7.42.zip"
	mysqlclientat57_cmd_zip := exec.Command("curl", "-L", mysqlclientat57_zip_url, "-o", "package.zip")
	err = mysqlclientat57_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mysqlclientat57_bin_url := "https://dev.mysql.com/get/Downloads/MySQL-5.7/mysql-boost-5.7.42.bin"
	mysqlclientat57_cmd_bin := exec.Command("curl", "-L", mysqlclientat57_bin_url, "-o", "binary.bin")
	err = mysqlclientat57_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mysqlclientat57_src_url := "https://dev.mysql.com/get/Downloads/MySQL-5.7/mysql-boost-5.7.42.src.tar.gz"
	mysqlclientat57_cmd_src := exec.Command("curl", "-L", mysqlclientat57_src_url, "-o", "source.tar.gz")
	err = mysqlclientat57_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mysqlclientat57_cmd_direct := exec.Command("./binary")
	err = mysqlclientat57_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: openssl@1.1")
exec.Command("latte", "install", "openssl@1.1")
}
