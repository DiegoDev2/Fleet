package main

// Gpatch - Apply a diff file to an original
// Homepage: https://savannah.gnu.org/projects/patch/

import (
	"fmt"
	
	"os/exec"
)

func installGpatch() {
	// Método 1: Descargar y extraer .tar.gz
	gpatch_tar_url := "https://ftp.gnu.org/gnu/patch/patch-2.7.6.tar.xz"
	gpatch_cmd_tar := exec.Command("curl", "-L", gpatch_tar_url, "-o", "package.tar.gz")
	err := gpatch_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gpatch_zip_url := "https://ftp.gnu.org/gnu/patch/patch-2.7.6.tar.xz"
	gpatch_cmd_zip := exec.Command("curl", "-L", gpatch_zip_url, "-o", "package.zip")
	err = gpatch_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gpatch_bin_url := "https://ftp.gnu.org/gnu/patch/patch-2.7.6.tar.xz"
	gpatch_cmd_bin := exec.Command("curl", "-L", gpatch_bin_url, "-o", "binary.bin")
	err = gpatch_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gpatch_src_url := "https://ftp.gnu.org/gnu/patch/patch-2.7.6.tar.xz"
	gpatch_cmd_src := exec.Command("curl", "-L", gpatch_src_url, "-o", "source.tar.gz")
	err = gpatch_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gpatch_cmd_direct := exec.Command("./binary")
	err = gpatch_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
