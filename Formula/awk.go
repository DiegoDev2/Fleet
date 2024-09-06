package main

// Awk - Text processing scripting language
// Homepage: https://www.cs.princeton.edu/~bwk/btl.mirror/

import (
	"fmt"
	
	"os/exec"
)

func installAwk() {
	// Método 1: Descargar y extraer .tar.gz
	awk_tar_url := "https://github.com/onetrueawk/awk/archive/refs/tags/20240728.tar.gz"
	awk_cmd_tar := exec.Command("curl", "-L", awk_tar_url, "-o", "package.tar.gz")
	err := awk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	awk_zip_url := "https://github.com/onetrueawk/awk/archive/refs/tags/20240728.zip"
	awk_cmd_zip := exec.Command("curl", "-L", awk_zip_url, "-o", "package.zip")
	err = awk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	awk_bin_url := "https://github.com/onetrueawk/awk/archive/refs/tags/20240728.bin"
	awk_cmd_bin := exec.Command("curl", "-L", awk_bin_url, "-o", "binary.bin")
	err = awk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	awk_src_url := "https://github.com/onetrueawk/awk/archive/refs/tags/20240728.src.tar.gz"
	awk_cmd_src := exec.Command("curl", "-L", awk_src_url, "-o", "source.tar.gz")
	err = awk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	awk_cmd_direct := exec.Command("./binary")
	err = awk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
