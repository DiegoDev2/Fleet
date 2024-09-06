package main

// Radare2 - Reverse engineering framework
// Homepage: https://radare.org

import (
	"fmt"
	
	"os/exec"
)

func installRadare2() {
	// Método 1: Descargar y extraer .tar.gz
	radare2_tar_url := "https://github.com/radareorg/radare2/archive/refs/tags/5.9.4.tar.gz"
	radare2_cmd_tar := exec.Command("curl", "-L", radare2_tar_url, "-o", "package.tar.gz")
	err := radare2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	radare2_zip_url := "https://github.com/radareorg/radare2/archive/refs/tags/5.9.4.zip"
	radare2_cmd_zip := exec.Command("curl", "-L", radare2_zip_url, "-o", "package.zip")
	err = radare2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	radare2_bin_url := "https://github.com/radareorg/radare2/archive/refs/tags/5.9.4.bin"
	radare2_cmd_bin := exec.Command("curl", "-L", radare2_bin_url, "-o", "binary.bin")
	err = radare2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	radare2_src_url := "https://github.com/radareorg/radare2/archive/refs/tags/5.9.4.src.tar.gz"
	radare2_cmd_src := exec.Command("curl", "-L", radare2_src_url, "-o", "source.tar.gz")
	err = radare2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	radare2_cmd_direct := exec.Command("./binary")
	err = radare2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
