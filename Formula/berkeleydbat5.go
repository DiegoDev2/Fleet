package main

// BerkeleyDbAT5 - High performance key/value database
// Homepage: https://www.oracle.com/database/technologies/related/berkeleydb.html

import (
	"fmt"
	
	"os/exec"
)

func installBerkeleyDbAT5() {
	// Método 1: Descargar y extraer .tar.gz
	berkeleydbat5_tar_url := "https://download.oracle.com/berkeley-db/db-5.3.28.tar.gz"
	berkeleydbat5_cmd_tar := exec.Command("curl", "-L", berkeleydbat5_tar_url, "-o", "package.tar.gz")
	err := berkeleydbat5_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	berkeleydbat5_zip_url := "https://download.oracle.com/berkeley-db/db-5.3.28.zip"
	berkeleydbat5_cmd_zip := exec.Command("curl", "-L", berkeleydbat5_zip_url, "-o", "package.zip")
	err = berkeleydbat5_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	berkeleydbat5_bin_url := "https://download.oracle.com/berkeley-db/db-5.3.28.bin"
	berkeleydbat5_cmd_bin := exec.Command("curl", "-L", berkeleydbat5_bin_url, "-o", "binary.bin")
	err = berkeleydbat5_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	berkeleydbat5_src_url := "https://download.oracle.com/berkeley-db/db-5.3.28.src.tar.gz"
	berkeleydbat5_cmd_src := exec.Command("curl", "-L", berkeleydbat5_src_url, "-o", "source.tar.gz")
	err = berkeleydbat5_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	berkeleydbat5_cmd_direct := exec.Command("./binary")
	err = berkeleydbat5_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
