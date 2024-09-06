package main

// MysqlAT84 - Open source relational database management system
// Homepage: https://dev.mysql.com/doc/refman/8.4/en/

import (
	"fmt"
	
	"os/exec"
)

func installMysqlAT84() {
	// Método 1: Descargar y extraer .tar.gz
	mysqlat84_tar_url := "https://cdn.mysql.com/Downloads/MySQL-8.4/mysql-8.4.2.tar.gz"
	mysqlat84_cmd_tar := exec.Command("curl", "-L", mysqlat84_tar_url, "-o", "package.tar.gz")
	err := mysqlat84_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mysqlat84_zip_url := "https://cdn.mysql.com/Downloads/MySQL-8.4/mysql-8.4.2.zip"
	mysqlat84_cmd_zip := exec.Command("curl", "-L", mysqlat84_zip_url, "-o", "package.zip")
	err = mysqlat84_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mysqlat84_bin_url := "https://cdn.mysql.com/Downloads/MySQL-8.4/mysql-8.4.2.bin"
	mysqlat84_cmd_bin := exec.Command("curl", "-L", mysqlat84_bin_url, "-o", "binary.bin")
	err = mysqlat84_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mysqlat84_src_url := "https://cdn.mysql.com/Downloads/MySQL-8.4/mysql-8.4.2.src.tar.gz"
	mysqlat84_cmd_src := exec.Command("curl", "-L", mysqlat84_src_url, "-o", "source.tar.gz")
	err = mysqlat84_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mysqlat84_cmd_direct := exec.Command("./binary")
	err = mysqlat84_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: abseil")
	exec.Command("latte", "install", "abseil").Run()
	fmt.Println("Instalando dependencia: icu4c")
	exec.Command("latte", "install", "icu4c").Run()
	fmt.Println("Instalando dependencia: libfido2")
	exec.Command("latte", "install", "libfido2").Run()
	fmt.Println("Instalando dependencia: lz4")
	exec.Command("latte", "install", "lz4").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: protobuf")
	exec.Command("latte", "install", "protobuf").Run()
	fmt.Println("Instalando dependencia: zlib")
	exec.Command("latte", "install", "zlib").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
	fmt.Println("Instalando dependencia: llvm")
	exec.Command("latte", "install", "llvm").Run()
	fmt.Println("Instalando dependencia: patchelf")
	exec.Command("latte", "install", "patchelf").Run()
	fmt.Println("Instalando dependencia: libtirpc")
	exec.Command("latte", "install", "libtirpc").Run()
}
