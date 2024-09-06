package main

// Cscope - Tool for browsing source code
// Homepage: https://cscope.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installCscope() {
	// Método 1: Descargar y extraer .tar.gz
	cscope_tar_url := "https://downloads.sourceforge.net/project/cscope/cscope/v15.9/cscope-15.9.tar.gz"
	cscope_cmd_tar := exec.Command("curl", "-L", cscope_tar_url, "-o", "package.tar.gz")
	err := cscope_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cscope_zip_url := "https://downloads.sourceforge.net/project/cscope/cscope/v15.9/cscope-15.9.zip"
	cscope_cmd_zip := exec.Command("curl", "-L", cscope_zip_url, "-o", "package.zip")
	err = cscope_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cscope_bin_url := "https://downloads.sourceforge.net/project/cscope/cscope/v15.9/cscope-15.9.bin"
	cscope_cmd_bin := exec.Command("curl", "-L", cscope_bin_url, "-o", "binary.bin")
	err = cscope_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cscope_src_url := "https://downloads.sourceforge.net/project/cscope/cscope/v15.9/cscope-15.9.src.tar.gz"
	cscope_cmd_src := exec.Command("curl", "-L", cscope_src_url, "-o", "source.tar.gz")
	err = cscope_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cscope_cmd_direct := exec.Command("./binary")
	err = cscope_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
