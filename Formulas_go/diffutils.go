package main

// Diffutils - File comparison utilities
// Homepage: https://www.gnu.org/software/diffutils/

import (
	"fmt"
	
	"os/exec"
)

func installDiffutils() {
	// Método 1: Descargar y extraer .tar.gz
	diffutils_tar_url := "https://ftp.gnu.org/gnu/diffutils/diffutils-3.10.tar.xz"
	diffutils_cmd_tar := exec.Command("curl", "-L", diffutils_tar_url, "-o", "package.tar.gz")
	err := diffutils_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	diffutils_zip_url := "https://ftp.gnu.org/gnu/diffutils/diffutils-3.10.tar.xz"
	diffutils_cmd_zip := exec.Command("curl", "-L", diffutils_zip_url, "-o", "package.zip")
	err = diffutils_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	diffutils_bin_url := "https://ftp.gnu.org/gnu/diffutils/diffutils-3.10.tar.xz"
	diffutils_cmd_bin := exec.Command("curl", "-L", diffutils_bin_url, "-o", "binary.bin")
	err = diffutils_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	diffutils_src_url := "https://ftp.gnu.org/gnu/diffutils/diffutils-3.10.tar.xz"
	diffutils_cmd_src := exec.Command("curl", "-L", diffutils_src_url, "-o", "source.tar.gz")
	err = diffutils_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	diffutils_cmd_direct := exec.Command("./binary")
	err = diffutils_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
