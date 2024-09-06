package main

// Sqlfmt - SQL formatter with width-aware output
// Homepage: https://sqlfum.pt/

import (
	"fmt"
	
	"os/exec"
)

func installSqlfmt() {
	// Método 1: Descargar y extraer .tar.gz
	sqlfmt_tar_url := "https://github.com/mjibson/sqlfmt/archive/refs/tags/v0.5.0.tar.gz"
	sqlfmt_cmd_tar := exec.Command("curl", "-L", sqlfmt_tar_url, "-o", "package.tar.gz")
	err := sqlfmt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sqlfmt_zip_url := "https://github.com/mjibson/sqlfmt/archive/refs/tags/v0.5.0.zip"
	sqlfmt_cmd_zip := exec.Command("curl", "-L", sqlfmt_zip_url, "-o", "package.zip")
	err = sqlfmt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sqlfmt_bin_url := "https://github.com/mjibson/sqlfmt/archive/refs/tags/v0.5.0.bin"
	sqlfmt_cmd_bin := exec.Command("curl", "-L", sqlfmt_bin_url, "-o", "binary.bin")
	err = sqlfmt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sqlfmt_src_url := "https://github.com/mjibson/sqlfmt/archive/refs/tags/v0.5.0.src.tar.gz"
	sqlfmt_cmd_src := exec.Command("curl", "-L", sqlfmt_src_url, "-o", "source.tar.gz")
	err = sqlfmt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sqlfmt_cmd_direct := exec.Command("./binary")
	err = sqlfmt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
