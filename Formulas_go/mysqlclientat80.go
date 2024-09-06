package main

// MysqlClientAT80 - Open source relational database management system
// Homepage: https://dev.mysql.com/doc/refman/8.0/en/

import (
	"fmt"
	
	"os/exec"
)

func installMysqlClientAT80() {
	// Método 1: Descargar y extraer .tar.gz
	mysqlclientat80_tar_url := "https://cdn.mysql.com/Downloads/MySQL-8.0/mysql-boost-8.0.39.tar.gz"
	mysqlclientat80_cmd_tar := exec.Command("curl", "-L", mysqlclientat80_tar_url, "-o", "package.tar.gz")
	err := mysqlclientat80_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mysqlclientat80_zip_url := "https://cdn.mysql.com/Downloads/MySQL-8.0/mysql-boost-8.0.39.zip"
	mysqlclientat80_cmd_zip := exec.Command("curl", "-L", mysqlclientat80_zip_url, "-o", "package.zip")
	err = mysqlclientat80_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mysqlclientat80_bin_url := "https://cdn.mysql.com/Downloads/MySQL-8.0/mysql-boost-8.0.39.bin"
	mysqlclientat80_cmd_bin := exec.Command("curl", "-L", mysqlclientat80_bin_url, "-o", "binary.bin")
	err = mysqlclientat80_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mysqlclientat80_src_url := "https://cdn.mysql.com/Downloads/MySQL-8.0/mysql-boost-8.0.39.src.tar.gz"
	mysqlclientat80_cmd_src := exec.Command("curl", "-L", mysqlclientat80_src_url, "-o", "source.tar.gz")
	err = mysqlclientat80_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mysqlclientat80_cmd_direct := exec.Command("./binary")
	err = mysqlclientat80_cmd_direct.Run()
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
