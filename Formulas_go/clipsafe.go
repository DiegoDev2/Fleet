package main

// Clipsafe - Command-line interface to Password Safe
// Homepage: https://waxandwane.org/clipsafe.html

import (
	"fmt"
	
	"os/exec"
)

func installClipsafe() {
	// Método 1: Descargar y extraer .tar.gz
	clipsafe_tar_url := "https://waxandwane.org/download/clipsafe-1.1.tar.gz"
	clipsafe_cmd_tar := exec.Command("curl", "-L", clipsafe_tar_url, "-o", "package.tar.gz")
	err := clipsafe_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	clipsafe_zip_url := "https://waxandwane.org/download/clipsafe-1.1.zip"
	clipsafe_cmd_zip := exec.Command("curl", "-L", clipsafe_zip_url, "-o", "package.zip")
	err = clipsafe_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	clipsafe_bin_url := "https://waxandwane.org/download/clipsafe-1.1.bin"
	clipsafe_cmd_bin := exec.Command("curl", "-L", clipsafe_bin_url, "-o", "binary.bin")
	err = clipsafe_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	clipsafe_src_url := "https://waxandwane.org/download/clipsafe-1.1.src.tar.gz"
	clipsafe_cmd_src := exec.Command("curl", "-L", clipsafe_src_url, "-o", "source.tar.gz")
	err = clipsafe_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	clipsafe_cmd_direct := exec.Command("./binary")
	err = clipsafe_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
