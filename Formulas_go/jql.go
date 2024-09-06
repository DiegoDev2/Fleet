package main

// Jql - JSON query language CLI tool
// Homepage: https://github.com/yamafaktory/jql

import (
	"fmt"
	
	"os/exec"
)

func installJql() {
	// Método 1: Descargar y extraer .tar.gz
	jql_tar_url := "https://github.com/yamafaktory/jql/archive/refs/tags/jql-v7.1.13.tar.gz"
	jql_cmd_tar := exec.Command("curl", "-L", jql_tar_url, "-o", "package.tar.gz")
	err := jql_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jql_zip_url := "https://github.com/yamafaktory/jql/archive/refs/tags/jql-v7.1.13.zip"
	jql_cmd_zip := exec.Command("curl", "-L", jql_zip_url, "-o", "package.zip")
	err = jql_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jql_bin_url := "https://github.com/yamafaktory/jql/archive/refs/tags/jql-v7.1.13.bin"
	jql_cmd_bin := exec.Command("curl", "-L", jql_bin_url, "-o", "binary.bin")
	err = jql_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jql_src_url := "https://github.com/yamafaktory/jql/archive/refs/tags/jql-v7.1.13.src.tar.gz"
	jql_cmd_src := exec.Command("curl", "-L", jql_src_url, "-o", "source.tar.gz")
	err = jql_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jql_cmd_direct := exec.Command("./binary")
	err = jql_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
