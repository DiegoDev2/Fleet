package main

// Gforth - Implementation of the ANS Forth language
// Homepage: https://www.gnu.org/software/gforth/

import (
	"fmt"
	
	"os/exec"
)

func installGforth() {
	// Método 1: Descargar y extraer .tar.gz
	gforth_tar_url := "https://ftp.gnu.org/gnu/gforth/gforth-0.7.3.tar.gz"
	gforth_cmd_tar := exec.Command("curl", "-L", gforth_tar_url, "-o", "package.tar.gz")
	err := gforth_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gforth_zip_url := "https://ftp.gnu.org/gnu/gforth/gforth-0.7.3.zip"
	gforth_cmd_zip := exec.Command("curl", "-L", gforth_zip_url, "-o", "package.zip")
	err = gforth_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gforth_bin_url := "https://ftp.gnu.org/gnu/gforth/gforth-0.7.3.bin"
	gforth_cmd_bin := exec.Command("curl", "-L", gforth_bin_url, "-o", "binary.bin")
	err = gforth_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gforth_src_url := "https://ftp.gnu.org/gnu/gforth/gforth-0.7.3.src.tar.gz"
	gforth_cmd_src := exec.Command("curl", "-L", gforth_src_url, "-o", "source.tar.gz")
	err = gforth_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gforth_cmd_direct := exec.Command("./binary")
	err = gforth_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: emacs")
	exec.Command("latte", "install", "emacs").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
}
