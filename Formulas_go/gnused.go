package main

// GnuSed - GNU implementation of the famous stream editor
// Homepage: https://www.gnu.org/software/sed/

import (
	"fmt"
	
	"os/exec"
)

func installGnuSed() {
	// Método 1: Descargar y extraer .tar.gz
	gnused_tar_url := "https://ftp.gnu.org/gnu/sed/sed-4.9.tar.xz"
	gnused_cmd_tar := exec.Command("curl", "-L", gnused_tar_url, "-o", "package.tar.gz")
	err := gnused_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gnused_zip_url := "https://ftp.gnu.org/gnu/sed/sed-4.9.tar.xz"
	gnused_cmd_zip := exec.Command("curl", "-L", gnused_zip_url, "-o", "package.zip")
	err = gnused_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gnused_bin_url := "https://ftp.gnu.org/gnu/sed/sed-4.9.tar.xz"
	gnused_cmd_bin := exec.Command("curl", "-L", gnused_bin_url, "-o", "binary.bin")
	err = gnused_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gnused_src_url := "https://ftp.gnu.org/gnu/sed/sed-4.9.tar.xz"
	gnused_cmd_src := exec.Command("curl", "-L", gnused_src_url, "-o", "source.tar.gz")
	err = gnused_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gnused_cmd_direct := exec.Command("./binary")
	err = gnused_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
