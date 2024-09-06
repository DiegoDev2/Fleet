package main

// BisonAT27 - Parser generator
// Homepage: https://www.gnu.org/software/bison/

import (
	"fmt"
	
	"os/exec"
)

func installBisonAT27() {
	// Método 1: Descargar y extraer .tar.gz
	bisonat27_tar_url := "https://ftp.gnu.org/gnu/bison/bison-2.7.1.tar.gz"
	bisonat27_cmd_tar := exec.Command("curl", "-L", bisonat27_tar_url, "-o", "package.tar.gz")
	err := bisonat27_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bisonat27_zip_url := "https://ftp.gnu.org/gnu/bison/bison-2.7.1.zip"
	bisonat27_cmd_zip := exec.Command("curl", "-L", bisonat27_zip_url, "-o", "package.zip")
	err = bisonat27_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bisonat27_bin_url := "https://ftp.gnu.org/gnu/bison/bison-2.7.1.bin"
	bisonat27_cmd_bin := exec.Command("curl", "-L", bisonat27_bin_url, "-o", "binary.bin")
	err = bisonat27_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bisonat27_src_url := "https://ftp.gnu.org/gnu/bison/bison-2.7.1.src.tar.gz"
	bisonat27_cmd_src := exec.Command("curl", "-L", bisonat27_src_url, "-o", "source.tar.gz")
	err = bisonat27_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bisonat27_cmd_direct := exec.Command("./binary")
	err = bisonat27_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
