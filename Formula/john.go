package main

// John - Featureful UNIX password cracker
// Homepage: https://www.openwall.com/john/

import (
	"fmt"
	
	"os/exec"
)

func installJohn() {
	// Método 1: Descargar y extraer .tar.gz
	john_tar_url := "https://www.openwall.com/john/k/john-1.9.0.tar.xz"
	john_cmd_tar := exec.Command("curl", "-L", john_tar_url, "-o", "package.tar.gz")
	err := john_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	john_zip_url := "https://www.openwall.com/john/k/john-1.9.0.tar.xz"
	john_cmd_zip := exec.Command("curl", "-L", john_zip_url, "-o", "package.zip")
	err = john_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	john_bin_url := "https://www.openwall.com/john/k/john-1.9.0.tar.xz"
	john_cmd_bin := exec.Command("curl", "-L", john_bin_url, "-o", "binary.bin")
	err = john_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	john_src_url := "https://www.openwall.com/john/k/john-1.9.0.tar.xz"
	john_cmd_src := exec.Command("curl", "-L", john_src_url, "-o", "source.tar.gz")
	err = john_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	john_cmd_direct := exec.Command("./binary")
	err = john_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
