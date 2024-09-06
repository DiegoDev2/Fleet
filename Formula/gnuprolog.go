package main

// GnuProlog - Prolog compiler with constraint solving
// Homepage: http://www.gprolog.org/

import (
	"fmt"
	
	"os/exec"
)

func installGnuProlog() {
	// Método 1: Descargar y extraer .tar.gz
	gnuprolog_tar_url := "http://www.gprolog.org/gprolog-1.5.0.tar.gz"
	gnuprolog_cmd_tar := exec.Command("curl", "-L", gnuprolog_tar_url, "-o", "package.tar.gz")
	err := gnuprolog_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gnuprolog_zip_url := "http://www.gprolog.org/gprolog-1.5.0.zip"
	gnuprolog_cmd_zip := exec.Command("curl", "-L", gnuprolog_zip_url, "-o", "package.zip")
	err = gnuprolog_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gnuprolog_bin_url := "http://www.gprolog.org/gprolog-1.5.0.bin"
	gnuprolog_cmd_bin := exec.Command("curl", "-L", gnuprolog_bin_url, "-o", "binary.bin")
	err = gnuprolog_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gnuprolog_src_url := "http://www.gprolog.org/gprolog-1.5.0.src.tar.gz"
	gnuprolog_cmd_src := exec.Command("curl", "-L", gnuprolog_src_url, "-o", "source.tar.gz")
	err = gnuprolog_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gnuprolog_cmd_direct := exec.Command("./binary")
	err = gnuprolog_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
