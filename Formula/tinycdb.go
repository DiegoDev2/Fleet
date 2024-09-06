package main

// Tinycdb - Create and read constant databases
// Homepage: https://www.corpit.ru/mjt/tinycdb.html

import (
	"fmt"
	
	"os/exec"
)

func installTinycdb() {
	// Método 1: Descargar y extraer .tar.gz
	tinycdb_tar_url := "https://www.corpit.ru/mjt/tinycdb/tinycdb-0.81.tar.gz"
	tinycdb_cmd_tar := exec.Command("curl", "-L", tinycdb_tar_url, "-o", "package.tar.gz")
	err := tinycdb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tinycdb_zip_url := "https://www.corpit.ru/mjt/tinycdb/tinycdb-0.81.zip"
	tinycdb_cmd_zip := exec.Command("curl", "-L", tinycdb_zip_url, "-o", "package.zip")
	err = tinycdb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tinycdb_bin_url := "https://www.corpit.ru/mjt/tinycdb/tinycdb-0.81.bin"
	tinycdb_cmd_bin := exec.Command("curl", "-L", tinycdb_bin_url, "-o", "binary.bin")
	err = tinycdb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tinycdb_src_url := "https://www.corpit.ru/mjt/tinycdb/tinycdb-0.81.src.tar.gz"
	tinycdb_cmd_src := exec.Command("curl", "-L", tinycdb_src_url, "-o", "source.tar.gz")
	err = tinycdb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tinycdb_cmd_direct := exec.Command("./binary")
	err = tinycdb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
