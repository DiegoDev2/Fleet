package main

// Itex2mml - Text filter to convert itex equations to MathML
// Homepage: https://golem.ph.utexas.edu/~distler/blog/itex2MML.html

import (
	"fmt"
	
	"os/exec"
)

func installItex2mml() {
	// Método 1: Descargar y extraer .tar.gz
	itex2mml_tar_url := "https://golem.ph.utexas.edu/~distler/blog/files/itexToMML-1.6.1.tar.gz"
	itex2mml_cmd_tar := exec.Command("curl", "-L", itex2mml_tar_url, "-o", "package.tar.gz")
	err := itex2mml_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	itex2mml_zip_url := "https://golem.ph.utexas.edu/~distler/blog/files/itexToMML-1.6.1.zip"
	itex2mml_cmd_zip := exec.Command("curl", "-L", itex2mml_zip_url, "-o", "package.zip")
	err = itex2mml_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	itex2mml_bin_url := "https://golem.ph.utexas.edu/~distler/blog/files/itexToMML-1.6.1.bin"
	itex2mml_cmd_bin := exec.Command("curl", "-L", itex2mml_bin_url, "-o", "binary.bin")
	err = itex2mml_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	itex2mml_src_url := "https://golem.ph.utexas.edu/~distler/blog/files/itexToMML-1.6.1.src.tar.gz"
	itex2mml_cmd_src := exec.Command("curl", "-L", itex2mml_src_url, "-o", "source.tar.gz")
	err = itex2mml_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	itex2mml_cmd_direct := exec.Command("./binary")
	err = itex2mml_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
