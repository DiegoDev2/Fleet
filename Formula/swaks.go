package main

// Swaks - SMTP command-line test tool
// Homepage: https://www.jetmore.org/john/code/swaks/

import (
	"fmt"
	
	"os/exec"
)

func installSwaks() {
	// Método 1: Descargar y extraer .tar.gz
	swaks_tar_url := "https://www.jetmore.org/john/code/swaks/files/swaks-20240103.0.tar.gz"
	swaks_cmd_tar := exec.Command("curl", "-L", swaks_tar_url, "-o", "package.tar.gz")
	err := swaks_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	swaks_zip_url := "https://www.jetmore.org/john/code/swaks/files/swaks-20240103.0.zip"
	swaks_cmd_zip := exec.Command("curl", "-L", swaks_zip_url, "-o", "package.zip")
	err = swaks_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	swaks_bin_url := "https://www.jetmore.org/john/code/swaks/files/swaks-20240103.0.bin"
	swaks_cmd_bin := exec.Command("curl", "-L", swaks_bin_url, "-o", "binary.bin")
	err = swaks_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	swaks_src_url := "https://www.jetmore.org/john/code/swaks/files/swaks-20240103.0.src.tar.gz"
	swaks_cmd_src := exec.Command("curl", "-L", swaks_src_url, "-o", "source.tar.gz")
	err = swaks_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	swaks_cmd_direct := exec.Command("./binary")
	err = swaks_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
