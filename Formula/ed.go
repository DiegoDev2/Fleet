package main

// Ed - Classic UNIX line editor
// Homepage: https://www.gnu.org/software/ed/ed.html

import (
	"fmt"
	
	"os/exec"
)

func installEd() {
	// Método 1: Descargar y extraer .tar.gz
	ed_tar_url := "https://ftp.gnu.org/gnu/ed/ed-1.20.2.tar.lz"
	ed_cmd_tar := exec.Command("curl", "-L", ed_tar_url, "-o", "package.tar.gz")
	err := ed_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ed_zip_url := "https://ftp.gnu.org/gnu/ed/ed-1.20.2.tar.lz"
	ed_cmd_zip := exec.Command("curl", "-L", ed_zip_url, "-o", "package.zip")
	err = ed_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ed_bin_url := "https://ftp.gnu.org/gnu/ed/ed-1.20.2.tar.lz"
	ed_cmd_bin := exec.Command("curl", "-L", ed_bin_url, "-o", "binary.bin")
	err = ed_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ed_src_url := "https://ftp.gnu.org/gnu/ed/ed-1.20.2.tar.lz"
	ed_cmd_src := exec.Command("curl", "-L", ed_src_url, "-o", "source.tar.gz")
	err = ed_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ed_cmd_direct := exec.Command("./binary")
	err = ed_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
