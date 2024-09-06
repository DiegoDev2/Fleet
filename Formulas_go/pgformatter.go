package main

// Pgformatter - PostgreSQL syntax beautifier
// Homepage: https://sqlformat.darold.net/

import (
	"fmt"
	
	"os/exec"
)

func installPgformatter() {
	// Método 1: Descargar y extraer .tar.gz
	pgformatter_tar_url := "https://github.com/darold/pgFormatter/archive/refs/tags/v5.5.tar.gz"
	pgformatter_cmd_tar := exec.Command("curl", "-L", pgformatter_tar_url, "-o", "package.tar.gz")
	err := pgformatter_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pgformatter_zip_url := "https://github.com/darold/pgFormatter/archive/refs/tags/v5.5.zip"
	pgformatter_cmd_zip := exec.Command("curl", "-L", pgformatter_zip_url, "-o", "package.zip")
	err = pgformatter_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pgformatter_bin_url := "https://github.com/darold/pgFormatter/archive/refs/tags/v5.5.bin"
	pgformatter_cmd_bin := exec.Command("curl", "-L", pgformatter_bin_url, "-o", "binary.bin")
	err = pgformatter_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pgformatter_src_url := "https://github.com/darold/pgFormatter/archive/refs/tags/v5.5.src.tar.gz"
	pgformatter_cmd_src := exec.Command("curl", "-L", pgformatter_src_url, "-o", "source.tar.gz")
	err = pgformatter_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pgformatter_cmd_direct := exec.Command("./binary")
	err = pgformatter_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
