package main

// CdDiscid - Read CD and get CDDB discid information
// Homepage: https://linukz.org/cd-discid.shtml

import (
	"fmt"
	
	"os/exec"
)

func installCdDiscid() {
	// Método 1: Descargar y extraer .tar.gz
	cddiscid_tar_url := "https://linukz.org/download/cd-discid-1.4.tar.gz"
	cddiscid_cmd_tar := exec.Command("curl", "-L", cddiscid_tar_url, "-o", "package.tar.gz")
	err := cddiscid_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cddiscid_zip_url := "https://linukz.org/download/cd-discid-1.4.zip"
	cddiscid_cmd_zip := exec.Command("curl", "-L", cddiscid_zip_url, "-o", "package.zip")
	err = cddiscid_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cddiscid_bin_url := "https://linukz.org/download/cd-discid-1.4.bin"
	cddiscid_cmd_bin := exec.Command("curl", "-L", cddiscid_bin_url, "-o", "binary.bin")
	err = cddiscid_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cddiscid_src_url := "https://linukz.org/download/cd-discid-1.4.src.tar.gz"
	cddiscid_cmd_src := exec.Command("curl", "-L", cddiscid_src_url, "-o", "source.tar.gz")
	err = cddiscid_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cddiscid_cmd_direct := exec.Command("./binary")
	err = cddiscid_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
