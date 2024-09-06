package main

// Chkrootkit - Rootkit detector
// Homepage: https://www.chkrootkit.org/

import (
	"fmt"
	
	"os/exec"
)

func installChkrootkit() {
	// Método 1: Descargar y extraer .tar.gz
	chkrootkit_tar_url := "ftp://ftp.chkrootkit.org/pub/seg/pac/chkrootkit-0.58.tar.gz"
	chkrootkit_cmd_tar := exec.Command("curl", "-L", chkrootkit_tar_url, "-o", "package.tar.gz")
	err := chkrootkit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	chkrootkit_zip_url := "ftp://ftp.chkrootkit.org/pub/seg/pac/chkrootkit-0.58.zip"
	chkrootkit_cmd_zip := exec.Command("curl", "-L", chkrootkit_zip_url, "-o", "package.zip")
	err = chkrootkit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	chkrootkit_bin_url := "ftp://ftp.chkrootkit.org/pub/seg/pac/chkrootkit-0.58.bin"
	chkrootkit_cmd_bin := exec.Command("curl", "-L", chkrootkit_bin_url, "-o", "binary.bin")
	err = chkrootkit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	chkrootkit_src_url := "ftp://ftp.chkrootkit.org/pub/seg/pac/chkrootkit-0.58.src.tar.gz"
	chkrootkit_cmd_src := exec.Command("curl", "-L", chkrootkit_src_url, "-o", "source.tar.gz")
	err = chkrootkit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	chkrootkit_cmd_direct := exec.Command("./binary")
	err = chkrootkit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
