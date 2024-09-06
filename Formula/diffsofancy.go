package main

// DiffSoFancy - Good-lookin' diffs with diff-highlight and more
// Homepage: https://github.com/so-fancy/diff-so-fancy

import (
	"fmt"
	
	"os/exec"
)

func installDiffSoFancy() {
	// Método 1: Descargar y extraer .tar.gz
	diffsofancy_tar_url := "https://github.com/so-fancy/diff-so-fancy/archive/refs/tags/v1.4.4.tar.gz"
	diffsofancy_cmd_tar := exec.Command("curl", "-L", diffsofancy_tar_url, "-o", "package.tar.gz")
	err := diffsofancy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	diffsofancy_zip_url := "https://github.com/so-fancy/diff-so-fancy/archive/refs/tags/v1.4.4.zip"
	diffsofancy_cmd_zip := exec.Command("curl", "-L", diffsofancy_zip_url, "-o", "package.zip")
	err = diffsofancy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	diffsofancy_bin_url := "https://github.com/so-fancy/diff-so-fancy/archive/refs/tags/v1.4.4.bin"
	diffsofancy_cmd_bin := exec.Command("curl", "-L", diffsofancy_bin_url, "-o", "binary.bin")
	err = diffsofancy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	diffsofancy_src_url := "https://github.com/so-fancy/diff-so-fancy/archive/refs/tags/v1.4.4.src.tar.gz"
	diffsofancy_cmd_src := exec.Command("curl", "-L", diffsofancy_src_url, "-o", "source.tar.gz")
	err = diffsofancy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	diffsofancy_cmd_direct := exec.Command("./binary")
	err = diffsofancy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
