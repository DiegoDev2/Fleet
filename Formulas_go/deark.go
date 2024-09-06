package main

// Deark - File conversion utility for older formats
// Homepage: https://entropymine.com/deark/

import (
	"fmt"
	
	"os/exec"
)

func installDeark() {
	// Método 1: Descargar y extraer .tar.gz
	deark_tar_url := "https://entropymine.com/deark/releases/deark-1.6.8.tar.gz"
	deark_cmd_tar := exec.Command("curl", "-L", deark_tar_url, "-o", "package.tar.gz")
	err := deark_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	deark_zip_url := "https://entropymine.com/deark/releases/deark-1.6.8.zip"
	deark_cmd_zip := exec.Command("curl", "-L", deark_zip_url, "-o", "package.zip")
	err = deark_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	deark_bin_url := "https://entropymine.com/deark/releases/deark-1.6.8.bin"
	deark_cmd_bin := exec.Command("curl", "-L", deark_bin_url, "-o", "binary.bin")
	err = deark_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	deark_src_url := "https://entropymine.com/deark/releases/deark-1.6.8.src.tar.gz"
	deark_cmd_src := exec.Command("curl", "-L", deark_src_url, "-o", "source.tar.gz")
	err = deark_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	deark_cmd_direct := exec.Command("./binary")
	err = deark_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
