package main

// Viewvc - Browser interface for CVS and Subversion repositories
// Homepage: https://www.viewvc.org

import (
	"fmt"
	
	"os/exec"
)

func installViewvc() {
	// Método 1: Descargar y extraer .tar.gz
	viewvc_tar_url := "https://github.com/viewvc/viewvc/releases/download/1.2.3/viewvc-1.2.3.tar.gz"
	viewvc_cmd_tar := exec.Command("curl", "-L", viewvc_tar_url, "-o", "package.tar.gz")
	err := viewvc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	viewvc_zip_url := "https://github.com/viewvc/viewvc/releases/download/1.2.3/viewvc-1.2.3.zip"
	viewvc_cmd_zip := exec.Command("curl", "-L", viewvc_zip_url, "-o", "package.zip")
	err = viewvc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	viewvc_bin_url := "https://github.com/viewvc/viewvc/releases/download/1.2.3/viewvc-1.2.3.bin"
	viewvc_cmd_bin := exec.Command("curl", "-L", viewvc_bin_url, "-o", "binary.bin")
	err = viewvc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	viewvc_src_url := "https://github.com/viewvc/viewvc/releases/download/1.2.3/viewvc-1.2.3.src.tar.gz"
	viewvc_cmd_src := exec.Command("curl", "-L", viewvc_src_url, "-o", "source.tar.gz")
	err = viewvc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	viewvc_cmd_direct := exec.Command("./binary")
	err = viewvc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
