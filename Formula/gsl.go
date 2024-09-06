package main

// Gsl - Numerical library for C and C++
// Homepage: https://www.gnu.org/software/gsl/

import (
	"fmt"
	
	"os/exec"
)

func installGsl() {
	// Método 1: Descargar y extraer .tar.gz
	gsl_tar_url := "https://ftp.gnu.org/gnu/gsl/gsl-2.8.tar.gz"
	gsl_cmd_tar := exec.Command("curl", "-L", gsl_tar_url, "-o", "package.tar.gz")
	err := gsl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gsl_zip_url := "https://ftp.gnu.org/gnu/gsl/gsl-2.8.zip"
	gsl_cmd_zip := exec.Command("curl", "-L", gsl_zip_url, "-o", "package.zip")
	err = gsl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gsl_bin_url := "https://ftp.gnu.org/gnu/gsl/gsl-2.8.bin"
	gsl_cmd_bin := exec.Command("curl", "-L", gsl_bin_url, "-o", "binary.bin")
	err = gsl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gsl_src_url := "https://ftp.gnu.org/gnu/gsl/gsl-2.8.src.tar.gz"
	gsl_cmd_src := exec.Command("curl", "-L", gsl_src_url, "-o", "source.tar.gz")
	err = gsl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gsl_cmd_direct := exec.Command("./binary")
	err = gsl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
