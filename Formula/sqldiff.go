package main

// Sqldiff - Displays the differences between SQLite databases
// Homepage: https://www.sqlite.org/sqldiff.html

import (
	"fmt"
	
	"os/exec"
)

func installSqldiff() {
	// Método 1: Descargar y extraer .tar.gz
	sqldiff_tar_url := "https://www.sqlite.org/2024/sqlite-src-3460100.zip"
	sqldiff_cmd_tar := exec.Command("curl", "-L", sqldiff_tar_url, "-o", "package.tar.gz")
	err := sqldiff_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sqldiff_zip_url := "https://www.sqlite.org/2024/sqlite-src-3460100.zip"
	sqldiff_cmd_zip := exec.Command("curl", "-L", sqldiff_zip_url, "-o", "package.zip")
	err = sqldiff_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sqldiff_bin_url := "https://www.sqlite.org/2024/sqlite-src-3460100.zip"
	sqldiff_cmd_bin := exec.Command("curl", "-L", sqldiff_bin_url, "-o", "binary.bin")
	err = sqldiff_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sqldiff_src_url := "https://www.sqlite.org/2024/sqlite-src-3460100.zip"
	sqldiff_cmd_src := exec.Command("curl", "-L", sqldiff_src_url, "-o", "source.tar.gz")
	err = sqldiff_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sqldiff_cmd_direct := exec.Command("./binary")
	err = sqldiff_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
