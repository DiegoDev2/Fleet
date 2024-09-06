package main

// Dbhash - Computes the SHA1 hash of schema and content of a SQLite database
// Homepage: https://www.sqlite.org/dbhash.html

import (
	"fmt"
	
	"os/exec"
)

func installDbhash() {
	// Método 1: Descargar y extraer .tar.gz
	dbhash_tar_url := "https://www.sqlite.org/2024/sqlite-src-3460100.zip"
	dbhash_cmd_tar := exec.Command("curl", "-L", dbhash_tar_url, "-o", "package.tar.gz")
	err := dbhash_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dbhash_zip_url := "https://www.sqlite.org/2024/sqlite-src-3460100.zip"
	dbhash_cmd_zip := exec.Command("curl", "-L", dbhash_zip_url, "-o", "package.zip")
	err = dbhash_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dbhash_bin_url := "https://www.sqlite.org/2024/sqlite-src-3460100.zip"
	dbhash_cmd_bin := exec.Command("curl", "-L", dbhash_bin_url, "-o", "binary.bin")
	err = dbhash_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dbhash_src_url := "https://www.sqlite.org/2024/sqlite-src-3460100.zip"
	dbhash_cmd_src := exec.Command("curl", "-L", dbhash_src_url, "-o", "source.tar.gz")
	err = dbhash_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dbhash_cmd_direct := exec.Command("./binary")
	err = dbhash_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
