package main

// Xorriso - ISO9660+RR manipulation tool
// Homepage: https://www.gnu.org/software/xorriso/

import (
	"fmt"
	
	"os/exec"
)

func installXorriso() {
	// Método 1: Descargar y extraer .tar.gz
	xorriso_tar_url := "https://ftp.gnu.org/gnu/xorriso/xorriso-1.5.6.tar.gz"
	xorriso_cmd_tar := exec.Command("curl", "-L", xorriso_tar_url, "-o", "package.tar.gz")
	err := xorriso_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xorriso_zip_url := "https://ftp.gnu.org/gnu/xorriso/xorriso-1.5.6.zip"
	xorriso_cmd_zip := exec.Command("curl", "-L", xorriso_zip_url, "-o", "package.zip")
	err = xorriso_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xorriso_bin_url := "https://ftp.gnu.org/gnu/xorriso/xorriso-1.5.6.bin"
	xorriso_cmd_bin := exec.Command("curl", "-L", xorriso_bin_url, "-o", "binary.bin")
	err = xorriso_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xorriso_src_url := "https://ftp.gnu.org/gnu/xorriso/xorriso-1.5.6.src.tar.gz"
	xorriso_cmd_src := exec.Command("curl", "-L", xorriso_src_url, "-o", "source.tar.gz")
	err = xorriso_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xorriso_cmd_direct := exec.Command("./binary")
	err = xorriso_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
