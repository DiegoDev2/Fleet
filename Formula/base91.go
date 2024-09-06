package main

// Base91 - Utility to encode and decode base91 files
// Homepage: https://base91.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installBase91() {
	// Método 1: Descargar y extraer .tar.gz
	base91_tar_url := "https://downloads.sourceforge.net/project/base91/basE91/0.6.0/base91-0.6.0.tar.gz"
	base91_cmd_tar := exec.Command("curl", "-L", base91_tar_url, "-o", "package.tar.gz")
	err := base91_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	base91_zip_url := "https://downloads.sourceforge.net/project/base91/basE91/0.6.0/base91-0.6.0.zip"
	base91_cmd_zip := exec.Command("curl", "-L", base91_zip_url, "-o", "package.zip")
	err = base91_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	base91_bin_url := "https://downloads.sourceforge.net/project/base91/basE91/0.6.0/base91-0.6.0.bin"
	base91_cmd_bin := exec.Command("curl", "-L", base91_bin_url, "-o", "binary.bin")
	err = base91_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	base91_src_url := "https://downloads.sourceforge.net/project/base91/basE91/0.6.0/base91-0.6.0.src.tar.gz"
	base91_cmd_src := exec.Command("curl", "-L", base91_src_url, "-o", "source.tar.gz")
	err = base91_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	base91_cmd_direct := exec.Command("./binary")
	err = base91_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
