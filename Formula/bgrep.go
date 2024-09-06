package main

// Bgrep - Like grep but for binary strings
// Homepage: https://github.com/tmbinc/bgrep

import (
	"fmt"
	
	"os/exec"
)

func installBgrep() {
	// Método 1: Descargar y extraer .tar.gz
	bgrep_tar_url := "https://github.com/tmbinc/bgrep/archive/refs/tags/bgrep-0.2.tar.gz"
	bgrep_cmd_tar := exec.Command("curl", "-L", bgrep_tar_url, "-o", "package.tar.gz")
	err := bgrep_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bgrep_zip_url := "https://github.com/tmbinc/bgrep/archive/refs/tags/bgrep-0.2.zip"
	bgrep_cmd_zip := exec.Command("curl", "-L", bgrep_zip_url, "-o", "package.zip")
	err = bgrep_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bgrep_bin_url := "https://github.com/tmbinc/bgrep/archive/refs/tags/bgrep-0.2.bin"
	bgrep_cmd_bin := exec.Command("curl", "-L", bgrep_bin_url, "-o", "binary.bin")
	err = bgrep_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bgrep_src_url := "https://github.com/tmbinc/bgrep/archive/refs/tags/bgrep-0.2.src.tar.gz"
	bgrep_cmd_src := exec.Command("curl", "-L", bgrep_src_url, "-o", "source.tar.gz")
	err = bgrep_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bgrep_cmd_direct := exec.Command("./binary")
	err = bgrep_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
