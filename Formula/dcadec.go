package main

// Dcadec - DTS Coherent Acoustics decoder with support for HD extensions
// Homepage: https://github.com/foo86/dcadec

import (
	"fmt"
	
	"os/exec"
)

func installDcadec() {
	// Método 1: Descargar y extraer .tar.gz
	dcadec_tar_url := "https://github.com/foo86/dcadec.git"
	dcadec_cmd_tar := exec.Command("curl", "-L", dcadec_tar_url, "-o", "package.tar.gz")
	err := dcadec_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dcadec_zip_url := "https://github.com/foo86/dcadec.git"
	dcadec_cmd_zip := exec.Command("curl", "-L", dcadec_zip_url, "-o", "package.zip")
	err = dcadec_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dcadec_bin_url := "https://github.com/foo86/dcadec.git"
	dcadec_cmd_bin := exec.Command("curl", "-L", dcadec_bin_url, "-o", "binary.bin")
	err = dcadec_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dcadec_src_url := "https://github.com/foo86/dcadec.git"
	dcadec_cmd_src := exec.Command("curl", "-L", dcadec_src_url, "-o", "source.tar.gz")
	err = dcadec_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dcadec_cmd_direct := exec.Command("./binary")
	err = dcadec_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
