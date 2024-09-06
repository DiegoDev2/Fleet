package main

// GnuWhich - GNU implementation of which utility
// Homepage: https://savannah.gnu.org/projects/which/

import (
	"fmt"
	
	"os/exec"
)

func installGnuWhich() {
	// Método 1: Descargar y extraer .tar.gz
	gnuwhich_tar_url := "https://ftp.gnu.org/gnu/which/which-2.21.tar.gz"
	gnuwhich_cmd_tar := exec.Command("curl", "-L", gnuwhich_tar_url, "-o", "package.tar.gz")
	err := gnuwhich_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gnuwhich_zip_url := "https://ftp.gnu.org/gnu/which/which-2.21.zip"
	gnuwhich_cmd_zip := exec.Command("curl", "-L", gnuwhich_zip_url, "-o", "package.zip")
	err = gnuwhich_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gnuwhich_bin_url := "https://ftp.gnu.org/gnu/which/which-2.21.bin"
	gnuwhich_cmd_bin := exec.Command("curl", "-L", gnuwhich_bin_url, "-o", "binary.bin")
	err = gnuwhich_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gnuwhich_src_url := "https://ftp.gnu.org/gnu/which/which-2.21.src.tar.gz"
	gnuwhich_cmd_src := exec.Command("curl", "-L", gnuwhich_src_url, "-o", "source.tar.gz")
	err = gnuwhich_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gnuwhich_cmd_direct := exec.Command("./binary")
	err = gnuwhich_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
