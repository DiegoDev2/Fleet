package main

// Spin - Efficient verification tool of multi-threaded software
// Homepage: https://spinroot.com/spin/whatispin.html

import (
	"fmt"
	
	"os/exec"
)

func installSpin() {
	// Método 1: Descargar y extraer .tar.gz
	spin_tar_url := "https://github.com/nimble-code/Spin/archive/refs/tags/version-6.5.2.tar.gz"
	spin_cmd_tar := exec.Command("curl", "-L", spin_tar_url, "-o", "package.tar.gz")
	err := spin_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	spin_zip_url := "https://github.com/nimble-code/Spin/archive/refs/tags/version-6.5.2.zip"
	spin_cmd_zip := exec.Command("curl", "-L", spin_zip_url, "-o", "package.zip")
	err = spin_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	spin_bin_url := "https://github.com/nimble-code/Spin/archive/refs/tags/version-6.5.2.bin"
	spin_cmd_bin := exec.Command("curl", "-L", spin_bin_url, "-o", "binary.bin")
	err = spin_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	spin_src_url := "https://github.com/nimble-code/Spin/archive/refs/tags/version-6.5.2.src.tar.gz"
	spin_cmd_src := exec.Command("curl", "-L", spin_src_url, "-o", "source.tar.gz")
	err = spin_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	spin_cmd_direct := exec.Command("./binary")
	err = spin_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
