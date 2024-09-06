package main

// Vbindiff - Visual Binary Diff
// Homepage: https://www.cjmweb.net/vbindiff/

import (
	"fmt"
	
	"os/exec"
)

func installVbindiff() {
	// Método 1: Descargar y extraer .tar.gz
	vbindiff_tar_url := "https://www.cjmweb.net/vbindiff/vbindiff-3.0_beta5.tar.gz"
	vbindiff_cmd_tar := exec.Command("curl", "-L", vbindiff_tar_url, "-o", "package.tar.gz")
	err := vbindiff_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vbindiff_zip_url := "https://www.cjmweb.net/vbindiff/vbindiff-3.0_beta5.zip"
	vbindiff_cmd_zip := exec.Command("curl", "-L", vbindiff_zip_url, "-o", "package.zip")
	err = vbindiff_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vbindiff_bin_url := "https://www.cjmweb.net/vbindiff/vbindiff-3.0_beta5.bin"
	vbindiff_cmd_bin := exec.Command("curl", "-L", vbindiff_bin_url, "-o", "binary.bin")
	err = vbindiff_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vbindiff_src_url := "https://www.cjmweb.net/vbindiff/vbindiff-3.0_beta5.src.tar.gz"
	vbindiff_cmd_src := exec.Command("curl", "-L", vbindiff_src_url, "-o", "source.tar.gz")
	err = vbindiff_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vbindiff_cmd_direct := exec.Command("./binary")
	err = vbindiff_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
