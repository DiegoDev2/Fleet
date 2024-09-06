package main

// Cdb - Create and read constant databases
// Homepage: https://cr.yp.to/cdb.html

import (
	"fmt"
	
	"os/exec"
)

func installCdb() {
	// Método 1: Descargar y extraer .tar.gz
	cdb_tar_url := "https://cr.yp.to/cdb/cdb-0.75.tar.gz"
	cdb_cmd_tar := exec.Command("curl", "-L", cdb_tar_url, "-o", "package.tar.gz")
	err := cdb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cdb_zip_url := "https://cr.yp.to/cdb/cdb-0.75.zip"
	cdb_cmd_zip := exec.Command("curl", "-L", cdb_zip_url, "-o", "package.zip")
	err = cdb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cdb_bin_url := "https://cr.yp.to/cdb/cdb-0.75.bin"
	cdb_cmd_bin := exec.Command("curl", "-L", cdb_bin_url, "-o", "binary.bin")
	err = cdb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cdb_src_url := "https://cr.yp.to/cdb/cdb-0.75.src.tar.gz"
	cdb_cmd_src := exec.Command("curl", "-L", cdb_src_url, "-o", "source.tar.gz")
	err = cdb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cdb_cmd_direct := exec.Command("./binary")
	err = cdb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
