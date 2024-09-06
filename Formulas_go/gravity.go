package main

// Gravity - Embeddable programming language
// Homepage: https://marcobambini.github.io/gravity/

import (
	"fmt"
	
	"os/exec"
)

func installGravity() {
	// Método 1: Descargar y extraer .tar.gz
	gravity_tar_url := "https://github.com/marcobambini/gravity/archive/refs/tags/0.8.5.tar.gz"
	gravity_cmd_tar := exec.Command("curl", "-L", gravity_tar_url, "-o", "package.tar.gz")
	err := gravity_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gravity_zip_url := "https://github.com/marcobambini/gravity/archive/refs/tags/0.8.5.zip"
	gravity_cmd_zip := exec.Command("curl", "-L", gravity_zip_url, "-o", "package.zip")
	err = gravity_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gravity_bin_url := "https://github.com/marcobambini/gravity/archive/refs/tags/0.8.5.bin"
	gravity_cmd_bin := exec.Command("curl", "-L", gravity_bin_url, "-o", "binary.bin")
	err = gravity_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gravity_src_url := "https://github.com/marcobambini/gravity/archive/refs/tags/0.8.5.src.tar.gz"
	gravity_cmd_src := exec.Command("curl", "-L", gravity_src_url, "-o", "source.tar.gz")
	err = gravity_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gravity_cmd_direct := exec.Command("./binary")
	err = gravity_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
