package main

// Fzy - Fast, simple fuzzy text selector with an advanced scoring algorithm
// Homepage: https://github.com/jhawthorn/fzy

import (
	"fmt"
	
	"os/exec"
)

func installFzy() {
	// Método 1: Descargar y extraer .tar.gz
	fzy_tar_url := "https://github.com/jhawthorn/fzy/releases/download/1.0/fzy-1.0.tar.gz"
	fzy_cmd_tar := exec.Command("curl", "-L", fzy_tar_url, "-o", "package.tar.gz")
	err := fzy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fzy_zip_url := "https://github.com/jhawthorn/fzy/releases/download/1.0/fzy-1.0.zip"
	fzy_cmd_zip := exec.Command("curl", "-L", fzy_zip_url, "-o", "package.zip")
	err = fzy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fzy_bin_url := "https://github.com/jhawthorn/fzy/releases/download/1.0/fzy-1.0.bin"
	fzy_cmd_bin := exec.Command("curl", "-L", fzy_bin_url, "-o", "binary.bin")
	err = fzy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fzy_src_url := "https://github.com/jhawthorn/fzy/releases/download/1.0/fzy-1.0.src.tar.gz"
	fzy_cmd_src := exec.Command("curl", "-L", fzy_src_url, "-o", "source.tar.gz")
	err = fzy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fzy_cmd_direct := exec.Command("./binary")
	err = fzy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
