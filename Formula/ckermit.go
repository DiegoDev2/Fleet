package main

// CKermit - Scriptable network and serial communication for UNIX and VMS
// Homepage: https://www.kermitproject.org/

import (
	"fmt"
	
	"os/exec"
)

func installCKermit() {
	// Método 1: Descargar y extraer .tar.gz
	ckermit_tar_url := "https://www.kermitproject.org/ftp/kermit/archives/cku302.tar.gz"
	ckermit_cmd_tar := exec.Command("curl", "-L", ckermit_tar_url, "-o", "package.tar.gz")
	err := ckermit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ckermit_zip_url := "https://www.kermitproject.org/ftp/kermit/archives/cku302.zip"
	ckermit_cmd_zip := exec.Command("curl", "-L", ckermit_zip_url, "-o", "package.zip")
	err = ckermit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ckermit_bin_url := "https://www.kermitproject.org/ftp/kermit/archives/cku302.bin"
	ckermit_cmd_bin := exec.Command("curl", "-L", ckermit_bin_url, "-o", "binary.bin")
	err = ckermit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ckermit_src_url := "https://www.kermitproject.org/ftp/kermit/archives/cku302.src.tar.gz"
	ckermit_cmd_src := exec.Command("curl", "-L", ckermit_src_url, "-o", "source.tar.gz")
	err = ckermit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ckermit_cmd_direct := exec.Command("./binary")
	err = ckermit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
