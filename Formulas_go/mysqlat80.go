package main

// MysqlAT80 - Open source relational database management system
// Homepage: https://dev.mysql.com/doc/refman/8.0/en/

import (
	"fmt"
	
	"os/exec"
)

func installMysqlAT80() {
	// Método 1: Descargar y extraer .tar.gz
	mysqlat80_tar_url := "https://cdn.mysql.com/Downloads/MySQL-8.0/mysql-boost-8.0.39.tar.gz"
	mysqlat80_cmd_tar := exec.Command("curl", "-L", mysqlat80_tar_url, "-o", "package.tar.gz")
	err := mysqlat80_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mysqlat80_zip_url := "https://cdn.mysql.com/Downloads/MySQL-8.0/mysql-boost-8.0.39.zip"
	mysqlat80_cmd_zip := exec.Command("curl", "-L", mysqlat80_zip_url, "-o", "package.zip")
	err = mysqlat80_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mysqlat80_bin_url := "https://cdn.mysql.com/Downloads/MySQL-8.0/mysql-boost-8.0.39.bin"
	mysqlat80_cmd_bin := exec.Command("curl", "-L", mysqlat80_bin_url, "-o", "binary.bin")
	err = mysqlat80_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mysqlat80_src_url := "https://cdn.mysql.com/Downloads/MySQL-8.0/mysql-boost-8.0.39.src.tar.gz"
	mysqlat80_cmd_src := exec.Command("curl", "-L", mysqlat80_src_url, "-o", "source.tar.gz")
	err = mysqlat80_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mysqlat80_cmd_direct := exec.Command("./binary")
	err = mysqlat80_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: abseil")
exec.Command("latte", "install", "abseil")
	fmt.Println("Instalando dependencia: icu4c")
exec.Command("latte", "install", "icu4c")
	fmt.Println("Instalando dependencia: libevent")
exec.Command("latte", "install", "libevent")
	fmt.Println("Instalando dependencia: libfido2")
exec.Command("latte", "install", "libfido2")
	fmt.Println("Instalando dependencia: lz4")
exec.Command("latte", "install", "lz4")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: protobuf")
exec.Command("latte", "install", "protobuf")
	fmt.Println("Instalando dependencia: zlib")
exec.Command("latte", "install", "zlib")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
	fmt.Println("Instalando dependencia: patchelf")
exec.Command("latte", "install", "patchelf")
	fmt.Println("Instalando dependencia: libtirpc")
exec.Command("latte", "install", "libtirpc")
}
