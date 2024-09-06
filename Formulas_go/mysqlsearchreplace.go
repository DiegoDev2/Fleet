package main

// MysqlSearchReplace - Database search and replace script in PHP
// Homepage: https://interconnectit.com/products/search-and-replace-for-wordpress-databases/

import (
	"fmt"
	
	"os/exec"
)

func installMysqlSearchReplace() {
	// Método 1: Descargar y extraer .tar.gz
	mysqlsearchreplace_tar_url := "https://github.com/interconnectit/Search-Replace-DB/archive/refs/tags/4.1.2.tar.gz"
	mysqlsearchreplace_cmd_tar := exec.Command("curl", "-L", mysqlsearchreplace_tar_url, "-o", "package.tar.gz")
	err := mysqlsearchreplace_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mysqlsearchreplace_zip_url := "https://github.com/interconnectit/Search-Replace-DB/archive/refs/tags/4.1.2.zip"
	mysqlsearchreplace_cmd_zip := exec.Command("curl", "-L", mysqlsearchreplace_zip_url, "-o", "package.zip")
	err = mysqlsearchreplace_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mysqlsearchreplace_bin_url := "https://github.com/interconnectit/Search-Replace-DB/archive/refs/tags/4.1.2.bin"
	mysqlsearchreplace_cmd_bin := exec.Command("curl", "-L", mysqlsearchreplace_bin_url, "-o", "binary.bin")
	err = mysqlsearchreplace_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mysqlsearchreplace_src_url := "https://github.com/interconnectit/Search-Replace-DB/archive/refs/tags/4.1.2.src.tar.gz"
	mysqlsearchreplace_cmd_src := exec.Command("curl", "-L", mysqlsearchreplace_src_url, "-o", "source.tar.gz")
	err = mysqlsearchreplace_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mysqlsearchreplace_cmd_direct := exec.Command("./binary")
	err = mysqlsearchreplace_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
