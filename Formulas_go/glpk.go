package main

// Glpk - Library for Linear and Mixed-Integer Programming
// Homepage: https://www.gnu.org/software/glpk/

import (
	"fmt"
	
	"os/exec"
)

func installGlpk() {
	// Método 1: Descargar y extraer .tar.gz
	glpk_tar_url := "https://ftp.gnu.org/gnu/glpk/glpk-5.0.tar.gz"
	glpk_cmd_tar := exec.Command("curl", "-L", glpk_tar_url, "-o", "package.tar.gz")
	err := glpk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	glpk_zip_url := "https://ftp.gnu.org/gnu/glpk/glpk-5.0.zip"
	glpk_cmd_zip := exec.Command("curl", "-L", glpk_zip_url, "-o", "package.zip")
	err = glpk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	glpk_bin_url := "https://ftp.gnu.org/gnu/glpk/glpk-5.0.bin"
	glpk_cmd_bin := exec.Command("curl", "-L", glpk_bin_url, "-o", "binary.bin")
	err = glpk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	glpk_src_url := "https://ftp.gnu.org/gnu/glpk/glpk-5.0.src.tar.gz"
	glpk_cmd_src := exec.Command("curl", "-L", glpk_src_url, "-o", "source.tar.gz")
	err = glpk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	glpk_cmd_direct := exec.Command("./binary")
	err = glpk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
}
