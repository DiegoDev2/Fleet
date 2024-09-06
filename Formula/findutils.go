package main

// Findutils - Collection of GNU find, xargs, and locate
// Homepage: https://www.gnu.org/software/findutils/

import (
	"fmt"
	
	"os/exec"
)

func installFindutils() {
	// Método 1: Descargar y extraer .tar.gz
	findutils_tar_url := "https://ftp.gnu.org/gnu/findutils/findutils-4.10.0.tar.xz"
	findutils_cmd_tar := exec.Command("curl", "-L", findutils_tar_url, "-o", "package.tar.gz")
	err := findutils_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	findutils_zip_url := "https://ftp.gnu.org/gnu/findutils/findutils-4.10.0.tar.xz"
	findutils_cmd_zip := exec.Command("curl", "-L", findutils_zip_url, "-o", "package.zip")
	err = findutils_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	findutils_bin_url := "https://ftp.gnu.org/gnu/findutils/findutils-4.10.0.tar.xz"
	findutils_cmd_bin := exec.Command("curl", "-L", findutils_bin_url, "-o", "binary.bin")
	err = findutils_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	findutils_src_url := "https://ftp.gnu.org/gnu/findutils/findutils-4.10.0.tar.xz"
	findutils_cmd_src := exec.Command("curl", "-L", findutils_src_url, "-o", "source.tar.gz")
	err = findutils_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	findutils_cmd_direct := exec.Command("./binary")
	err = findutils_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
