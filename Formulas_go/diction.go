package main

// Diction - GNU diction and style
// Homepage: https://www.gnu.org/software/diction/

import (
	"fmt"
	
	"os/exec"
)

func installDiction() {
	// Método 1: Descargar y extraer .tar.gz
	diction_tar_url := "https://ftp.gnu.org/gnu/diction/diction-1.11.tar.gz"
	diction_cmd_tar := exec.Command("curl", "-L", diction_tar_url, "-o", "package.tar.gz")
	err := diction_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	diction_zip_url := "https://ftp.gnu.org/gnu/diction/diction-1.11.zip"
	diction_cmd_zip := exec.Command("curl", "-L", diction_zip_url, "-o", "package.zip")
	err = diction_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	diction_bin_url := "https://ftp.gnu.org/gnu/diction/diction-1.11.bin"
	diction_cmd_bin := exec.Command("curl", "-L", diction_bin_url, "-o", "binary.bin")
	err = diction_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	diction_src_url := "https://ftp.gnu.org/gnu/diction/diction-1.11.src.tar.gz"
	diction_cmd_src := exec.Command("curl", "-L", diction_src_url, "-o", "source.tar.gz")
	err = diction_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	diction_cmd_direct := exec.Command("./binary")
	err = diction_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
