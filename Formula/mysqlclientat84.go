package main

// MysqlClientAT84 - Open source relational database management system
// Homepage: https://dev.mysql.com/doc/refman/8.4/en/

import (
	"fmt"
	
	"os/exec"
)

func installMysqlClientAT84() {
	// Método 1: Descargar y extraer .tar.gz
	mysqlclientat84_tar_url := "https://cdn.mysql.com/Downloads/MySQL-8.4/mysql-8.4.2.tar.gz"
	mysqlclientat84_cmd_tar := exec.Command("curl", "-L", mysqlclientat84_tar_url, "-o", "package.tar.gz")
	err := mysqlclientat84_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mysqlclientat84_zip_url := "https://cdn.mysql.com/Downloads/MySQL-8.4/mysql-8.4.2.zip"
	mysqlclientat84_cmd_zip := exec.Command("curl", "-L", mysqlclientat84_zip_url, "-o", "package.zip")
	err = mysqlclientat84_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mysqlclientat84_bin_url := "https://cdn.mysql.com/Downloads/MySQL-8.4/mysql-8.4.2.bin"
	mysqlclientat84_cmd_bin := exec.Command("curl", "-L", mysqlclientat84_bin_url, "-o", "binary.bin")
	err = mysqlclientat84_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mysqlclientat84_src_url := "https://cdn.mysql.com/Downloads/MySQL-8.4/mysql-8.4.2.src.tar.gz"
	mysqlclientat84_cmd_src := exec.Command("curl", "-L", mysqlclientat84_src_url, "-o", "source.tar.gz")
	err = mysqlclientat84_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mysqlclientat84_cmd_direct := exec.Command("./binary")
	err = mysqlclientat84_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bison")
	exec.Command("latte", "install", "bison").Run()
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libevent")
	exec.Command("latte", "install", "libevent").Run()
	fmt.Println("Instalando dependencia: libfido2")
	exec.Command("latte", "install", "libfido2").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: zlib")
	exec.Command("latte", "install", "zlib").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
}
