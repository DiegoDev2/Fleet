package main

// Tdb - Trivial DataBase, by the Samba project
// Homepage: https://tdb.samba.org/

import (
	"fmt"
	
	"os/exec"
)

func installTdb() {
	// Método 1: Descargar y extraer .tar.gz
	tdb_tar_url := "https://www.samba.org/ftp/tdb/tdb-1.4.12.tar.gz"
	tdb_cmd_tar := exec.Command("curl", "-L", tdb_tar_url, "-o", "package.tar.gz")
	err := tdb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tdb_zip_url := "https://www.samba.org/ftp/tdb/tdb-1.4.12.zip"
	tdb_cmd_zip := exec.Command("curl", "-L", tdb_zip_url, "-o", "package.zip")
	err = tdb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tdb_bin_url := "https://www.samba.org/ftp/tdb/tdb-1.4.12.bin"
	tdb_cmd_bin := exec.Command("curl", "-L", tdb_bin_url, "-o", "binary.bin")
	err = tdb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tdb_src_url := "https://www.samba.org/ftp/tdb/tdb-1.4.12.src.tar.gz"
	tdb_cmd_src := exec.Command("curl", "-L", tdb_src_url, "-o", "source.tar.gz")
	err = tdb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tdb_cmd_direct := exec.Command("./binary")
	err = tdb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
