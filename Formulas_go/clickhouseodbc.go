package main

// ClickhouseOdbc - Official ODBC driver implementation for accessing ClickHouse as a data source
// Homepage: https://github.com/ClickHouse/clickhouse-odbc

import (
	"fmt"
	
	"os/exec"
)

func installClickhouseOdbc() {
	// Método 1: Descargar y extraer .tar.gz
	clickhouseodbc_tar_url := "https://github.com/ClickHouse/clickhouse-odbc/archive/refs/tags/v1.2.1.20220905.tar.gz"
	clickhouseodbc_cmd_tar := exec.Command("curl", "-L", clickhouseodbc_tar_url, "-o", "package.tar.gz")
	err := clickhouseodbc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	clickhouseodbc_zip_url := "https://github.com/ClickHouse/clickhouse-odbc/archive/refs/tags/v1.2.1.20220905.zip"
	clickhouseodbc_cmd_zip := exec.Command("curl", "-L", clickhouseodbc_zip_url, "-o", "package.zip")
	err = clickhouseodbc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	clickhouseodbc_bin_url := "https://github.com/ClickHouse/clickhouse-odbc/archive/refs/tags/v1.2.1.20220905.bin"
	clickhouseodbc_cmd_bin := exec.Command("curl", "-L", clickhouseodbc_bin_url, "-o", "binary.bin")
	err = clickhouseodbc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	clickhouseodbc_src_url := "https://github.com/ClickHouse/clickhouse-odbc/archive/refs/tags/v1.2.1.20220905.src.tar.gz"
	clickhouseodbc_cmd_src := exec.Command("curl", "-L", clickhouseodbc_src_url, "-o", "source.tar.gz")
	err = clickhouseodbc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	clickhouseodbc_cmd_direct := exec.Command("./binary")
	err = clickhouseodbc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: folly")
exec.Command("latte", "install", "folly")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: icu4c")
exec.Command("latte", "install", "icu4c")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: poco")
exec.Command("latte", "install", "poco")
	fmt.Println("Instalando dependencia: libiodbc")
exec.Command("latte", "install", "libiodbc")
	fmt.Println("Instalando dependencia: pcre2")
exec.Command("latte", "install", "pcre2")
	fmt.Println("Instalando dependencia: unixodbc")
exec.Command("latte", "install", "unixodbc")
}
