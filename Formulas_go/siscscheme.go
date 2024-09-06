package main

// SiscScheme - Extensive Java based Scheme interpreter
// Homepage: https://sisc-scheme.org/

import (
	"fmt"
	
	"os/exec"
)

func installSiscScheme() {
	// Método 1: Descargar y extraer .tar.gz
	siscscheme_tar_url := "https://downloads.sourceforge.net/project/sisc/SISC%20Lite/1.16.6/sisc-lite-1.16.6.tar.gz"
	siscscheme_cmd_tar := exec.Command("curl", "-L", siscscheme_tar_url, "-o", "package.tar.gz")
	err := siscscheme_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	siscscheme_zip_url := "https://downloads.sourceforge.net/project/sisc/SISC%20Lite/1.16.6/sisc-lite-1.16.6.zip"
	siscscheme_cmd_zip := exec.Command("curl", "-L", siscscheme_zip_url, "-o", "package.zip")
	err = siscscheme_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	siscscheme_bin_url := "https://downloads.sourceforge.net/project/sisc/SISC%20Lite/1.16.6/sisc-lite-1.16.6.bin"
	siscscheme_cmd_bin := exec.Command("curl", "-L", siscscheme_bin_url, "-o", "binary.bin")
	err = siscscheme_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	siscscheme_src_url := "https://downloads.sourceforge.net/project/sisc/SISC%20Lite/1.16.6/sisc-lite-1.16.6.src.tar.gz"
	siscscheme_cmd_src := exec.Command("curl", "-L", siscscheme_src_url, "-o", "source.tar.gz")
	err = siscscheme_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	siscscheme_cmd_direct := exec.Command("./binary")
	err = siscscheme_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
