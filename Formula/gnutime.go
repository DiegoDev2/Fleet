package main

// GnuTime - GNU implementation of time utility
// Homepage: https://www.gnu.org/software/time/

import (
	"fmt"
	
	"os/exec"
)

func installGnuTime() {
	// Método 1: Descargar y extraer .tar.gz
	gnutime_tar_url := "https://ftp.gnu.org/gnu/time/time-1.9.tar.gz"
	gnutime_cmd_tar := exec.Command("curl", "-L", gnutime_tar_url, "-o", "package.tar.gz")
	err := gnutime_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gnutime_zip_url := "https://ftp.gnu.org/gnu/time/time-1.9.zip"
	gnutime_cmd_zip := exec.Command("curl", "-L", gnutime_zip_url, "-o", "package.zip")
	err = gnutime_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gnutime_bin_url := "https://ftp.gnu.org/gnu/time/time-1.9.bin"
	gnutime_cmd_bin := exec.Command("curl", "-L", gnutime_bin_url, "-o", "binary.bin")
	err = gnutime_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gnutime_src_url := "https://ftp.gnu.org/gnu/time/time-1.9.src.tar.gz"
	gnutime_cmd_src := exec.Command("curl", "-L", gnutime_src_url, "-o", "source.tar.gz")
	err = gnutime_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gnutime_cmd_direct := exec.Command("./binary")
	err = gnutime_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
