package main

// Ecm - Prepare CD image files so they compress better
// Homepage: https://web.archive.org/web/20140227165748/www.neillcorlett.com/ecm/

import (
	"fmt"
	
	"os/exec"
)

func installEcm() {
	// Método 1: Descargar y extraer .tar.gz
	ecm_tar_url := "https://web.archive.org/web/20091021035854/www.neillcorlett.com/downloads/ecm100.zip"
	ecm_cmd_tar := exec.Command("curl", "-L", ecm_tar_url, "-o", "package.tar.gz")
	err := ecm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ecm_zip_url := "https://web.archive.org/web/20091021035854/www.neillcorlett.com/downloads/ecm100.zip"
	ecm_cmd_zip := exec.Command("curl", "-L", ecm_zip_url, "-o", "package.zip")
	err = ecm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ecm_bin_url := "https://web.archive.org/web/20091021035854/www.neillcorlett.com/downloads/ecm100.zip"
	ecm_cmd_bin := exec.Command("curl", "-L", ecm_bin_url, "-o", "binary.bin")
	err = ecm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ecm_src_url := "https://web.archive.org/web/20091021035854/www.neillcorlett.com/downloads/ecm100.zip"
	ecm_cmd_src := exec.Command("curl", "-L", ecm_src_url, "-o", "source.tar.gz")
	err = ecm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ecm_cmd_direct := exec.Command("./binary")
	err = ecm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
