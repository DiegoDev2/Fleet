package main

// Radamsa - Test case generator for robustness testing (a.k.a. a \
// Homepage: https://gitlab.com/akihe/radamsa

import (
	"fmt"
	
	"os/exec"
)

func installRadamsa() {
	// Método 1: Descargar y extraer .tar.gz
	radamsa_tar_url := "https://gitlab.com/akihe/radamsa/-/archive/v0.7/radamsa-v0.7.tar.gz"
	radamsa_cmd_tar := exec.Command("curl", "-L", radamsa_tar_url, "-o", "package.tar.gz")
	err := radamsa_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	radamsa_zip_url := "https://gitlab.com/akihe/radamsa/-/archive/v0.7/radamsa-v0.7.zip"
	radamsa_cmd_zip := exec.Command("curl", "-L", radamsa_zip_url, "-o", "package.zip")
	err = radamsa_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	radamsa_bin_url := "https://gitlab.com/akihe/radamsa/-/archive/v0.7/radamsa-v0.7.bin"
	radamsa_cmd_bin := exec.Command("curl", "-L", radamsa_bin_url, "-o", "binary.bin")
	err = radamsa_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	radamsa_src_url := "https://gitlab.com/akihe/radamsa/-/archive/v0.7/radamsa-v0.7.src.tar.gz"
	radamsa_cmd_src := exec.Command("curl", "-L", radamsa_src_url, "-o", "source.tar.gz")
	err = radamsa_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	radamsa_cmd_direct := exec.Command("./binary")
	err = radamsa_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
