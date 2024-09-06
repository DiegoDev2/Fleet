package main

// Pth - GNU Portable THreads
// Homepage: https://www.gnu.org/software/pth/

import (
	"fmt"
	
	"os/exec"
)

func installPth() {
	// Método 1: Descargar y extraer .tar.gz
	pth_tar_url := "https://ftp.gnu.org/gnu/pth/pth-2.0.7.tar.gz"
	pth_cmd_tar := exec.Command("curl", "-L", pth_tar_url, "-o", "package.tar.gz")
	err := pth_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pth_zip_url := "https://ftp.gnu.org/gnu/pth/pth-2.0.7.zip"
	pth_cmd_zip := exec.Command("curl", "-L", pth_zip_url, "-o", "package.zip")
	err = pth_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pth_bin_url := "https://ftp.gnu.org/gnu/pth/pth-2.0.7.bin"
	pth_cmd_bin := exec.Command("curl", "-L", pth_bin_url, "-o", "binary.bin")
	err = pth_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pth_src_url := "https://ftp.gnu.org/gnu/pth/pth-2.0.7.src.tar.gz"
	pth_cmd_src := exec.Command("curl", "-L", pth_src_url, "-o", "source.tar.gz")
	err = pth_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pth_cmd_direct := exec.Command("./binary")
	err = pth_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
