package main

// Lastz - Pairwise aligner for DNA sequences
// Homepage: https://lastz.github.io/lastz/

import (
	"fmt"
	
	"os/exec"
)

func installLastz() {
	// Método 1: Descargar y extraer .tar.gz
	lastz_tar_url := "https://github.com/lastz/lastz/archive/refs/tags/1.04.22.tar.gz"
	lastz_cmd_tar := exec.Command("curl", "-L", lastz_tar_url, "-o", "package.tar.gz")
	err := lastz_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lastz_zip_url := "https://github.com/lastz/lastz/archive/refs/tags/1.04.22.zip"
	lastz_cmd_zip := exec.Command("curl", "-L", lastz_zip_url, "-o", "package.zip")
	err = lastz_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lastz_bin_url := "https://github.com/lastz/lastz/archive/refs/tags/1.04.22.bin"
	lastz_cmd_bin := exec.Command("curl", "-L", lastz_bin_url, "-o", "binary.bin")
	err = lastz_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lastz_src_url := "https://github.com/lastz/lastz/archive/refs/tags/1.04.22.src.tar.gz"
	lastz_cmd_src := exec.Command("curl", "-L", lastz_src_url, "-o", "source.tar.gz")
	err = lastz_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lastz_cmd_direct := exec.Command("./binary")
	err = lastz_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
