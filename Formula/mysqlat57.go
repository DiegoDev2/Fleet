package main

// MysqlAT57 - Open source relational database management system
// Homepage: https://dev.mysql.com/doc/refman/5.7/en/

import (
	"fmt"
	
	"os/exec"
)

func installMysqlAT57() {
	// Método 1: Descargar y extraer .tar.gz
	mysqlat57_tar_url := "https://cdn.mysql.com/Downloads/MySQL-5.7/mysql-boost-5.7.44.tar.gz"
	mysqlat57_cmd_tar := exec.Command("curl", "-L", mysqlat57_tar_url, "-o", "package.tar.gz")
	err := mysqlat57_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mysqlat57_zip_url := "https://cdn.mysql.com/Downloads/MySQL-5.7/mysql-boost-5.7.44.zip"
	mysqlat57_cmd_zip := exec.Command("curl", "-L", mysqlat57_zip_url, "-o", "package.zip")
	err = mysqlat57_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mysqlat57_bin_url := "https://cdn.mysql.com/Downloads/MySQL-5.7/mysql-boost-5.7.44.bin"
	mysqlat57_cmd_bin := exec.Command("curl", "-L", mysqlat57_bin_url, "-o", "binary.bin")
	err = mysqlat57_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mysqlat57_src_url := "https://cdn.mysql.com/Downloads/MySQL-5.7/mysql-boost-5.7.44.src.tar.gz"
	mysqlat57_cmd_src := exec.Command("curl", "-L", mysqlat57_src_url, "-o", "source.tar.gz")
	err = mysqlat57_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mysqlat57_cmd_direct := exec.Command("./binary")
	err = mysqlat57_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: libevent")
	exec.Command("latte", "install", "libevent").Run()
	fmt.Println("Instalando dependencia: lz4")
	exec.Command("latte", "install", "lz4").Run()
	fmt.Println("Instalando dependencia: openssl@1.1")
	exec.Command("latte", "install", "openssl@1.1").Run()
	fmt.Println("Instalando dependencia: protobuf")
	exec.Command("latte", "install", "protobuf").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libtirpc")
	exec.Command("latte", "install", "libtirpc").Run()
}
