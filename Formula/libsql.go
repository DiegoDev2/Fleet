package main

// Libsql - Fork of SQLite that is both Open Source, and Open Contributions
// Homepage: https://turso.tech/libsql

import (
	"fmt"
	
	"os/exec"
)

func installLibsql() {
	// Método 1: Descargar y extraer .tar.gz
	libsql_tar_url := "https://github.com/tursodatabase/libsql/releases/download/libsql-server-v0.24.23/source.tar.gz"
	libsql_cmd_tar := exec.Command("curl", "-L", libsql_tar_url, "-o", "package.tar.gz")
	err := libsql_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libsql_zip_url := "https://github.com/tursodatabase/libsql/releases/download/libsql-server-v0.24.23/source.zip"
	libsql_cmd_zip := exec.Command("curl", "-L", libsql_zip_url, "-o", "package.zip")
	err = libsql_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libsql_bin_url := "https://github.com/tursodatabase/libsql/releases/download/libsql-server-v0.24.23/source.bin"
	libsql_cmd_bin := exec.Command("curl", "-L", libsql_bin_url, "-o", "binary.bin")
	err = libsql_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libsql_src_url := "https://github.com/tursodatabase/libsql/releases/download/libsql-server-v0.24.23/source.src.tar.gz"
	libsql_cmd_src := exec.Command("curl", "-L", libsql_src_url, "-o", "source.tar.gz")
	err = libsql_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libsql_cmd_direct := exec.Command("./binary")
	err = libsql_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
