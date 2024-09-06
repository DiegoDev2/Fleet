package main

// Mysql - Open source relational database management system
// Homepage: https://dev.mysql.com/doc/refman/9.0/en/

import (
	"fmt"
	
	"os/exec"
)

func installMysql() {
	// Método 1: Descargar y extraer .tar.gz
	mysql_tar_url := "https://cdn.mysql.com/Downloads/MySQL-9.0/mysql-9.0.1.tar.gz"
	mysql_cmd_tar := exec.Command("curl", "-L", mysql_tar_url, "-o", "package.tar.gz")
	err := mysql_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mysql_zip_url := "https://cdn.mysql.com/Downloads/MySQL-9.0/mysql-9.0.1.zip"
	mysql_cmd_zip := exec.Command("curl", "-L", mysql_zip_url, "-o", "package.zip")
	err = mysql_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mysql_bin_url := "https://cdn.mysql.com/Downloads/MySQL-9.0/mysql-9.0.1.bin"
	mysql_cmd_bin := exec.Command("curl", "-L", mysql_bin_url, "-o", "binary.bin")
	err = mysql_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mysql_src_url := "https://cdn.mysql.com/Downloads/MySQL-9.0/mysql-9.0.1.src.tar.gz"
	mysql_cmd_src := exec.Command("curl", "-L", mysql_src_url, "-o", "source.tar.gz")
	err = mysql_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mysql_cmd_direct := exec.Command("./binary")
	err = mysql_cmd_direct.Run()
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
