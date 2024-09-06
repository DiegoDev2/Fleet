package main

// Polyglot - Protocol adapter to run UCI engines under XBoard
// Homepage: https://www.chessprogramming.org/PolyGlot

import (
	"fmt"
	
	"os/exec"
)

func installPolyglot() {
	// Método 1: Descargar y extraer .tar.gz
	polyglot_tar_url := "http://hgm.nubati.net/releases/polyglot-2.0.4.tar.gz"
	polyglot_cmd_tar := exec.Command("curl", "-L", polyglot_tar_url, "-o", "package.tar.gz")
	err := polyglot_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	polyglot_zip_url := "http://hgm.nubati.net/releases/polyglot-2.0.4.zip"
	polyglot_cmd_zip := exec.Command("curl", "-L", polyglot_zip_url, "-o", "package.zip")
	err = polyglot_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	polyglot_bin_url := "http://hgm.nubati.net/releases/polyglot-2.0.4.bin"
	polyglot_cmd_bin := exec.Command("curl", "-L", polyglot_bin_url, "-o", "binary.bin")
	err = polyglot_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	polyglot_src_url := "http://hgm.nubati.net/releases/polyglot-2.0.4.src.tar.gz"
	polyglot_cmd_src := exec.Command("curl", "-L", polyglot_src_url, "-o", "source.tar.gz")
	err = polyglot_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	polyglot_cmd_direct := exec.Command("./binary")
	err = polyglot_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
