package main

// BerkeleyDbAT4 - High performance key/value database
// Homepage: https://www.oracle.com/database/technologies/related/berkeleydb.html

import (
	"fmt"
	
	"os/exec"
)

func installBerkeleyDbAT4() {
	// Método 1: Descargar y extraer .tar.gz
	berkeleydbat4_tar_url := "https://download.oracle.com/berkeley-db/db-4.8.30.tar.gz"
	berkeleydbat4_cmd_tar := exec.Command("curl", "-L", berkeleydbat4_tar_url, "-o", "package.tar.gz")
	err := berkeleydbat4_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	berkeleydbat4_zip_url := "https://download.oracle.com/berkeley-db/db-4.8.30.zip"
	berkeleydbat4_cmd_zip := exec.Command("curl", "-L", berkeleydbat4_zip_url, "-o", "package.zip")
	err = berkeleydbat4_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	berkeleydbat4_bin_url := "https://download.oracle.com/berkeley-db/db-4.8.30.bin"
	berkeleydbat4_cmd_bin := exec.Command("curl", "-L", berkeleydbat4_bin_url, "-o", "binary.bin")
	err = berkeleydbat4_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	berkeleydbat4_src_url := "https://download.oracle.com/berkeley-db/db-4.8.30.src.tar.gz"
	berkeleydbat4_cmd_src := exec.Command("curl", "-L", berkeleydbat4_src_url, "-o", "source.tar.gz")
	err = berkeleydbat4_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	berkeleydbat4_cmd_direct := exec.Command("./binary")
	err = berkeleydbat4_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
