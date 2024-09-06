package main

// Rethinkdb - Open-source database for the realtime web
// Homepage: https://rethinkdb.com/

import (
	"fmt"
	
	"os/exec"
)

func installRethinkdb() {
	// Método 1: Descargar y extraer .tar.gz
	rethinkdb_tar_url := "https://download.rethinkdb.com/repository/raw/dist/rethinkdb-2.4.4.tgz"
	rethinkdb_cmd_tar := exec.Command("curl", "-L", rethinkdb_tar_url, "-o", "package.tar.gz")
	err := rethinkdb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rethinkdb_zip_url := "https://download.rethinkdb.com/repository/raw/dist/rethinkdb-2.4.4.tgz"
	rethinkdb_cmd_zip := exec.Command("curl", "-L", rethinkdb_zip_url, "-o", "package.zip")
	err = rethinkdb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rethinkdb_bin_url := "https://download.rethinkdb.com/repository/raw/dist/rethinkdb-2.4.4.tgz"
	rethinkdb_cmd_bin := exec.Command("curl", "-L", rethinkdb_bin_url, "-o", "binary.bin")
	err = rethinkdb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rethinkdb_src_url := "https://download.rethinkdb.com/repository/raw/dist/rethinkdb-2.4.4.tgz"
	rethinkdb_cmd_src := exec.Command("curl", "-L", rethinkdb_src_url, "-o", "source.tar.gz")
	err = rethinkdb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rethinkdb_cmd_direct := exec.Command("./binary")
	err = rethinkdb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: protobuf@21")
	exec.Command("latte", "install", "protobuf@21").Run()
}
