package main

// Mpgtx - Toolbox to manipulate MPEG files
// Homepage: https://mpgtx.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installMpgtx() {
	// Método 1: Descargar y extraer .tar.gz
	mpgtx_tar_url := "https://downloads.sourceforge.net/project/mpgtx/mpgtx/1.3.1/mpgtx-1.3.1.tar.gz"
	mpgtx_cmd_tar := exec.Command("curl", "-L", mpgtx_tar_url, "-o", "package.tar.gz")
	err := mpgtx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mpgtx_zip_url := "https://downloads.sourceforge.net/project/mpgtx/mpgtx/1.3.1/mpgtx-1.3.1.zip"
	mpgtx_cmd_zip := exec.Command("curl", "-L", mpgtx_zip_url, "-o", "package.zip")
	err = mpgtx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mpgtx_bin_url := "https://downloads.sourceforge.net/project/mpgtx/mpgtx/1.3.1/mpgtx-1.3.1.bin"
	mpgtx_cmd_bin := exec.Command("curl", "-L", mpgtx_bin_url, "-o", "binary.bin")
	err = mpgtx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mpgtx_src_url := "https://downloads.sourceforge.net/project/mpgtx/mpgtx/1.3.1/mpgtx-1.3.1.src.tar.gz"
	mpgtx_cmd_src := exec.Command("curl", "-L", mpgtx_src_url, "-o", "source.tar.gz")
	err = mpgtx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mpgtx_cmd_direct := exec.Command("./binary")
	err = mpgtx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
