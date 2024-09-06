package main

// Findent - Indent and beautify Fortran sources and generate dependency information
// Homepage: https://www.ratrabbit.nl/ratrabbit/findent/index.html

import (
	"fmt"
	
	"os/exec"
)

func installFindent() {
	// Método 1: Descargar y extraer .tar.gz
	findent_tar_url := "https://downloads.sourceforge.net/project/findent/findent-4.3.3.tar.gz"
	findent_cmd_tar := exec.Command("curl", "-L", findent_tar_url, "-o", "package.tar.gz")
	err := findent_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	findent_zip_url := "https://downloads.sourceforge.net/project/findent/findent-4.3.3.zip"
	findent_cmd_zip := exec.Command("curl", "-L", findent_zip_url, "-o", "package.zip")
	err = findent_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	findent_bin_url := "https://downloads.sourceforge.net/project/findent/findent-4.3.3.bin"
	findent_cmd_bin := exec.Command("curl", "-L", findent_bin_url, "-o", "binary.bin")
	err = findent_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	findent_src_url := "https://downloads.sourceforge.net/project/findent/findent-4.3.3.src.tar.gz"
	findent_cmd_src := exec.Command("curl", "-L", findent_src_url, "-o", "source.tar.gz")
	err = findent_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	findent_cmd_direct := exec.Command("./binary")
	err = findent_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
