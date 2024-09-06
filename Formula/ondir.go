package main

// Ondir - Automatically execute scripts as you traverse directories
// Homepage: https://swapoff.org/ondir.html

import (
	"fmt"
	
	"os/exec"
)

func installOndir() {
	// Método 1: Descargar y extraer .tar.gz
	ondir_tar_url := "https://swapoff.org/files/ondir/ondir-0.2.4.tar.gz"
	ondir_cmd_tar := exec.Command("curl", "-L", ondir_tar_url, "-o", "package.tar.gz")
	err := ondir_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ondir_zip_url := "https://swapoff.org/files/ondir/ondir-0.2.4.zip"
	ondir_cmd_zip := exec.Command("curl", "-L", ondir_zip_url, "-o", "package.zip")
	err = ondir_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ondir_bin_url := "https://swapoff.org/files/ondir/ondir-0.2.4.bin"
	ondir_cmd_bin := exec.Command("curl", "-L", ondir_bin_url, "-o", "binary.bin")
	err = ondir_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ondir_src_url := "https://swapoff.org/files/ondir/ondir-0.2.4.src.tar.gz"
	ondir_cmd_src := exec.Command("curl", "-L", ondir_src_url, "-o", "source.tar.gz")
	err = ondir_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ondir_cmd_direct := exec.Command("./binary")
	err = ondir_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
