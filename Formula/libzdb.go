package main

// Libzdb - Database connection pool library
// Homepage: https://tildeslash.com/libzdb/

import (
	"fmt"
	
	"os/exec"
)

func installLibzdb() {
	// Método 1: Descargar y extraer .tar.gz
	libzdb_tar_url := "https://tildeslash.com/libzdb/dist/libzdb-3.2.3.tar.gz"
	libzdb_cmd_tar := exec.Command("curl", "-L", libzdb_tar_url, "-o", "package.tar.gz")
	err := libzdb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libzdb_zip_url := "https://tildeslash.com/libzdb/dist/libzdb-3.2.3.zip"
	libzdb_cmd_zip := exec.Command("curl", "-L", libzdb_zip_url, "-o", "package.zip")
	err = libzdb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libzdb_bin_url := "https://tildeslash.com/libzdb/dist/libzdb-3.2.3.bin"
	libzdb_cmd_bin := exec.Command("curl", "-L", libzdb_bin_url, "-o", "binary.bin")
	err = libzdb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libzdb_src_url := "https://tildeslash.com/libzdb/dist/libzdb-3.2.3.src.tar.gz"
	libzdb_cmd_src := exec.Command("curl", "-L", libzdb_src_url, "-o", "source.tar.gz")
	err = libzdb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libzdb_cmd_direct := exec.Command("./binary")
	err = libzdb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libpq")
	exec.Command("latte", "install", "libpq").Run()
	fmt.Println("Instalando dependencia: mysql-client")
	exec.Command("latte", "install", "mysql-client").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: sqlite")
	exec.Command("latte", "install", "sqlite").Run()
}
