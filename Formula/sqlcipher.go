package main

// Sqlcipher - SQLite extension providing 256-bit AES encryption
// Homepage: https://www.zetetic.net/sqlcipher/

import (
	"fmt"
	
	"os/exec"
)

func installSqlcipher() {
	// Método 1: Descargar y extraer .tar.gz
	sqlcipher_tar_url := "https://github.com/sqlcipher/sqlcipher/archive/refs/tags/v4.6.1.tar.gz"
	sqlcipher_cmd_tar := exec.Command("curl", "-L", sqlcipher_tar_url, "-o", "package.tar.gz")
	err := sqlcipher_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sqlcipher_zip_url := "https://github.com/sqlcipher/sqlcipher/archive/refs/tags/v4.6.1.zip"
	sqlcipher_cmd_zip := exec.Command("curl", "-L", sqlcipher_zip_url, "-o", "package.zip")
	err = sqlcipher_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sqlcipher_bin_url := "https://github.com/sqlcipher/sqlcipher/archive/refs/tags/v4.6.1.bin"
	sqlcipher_cmd_bin := exec.Command("curl", "-L", sqlcipher_bin_url, "-o", "binary.bin")
	err = sqlcipher_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sqlcipher_src_url := "https://github.com/sqlcipher/sqlcipher/archive/refs/tags/v4.6.1.src.tar.gz"
	sqlcipher_cmd_src := exec.Command("curl", "-L", sqlcipher_src_url, "-o", "source.tar.gz")
	err = sqlcipher_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sqlcipher_cmd_direct := exec.Command("./binary")
	err = sqlcipher_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
