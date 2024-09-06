package main

// Finatra - Scala web framework inspired by Sinatra
// Homepage: http://finatra.info/

import (
	"fmt"
	
	"os/exec"
)

func installFinatra() {
	// Método 1: Descargar y extraer .tar.gz
	finatra_tar_url := "https://github.com/twitter/finatra/archive/refs/tags/1.5.3.tar.gz"
	finatra_cmd_tar := exec.Command("curl", "-L", finatra_tar_url, "-o", "package.tar.gz")
	err := finatra_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	finatra_zip_url := "https://github.com/twitter/finatra/archive/refs/tags/1.5.3.zip"
	finatra_cmd_zip := exec.Command("curl", "-L", finatra_zip_url, "-o", "package.zip")
	err = finatra_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	finatra_bin_url := "https://github.com/twitter/finatra/archive/refs/tags/1.5.3.bin"
	finatra_cmd_bin := exec.Command("curl", "-L", finatra_bin_url, "-o", "binary.bin")
	err = finatra_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	finatra_src_url := "https://github.com/twitter/finatra/archive/refs/tags/1.5.3.src.tar.gz"
	finatra_cmd_src := exec.Command("curl", "-L", finatra_src_url, "-o", "source.tar.gz")
	err = finatra_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	finatra_cmd_direct := exec.Command("./binary")
	err = finatra_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
