package main

// ParquetCli - Apache Parquet command-line tools and utilities
// Homepage: https://parquet.apache.org/

import (
	"fmt"
	
	"os/exec"
)

func installParquetCli() {
	// Método 1: Descargar y extraer .tar.gz
	parquetcli_tar_url := "https://github.com/apache/parquet-java/archive/refs/tags/apache-parquet-1.14.2.tar.gz"
	parquetcli_cmd_tar := exec.Command("curl", "-L", parquetcli_tar_url, "-o", "package.tar.gz")
	err := parquetcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	parquetcli_zip_url := "https://github.com/apache/parquet-java/archive/refs/tags/apache-parquet-1.14.2.zip"
	parquetcli_cmd_zip := exec.Command("curl", "-L", parquetcli_zip_url, "-o", "package.zip")
	err = parquetcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	parquetcli_bin_url := "https://github.com/apache/parquet-java/archive/refs/tags/apache-parquet-1.14.2.bin"
	parquetcli_cmd_bin := exec.Command("curl", "-L", parquetcli_bin_url, "-o", "binary.bin")
	err = parquetcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	parquetcli_src_url := "https://github.com/apache/parquet-java/archive/refs/tags/apache-parquet-1.14.2.src.tar.gz"
	parquetcli_cmd_src := exec.Command("curl", "-L", parquetcli_src_url, "-o", "source.tar.gz")
	err = parquetcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	parquetcli_cmd_direct := exec.Command("./binary")
	err = parquetcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: maven")
	exec.Command("latte", "install", "maven").Run()
	fmt.Println("Instalando dependencia: openjdk@21")
	exec.Command("latte", "install", "openjdk@21").Run()
}
