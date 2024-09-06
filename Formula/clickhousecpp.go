package main

// ClickhouseCpp - C++ client library for ClickHouse
// Homepage: https://github.com/ClickHouse/clickhouse-cpp

import (
	"fmt"
	
	"os/exec"
)

func installClickhouseCpp() {
	// Método 1: Descargar y extraer .tar.gz
	clickhousecpp_tar_url := "https://github.com/ClickHouse/clickhouse-cpp/archive/refs/tags/v2.5.1.tar.gz"
	clickhousecpp_cmd_tar := exec.Command("curl", "-L", clickhousecpp_tar_url, "-o", "package.tar.gz")
	err := clickhousecpp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	clickhousecpp_zip_url := "https://github.com/ClickHouse/clickhouse-cpp/archive/refs/tags/v2.5.1.zip"
	clickhousecpp_cmd_zip := exec.Command("curl", "-L", clickhousecpp_zip_url, "-o", "package.zip")
	err = clickhousecpp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	clickhousecpp_bin_url := "https://github.com/ClickHouse/clickhouse-cpp/archive/refs/tags/v2.5.1.bin"
	clickhousecpp_cmd_bin := exec.Command("curl", "-L", clickhousecpp_bin_url, "-o", "binary.bin")
	err = clickhousecpp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	clickhousecpp_src_url := "https://github.com/ClickHouse/clickhouse-cpp/archive/refs/tags/v2.5.1.src.tar.gz"
	clickhousecpp_cmd_src := exec.Command("curl", "-L", clickhousecpp_src_url, "-o", "source.tar.gz")
	err = clickhousecpp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	clickhousecpp_cmd_direct := exec.Command("./binary")
	err = clickhousecpp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: abseil")
	exec.Command("latte", "install", "abseil").Run()
	fmt.Println("Instalando dependencia: lz4")
	exec.Command("latte", "install", "lz4").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
