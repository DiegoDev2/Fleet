package main

// Tofrodos - Converts DOS <-> UNIX text files, alias tofromdos
// Homepage: https://www.thefreecountry.com/tofrodos/

import (
	"fmt"
	
	"os/exec"
)

func installTofrodos() {
	// Método 1: Descargar y extraer .tar.gz
	tofrodos_tar_url := "https://tofrodos.sourceforge.io/download/tofrodos-1.7.13.tar.gz"
	tofrodos_cmd_tar := exec.Command("curl", "-L", tofrodos_tar_url, "-o", "package.tar.gz")
	err := tofrodos_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tofrodos_zip_url := "https://tofrodos.sourceforge.io/download/tofrodos-1.7.13.zip"
	tofrodos_cmd_zip := exec.Command("curl", "-L", tofrodos_zip_url, "-o", "package.zip")
	err = tofrodos_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tofrodos_bin_url := "https://tofrodos.sourceforge.io/download/tofrodos-1.7.13.bin"
	tofrodos_cmd_bin := exec.Command("curl", "-L", tofrodos_bin_url, "-o", "binary.bin")
	err = tofrodos_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tofrodos_src_url := "https://tofrodos.sourceforge.io/download/tofrodos-1.7.13.src.tar.gz"
	tofrodos_cmd_src := exec.Command("curl", "-L", tofrodos_src_url, "-o", "source.tar.gz")
	err = tofrodos_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tofrodos_cmd_direct := exec.Command("./binary")
	err = tofrodos_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
