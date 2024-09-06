package main

// BerkeleyDb - High performance key/value database
// Homepage: https://www.oracle.com/database/technologies/related/berkeleydb.html

import (
	"fmt"
	
	"os/exec"
)

func installBerkeleyDb() {
	// Método 1: Descargar y extraer .tar.gz
	berkeleydb_tar_url := "https://download.oracle.com/berkeley-db/db-18.1.40.tar.gz"
	berkeleydb_cmd_tar := exec.Command("curl", "-L", berkeleydb_tar_url, "-o", "package.tar.gz")
	err := berkeleydb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	berkeleydb_zip_url := "https://download.oracle.com/berkeley-db/db-18.1.40.zip"
	berkeleydb_cmd_zip := exec.Command("curl", "-L", berkeleydb_zip_url, "-o", "package.zip")
	err = berkeleydb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	berkeleydb_bin_url := "https://download.oracle.com/berkeley-db/db-18.1.40.bin"
	berkeleydb_cmd_bin := exec.Command("curl", "-L", berkeleydb_bin_url, "-o", "binary.bin")
	err = berkeleydb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	berkeleydb_src_url := "https://download.oracle.com/berkeley-db/db-18.1.40.src.tar.gz"
	berkeleydb_cmd_src := exec.Command("curl", "-L", berkeleydb_src_url, "-o", "source.tar.gz")
	err = berkeleydb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	berkeleydb_cmd_direct := exec.Command("./binary")
	err = berkeleydb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
