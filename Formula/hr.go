package main

// Hr - <hr />, for your terminal window
// Homepage: https://github.com/LuRsT/hr

import (
	"fmt"
	
	"os/exec"
)

func installHr() {
	// Método 1: Descargar y extraer .tar.gz
	hr_tar_url := "https://github.com/LuRsT/hr/archive/refs/tags/1.4.tar.gz"
	hr_cmd_tar := exec.Command("curl", "-L", hr_tar_url, "-o", "package.tar.gz")
	err := hr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hr_zip_url := "https://github.com/LuRsT/hr/archive/refs/tags/1.4.zip"
	hr_cmd_zip := exec.Command("curl", "-L", hr_zip_url, "-o", "package.zip")
	err = hr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hr_bin_url := "https://github.com/LuRsT/hr/archive/refs/tags/1.4.bin"
	hr_cmd_bin := exec.Command("curl", "-L", hr_bin_url, "-o", "binary.bin")
	err = hr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hr_src_url := "https://github.com/LuRsT/hr/archive/refs/tags/1.4.src.tar.gz"
	hr_cmd_src := exec.Command("curl", "-L", hr_src_url, "-o", "source.tar.gz")
	err = hr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hr_cmd_direct := exec.Command("./binary")
	err = hr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
