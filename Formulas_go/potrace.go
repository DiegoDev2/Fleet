package main

// Potrace - Convert bitmaps to vector graphics
// Homepage: https://potrace.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installPotrace() {
	// Método 1: Descargar y extraer .tar.gz
	potrace_tar_url := "https://potrace.sourceforge.net/download/1.16/potrace-1.16.tar.gz"
	potrace_cmd_tar := exec.Command("curl", "-L", potrace_tar_url, "-o", "package.tar.gz")
	err := potrace_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	potrace_zip_url := "https://potrace.sourceforge.net/download/1.16/potrace-1.16.zip"
	potrace_cmd_zip := exec.Command("curl", "-L", potrace_zip_url, "-o", "package.zip")
	err = potrace_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	potrace_bin_url := "https://potrace.sourceforge.net/download/1.16/potrace-1.16.bin"
	potrace_cmd_bin := exec.Command("curl", "-L", potrace_bin_url, "-o", "binary.bin")
	err = potrace_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	potrace_src_url := "https://potrace.sourceforge.net/download/1.16/potrace-1.16.src.tar.gz"
	potrace_cmd_src := exec.Command("curl", "-L", potrace_src_url, "-o", "source.tar.gz")
	err = potrace_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	potrace_cmd_direct := exec.Command("./binary")
	err = potrace_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
