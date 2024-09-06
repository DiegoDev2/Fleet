package main

// Idutils - ID database and query tools
// Homepage: https://www.gnu.org/software/idutils/

import (
	"fmt"
	
	"os/exec"
)

func installIdutils() {
	// Método 1: Descargar y extraer .tar.gz
	idutils_tar_url := "https://ftp.gnu.org/gnu/idutils/idutils-4.6.tar.xz"
	idutils_cmd_tar := exec.Command("curl", "-L", idutils_tar_url, "-o", "package.tar.gz")
	err := idutils_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	idutils_zip_url := "https://ftp.gnu.org/gnu/idutils/idutils-4.6.tar.xz"
	idutils_cmd_zip := exec.Command("curl", "-L", idutils_zip_url, "-o", "package.zip")
	err = idutils_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	idutils_bin_url := "https://ftp.gnu.org/gnu/idutils/idutils-4.6.tar.xz"
	idutils_cmd_bin := exec.Command("curl", "-L", idutils_bin_url, "-o", "binary.bin")
	err = idutils_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	idutils_src_url := "https://ftp.gnu.org/gnu/idutils/idutils-4.6.tar.xz"
	idutils_cmd_src := exec.Command("curl", "-L", idutils_src_url, "-o", "source.tar.gz")
	err = idutils_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	idutils_cmd_direct := exec.Command("./binary")
	err = idutils_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
