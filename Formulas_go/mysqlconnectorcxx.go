package main

// MysqlConnectorCxx - MySQL database connector for C++ applications
// Homepage: https://dev.mysql.com/downloads/connector/cpp/

import (
	"fmt"
	
	"os/exec"
)

func installMysqlConnectorCxx() {
	// Método 1: Descargar y extraer .tar.gz
	mysqlconnectorcxx_tar_url := "https://dev.mysql.com/get/Downloads/Connector-C++/mysql-connector-c++-9.0.0-src.tar.gz"
	mysqlconnectorcxx_cmd_tar := exec.Command("curl", "-L", mysqlconnectorcxx_tar_url, "-o", "package.tar.gz")
	err := mysqlconnectorcxx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mysqlconnectorcxx_zip_url := "https://dev.mysql.com/get/Downloads/Connector-C++/mysql-connector-c++-9.0.0-src.zip"
	mysqlconnectorcxx_cmd_zip := exec.Command("curl", "-L", mysqlconnectorcxx_zip_url, "-o", "package.zip")
	err = mysqlconnectorcxx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mysqlconnectorcxx_bin_url := "https://dev.mysql.com/get/Downloads/Connector-C++/mysql-connector-c++-9.0.0-src.bin"
	mysqlconnectorcxx_cmd_bin := exec.Command("curl", "-L", mysqlconnectorcxx_bin_url, "-o", "binary.bin")
	err = mysqlconnectorcxx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mysqlconnectorcxx_src_url := "https://dev.mysql.com/get/Downloads/Connector-C++/mysql-connector-c++-9.0.0-src.src.tar.gz"
	mysqlconnectorcxx_cmd_src := exec.Command("curl", "-L", mysqlconnectorcxx_src_url, "-o", "source.tar.gz")
	err = mysqlconnectorcxx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mysqlconnectorcxx_cmd_direct := exec.Command("./binary")
	err = mysqlconnectorcxx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: rapidjson")
exec.Command("latte", "install", "rapidjson")
	fmt.Println("Instalando dependencia: lz4")
exec.Command("latte", "install", "lz4")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: protobuf@21")
exec.Command("latte", "install", "protobuf@21")
	fmt.Println("Instalando dependencia: zlib")
exec.Command("latte", "install", "zlib")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
}
