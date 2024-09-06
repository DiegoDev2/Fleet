package main

// Ghi - Work on GitHub issues on the command-line
// Homepage: https://github.com/drazisil/ghi

import (
	"fmt"
	
	"os/exec"
)

func installGhi() {
	// Método 1: Descargar y extraer .tar.gz
	ghi_tar_url := "https://github.com/drazisil/ghi/archive/refs/tags/1.2.1.tar.gz"
	ghi_cmd_tar := exec.Command("curl", "-L", ghi_tar_url, "-o", "package.tar.gz")
	err := ghi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ghi_zip_url := "https://github.com/drazisil/ghi/archive/refs/tags/1.2.1.zip"
	ghi_cmd_zip := exec.Command("curl", "-L", ghi_zip_url, "-o", "package.zip")
	err = ghi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ghi_bin_url := "https://github.com/drazisil/ghi/archive/refs/tags/1.2.1.bin"
	ghi_cmd_bin := exec.Command("curl", "-L", ghi_bin_url, "-o", "binary.bin")
	err = ghi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ghi_src_url := "https://github.com/drazisil/ghi/archive/refs/tags/1.2.1.src.tar.gz"
	ghi_cmd_src := exec.Command("curl", "-L", ghi_src_url, "-o", "source.tar.gz")
	err = ghi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ghi_cmd_direct := exec.Command("./binary")
	err = ghi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
