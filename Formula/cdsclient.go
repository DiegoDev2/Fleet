package main

// Cdsclient - Tools for querying CDS databases for astronomical data
// Homepage: https://cdsarc.u-strasbg.fr/doc/cdsclient.html

import (
	"fmt"
	
	"os/exec"
)

func installCdsclient() {
	// Método 1: Descargar y extraer .tar.gz
	cdsclient_tar_url := "https://cdsarc.u-strasbg.fr/ftp/pub/sw/cdsclient-3.84.tar.gz"
	cdsclient_cmd_tar := exec.Command("curl", "-L", cdsclient_tar_url, "-o", "package.tar.gz")
	err := cdsclient_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cdsclient_zip_url := "https://cdsarc.u-strasbg.fr/ftp/pub/sw/cdsclient-3.84.zip"
	cdsclient_cmd_zip := exec.Command("curl", "-L", cdsclient_zip_url, "-o", "package.zip")
	err = cdsclient_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cdsclient_bin_url := "https://cdsarc.u-strasbg.fr/ftp/pub/sw/cdsclient-3.84.bin"
	cdsclient_cmd_bin := exec.Command("curl", "-L", cdsclient_bin_url, "-o", "binary.bin")
	err = cdsclient_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cdsclient_src_url := "https://cdsarc.u-strasbg.fr/ftp/pub/sw/cdsclient-3.84.src.tar.gz"
	cdsclient_cmd_src := exec.Command("curl", "-L", cdsclient_src_url, "-o", "source.tar.gz")
	err = cdsclient_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cdsclient_cmd_direct := exec.Command("./binary")
	err = cdsclient_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
