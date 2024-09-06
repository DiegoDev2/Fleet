package main

// Cppi - Indent C preprocessor directives to reflect their nesting
// Homepage: https://www.gnu.org/software/cppi/

import (
	"fmt"
	
	"os/exec"
)

func installCppi() {
	// Método 1: Descargar y extraer .tar.gz
	cppi_tar_url := "https://ftp.gnu.org/gnu/cppi/cppi-1.18.tar.xz"
	cppi_cmd_tar := exec.Command("curl", "-L", cppi_tar_url, "-o", "package.tar.gz")
	err := cppi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cppi_zip_url := "https://ftp.gnu.org/gnu/cppi/cppi-1.18.tar.xz"
	cppi_cmd_zip := exec.Command("curl", "-L", cppi_zip_url, "-o", "package.zip")
	err = cppi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cppi_bin_url := "https://ftp.gnu.org/gnu/cppi/cppi-1.18.tar.xz"
	cppi_cmd_bin := exec.Command("curl", "-L", cppi_bin_url, "-o", "binary.bin")
	err = cppi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cppi_src_url := "https://ftp.gnu.org/gnu/cppi/cppi-1.18.tar.xz"
	cppi_cmd_src := exec.Command("curl", "-L", cppi_src_url, "-o", "source.tar.gz")
	err = cppi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cppi_cmd_direct := exec.Command("./binary")
	err = cppi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
