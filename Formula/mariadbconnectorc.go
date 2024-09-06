package main

// MariadbConnectorC - MariaDB database connector for C applications
// Homepage: https://mariadb.org/download/?tab=connector&prod=connector-c

import (
	"fmt"
	
	"os/exec"
)

func installMariadbConnectorC() {
	// Método 1: Descargar y extraer .tar.gz
	mariadbconnectorc_tar_url := "https://archive.mariadb.org/connector-c-3.4.1/mariadb-connector-c-3.4.1-src.tar.gz"
	mariadbconnectorc_cmd_tar := exec.Command("curl", "-L", mariadbconnectorc_tar_url, "-o", "package.tar.gz")
	err := mariadbconnectorc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mariadbconnectorc_zip_url := "https://archive.mariadb.org/connector-c-3.4.1/mariadb-connector-c-3.4.1-src.zip"
	mariadbconnectorc_cmd_zip := exec.Command("curl", "-L", mariadbconnectorc_zip_url, "-o", "package.zip")
	err = mariadbconnectorc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mariadbconnectorc_bin_url := "https://archive.mariadb.org/connector-c-3.4.1/mariadb-connector-c-3.4.1-src.bin"
	mariadbconnectorc_cmd_bin := exec.Command("curl", "-L", mariadbconnectorc_bin_url, "-o", "binary.bin")
	err = mariadbconnectorc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mariadbconnectorc_src_url := "https://archive.mariadb.org/connector-c-3.4.1/mariadb-connector-c-3.4.1-src.src.tar.gz"
	mariadbconnectorc_cmd_src := exec.Command("curl", "-L", mariadbconnectorc_src_url, "-o", "source.tar.gz")
	err = mariadbconnectorc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mariadbconnectorc_cmd_direct := exec.Command("./binary")
	err = mariadbconnectorc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
}
