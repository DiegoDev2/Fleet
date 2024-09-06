package main

// Manticoresearch - Open source text search engine
// Homepage: https://manticoresearch.com

import (
	"fmt"
	
	"os/exec"
)

func installManticoresearch() {
	// Método 1: Descargar y extraer .tar.gz
	manticoresearch_tar_url := "https://github.com/manticoresoftware/manticoresearch/archive/refs/tags/6.2.12.tar.gz"
	manticoresearch_cmd_tar := exec.Command("curl", "-L", manticoresearch_tar_url, "-o", "package.tar.gz")
	err := manticoresearch_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	manticoresearch_zip_url := "https://github.com/manticoresoftware/manticoresearch/archive/refs/tags/6.2.12.zip"
	manticoresearch_cmd_zip := exec.Command("curl", "-L", manticoresearch_zip_url, "-o", "package.zip")
	err = manticoresearch_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	manticoresearch_bin_url := "https://github.com/manticoresoftware/manticoresearch/archive/refs/tags/6.2.12.bin"
	manticoresearch_cmd_bin := exec.Command("curl", "-L", manticoresearch_bin_url, "-o", "binary.bin")
	err = manticoresearch_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	manticoresearch_src_url := "https://github.com/manticoresoftware/manticoresearch/archive/refs/tags/6.2.12.src.tar.gz"
	manticoresearch_cmd_src := exec.Command("curl", "-L", manticoresearch_src_url, "-o", "source.tar.gz")
	err = manticoresearch_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	manticoresearch_cmd_direct := exec.Command("./binary")
	err = manticoresearch_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: icu4c")
exec.Command("latte", "install", "icu4c")
	fmt.Println("Instalando dependencia: libpq")
exec.Command("latte", "install", "libpq")
	fmt.Println("Instalando dependencia: mysql-client@8.0")
exec.Command("latte", "install", "mysql-client@8.0")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: unixodbc")
exec.Command("latte", "install", "unixodbc")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
}
