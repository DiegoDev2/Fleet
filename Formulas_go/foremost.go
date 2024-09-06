package main

// Foremost - Console program to recover files based on their headers and footers
// Homepage: https://foremost.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installForemost() {
	// Método 1: Descargar y extraer .tar.gz
	foremost_tar_url := "https://foremost.sourceforge.net/pkg/foremost-1.5.7.tar.gz"
	foremost_cmd_tar := exec.Command("curl", "-L", foremost_tar_url, "-o", "package.tar.gz")
	err := foremost_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	foremost_zip_url := "https://foremost.sourceforge.net/pkg/foremost-1.5.7.zip"
	foremost_cmd_zip := exec.Command("curl", "-L", foremost_zip_url, "-o", "package.zip")
	err = foremost_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	foremost_bin_url := "https://foremost.sourceforge.net/pkg/foremost-1.5.7.bin"
	foremost_cmd_bin := exec.Command("curl", "-L", foremost_bin_url, "-o", "binary.bin")
	err = foremost_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	foremost_src_url := "https://foremost.sourceforge.net/pkg/foremost-1.5.7.src.tar.gz"
	foremost_cmd_src := exec.Command("curl", "-L", foremost_src_url, "-o", "source.tar.gz")
	err = foremost_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	foremost_cmd_direct := exec.Command("./binary")
	err = foremost_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
