package main

// Cpio - Copies files into or out of a cpio or tar archive
// Homepage: https://www.gnu.org/software/cpio/

import (
	"fmt"
	
	"os/exec"
)

func installCpio() {
	// Método 1: Descargar y extraer .tar.gz
	cpio_tar_url := "https://ftp.gnu.org/gnu/cpio/cpio-2.15.tar.bz2"
	cpio_cmd_tar := exec.Command("curl", "-L", cpio_tar_url, "-o", "package.tar.gz")
	err := cpio_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cpio_zip_url := "https://ftp.gnu.org/gnu/cpio/cpio-2.15.tar.bz2"
	cpio_cmd_zip := exec.Command("curl", "-L", cpio_zip_url, "-o", "package.zip")
	err = cpio_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cpio_bin_url := "https://ftp.gnu.org/gnu/cpio/cpio-2.15.tar.bz2"
	cpio_cmd_bin := exec.Command("curl", "-L", cpio_bin_url, "-o", "binary.bin")
	err = cpio_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cpio_src_url := "https://ftp.gnu.org/gnu/cpio/cpio-2.15.tar.bz2"
	cpio_cmd_src := exec.Command("curl", "-L", cpio_src_url, "-o", "source.tar.gz")
	err = cpio_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cpio_cmd_direct := exec.Command("./binary")
	err = cpio_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
