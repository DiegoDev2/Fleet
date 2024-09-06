package main

// GnuShogi - Japanese Chess
// Homepage: https://www.gnu.org/software/gnushogi/

import (
	"fmt"
	
	"os/exec"
)

func installGnuShogi() {
	// Método 1: Descargar y extraer .tar.gz
	gnushogi_tar_url := "https://ftp.gnu.org/gnu/gnushogi/gnushogi-1.4.2.tar.gz"
	gnushogi_cmd_tar := exec.Command("curl", "-L", gnushogi_tar_url, "-o", "package.tar.gz")
	err := gnushogi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gnushogi_zip_url := "https://ftp.gnu.org/gnu/gnushogi/gnushogi-1.4.2.zip"
	gnushogi_cmd_zip := exec.Command("curl", "-L", gnushogi_zip_url, "-o", "package.zip")
	err = gnushogi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gnushogi_bin_url := "https://ftp.gnu.org/gnu/gnushogi/gnushogi-1.4.2.bin"
	gnushogi_cmd_bin := exec.Command("curl", "-L", gnushogi_bin_url, "-o", "binary.bin")
	err = gnushogi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gnushogi_src_url := "https://ftp.gnu.org/gnu/gnushogi/gnushogi-1.4.2.src.tar.gz"
	gnushogi_cmd_src := exec.Command("curl", "-L", gnushogi_src_url, "-o", "source.tar.gz")
	err = gnushogi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gnushogi_cmd_direct := exec.Command("./binary")
	err = gnushogi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
