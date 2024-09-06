package main

// Vapor - Command-line tool for Vapor (Server-side Swift web framework)
// Homepage: https://vapor.codes

import (
	"fmt"
	
	"os/exec"
)

func installVapor() {
	// Método 1: Descargar y extraer .tar.gz
	vapor_tar_url := "https://github.com/vapor/toolbox/archive/refs/tags/18.7.5.tar.gz"
	vapor_cmd_tar := exec.Command("curl", "-L", vapor_tar_url, "-o", "package.tar.gz")
	err := vapor_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vapor_zip_url := "https://github.com/vapor/toolbox/archive/refs/tags/18.7.5.zip"
	vapor_cmd_zip := exec.Command("curl", "-L", vapor_zip_url, "-o", "package.zip")
	err = vapor_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vapor_bin_url := "https://github.com/vapor/toolbox/archive/refs/tags/18.7.5.bin"
	vapor_cmd_bin := exec.Command("curl", "-L", vapor_bin_url, "-o", "binary.bin")
	err = vapor_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vapor_src_url := "https://github.com/vapor/toolbox/archive/refs/tags/18.7.5.src.tar.gz"
	vapor_cmd_src := exec.Command("curl", "-L", vapor_src_url, "-o", "source.tar.gz")
	err = vapor_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vapor_cmd_direct := exec.Command("./binary")
	err = vapor_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
