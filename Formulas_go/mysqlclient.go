package main

// MysqlClient - Open source relational database management system
// Homepage: https://dev.mysql.com/doc/refman/9.0/en/

import (
	"fmt"
	
	"os/exec"
)

func installMysqlClient() {
	// Método 1: Descargar y extraer .tar.gz
	mysqlclient_tar_url := "https://cdn.mysql.com/Downloads/MySQL-9.0/mysql-9.0.1.tar.gz"
	mysqlclient_cmd_tar := exec.Command("curl", "-L", mysqlclient_tar_url, "-o", "package.tar.gz")
	err := mysqlclient_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mysqlclient_zip_url := "https://cdn.mysql.com/Downloads/MySQL-9.0/mysql-9.0.1.zip"
	mysqlclient_cmd_zip := exec.Command("curl", "-L", mysqlclient_zip_url, "-o", "package.zip")
	err = mysqlclient_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mysqlclient_bin_url := "https://cdn.mysql.com/Downloads/MySQL-9.0/mysql-9.0.1.bin"
	mysqlclient_cmd_bin := exec.Command("curl", "-L", mysqlclient_bin_url, "-o", "binary.bin")
	err = mysqlclient_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mysqlclient_src_url := "https://cdn.mysql.com/Downloads/MySQL-9.0/mysql-9.0.1.src.tar.gz"
	mysqlclient_cmd_src := exec.Command("curl", "-L", mysqlclient_src_url, "-o", "source.tar.gz")
	err = mysqlclient_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mysqlclient_cmd_direct := exec.Command("./binary")
	err = mysqlclient_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bison")
exec.Command("latte", "install", "bison")
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libevent")
exec.Command("latte", "install", "libevent")
	fmt.Println("Instalando dependencia: libfido2")
exec.Command("latte", "install", "libfido2")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: zlib")
exec.Command("latte", "install", "zlib")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
}
