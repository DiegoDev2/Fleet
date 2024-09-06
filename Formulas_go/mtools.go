package main

// Mtools - Tools for manipulating MSDOS files
// Homepage: https://www.gnu.org/software/mtools/

import (
	"fmt"
	
	"os/exec"
)

func installMtools() {
	// Método 1: Descargar y extraer .tar.gz
	mtools_tar_url := "https://ftp.gnu.org/gnu/mtools/mtools-4.0.44.tar.gz"
	mtools_cmd_tar := exec.Command("curl", "-L", mtools_tar_url, "-o", "package.tar.gz")
	err := mtools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mtools_zip_url := "https://ftp.gnu.org/gnu/mtools/mtools-4.0.44.zip"
	mtools_cmd_zip := exec.Command("curl", "-L", mtools_zip_url, "-o", "package.zip")
	err = mtools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mtools_bin_url := "https://ftp.gnu.org/gnu/mtools/mtools-4.0.44.bin"
	mtools_cmd_bin := exec.Command("curl", "-L", mtools_bin_url, "-o", "binary.bin")
	err = mtools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mtools_src_url := "https://ftp.gnu.org/gnu/mtools/mtools-4.0.44.src.tar.gz"
	mtools_cmd_src := exec.Command("curl", "-L", mtools_src_url, "-o", "source.tar.gz")
	err = mtools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mtools_cmd_direct := exec.Command("./binary")
	err = mtools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
