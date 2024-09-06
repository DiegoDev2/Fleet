package main

// Sq - Data wrangler with jq-like query language
// Homepage: https://sq.io

import (
	"fmt"
	
	"os/exec"
)

func installSq() {
	// Método 1: Descargar y extraer .tar.gz
	sq_tar_url := "https://github.com/neilotoole/sq/archive/refs/tags/v0.48.3.tar.gz"
	sq_cmd_tar := exec.Command("curl", "-L", sq_tar_url, "-o", "package.tar.gz")
	err := sq_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sq_zip_url := "https://github.com/neilotoole/sq/archive/refs/tags/v0.48.3.zip"
	sq_cmd_zip := exec.Command("curl", "-L", sq_zip_url, "-o", "package.zip")
	err = sq_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sq_bin_url := "https://github.com/neilotoole/sq/archive/refs/tags/v0.48.3.bin"
	sq_cmd_bin := exec.Command("curl", "-L", sq_bin_url, "-o", "binary.bin")
	err = sq_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sq_src_url := "https://github.com/neilotoole/sq/archive/refs/tags/v0.48.3.src.tar.gz"
	sq_cmd_src := exec.Command("curl", "-L", sq_src_url, "-o", "source.tar.gz")
	err = sq_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sq_cmd_direct := exec.Command("./binary")
	err = sq_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
