package main

// Cdpr - Cisco Discovery Protocol Reporter
// Homepage: https://www.monkeymental.com/

import (
	"fmt"
	
	"os/exec"
)

func installCdpr() {
	// Método 1: Descargar y extraer .tar.gz
	cdpr_tar_url := "https://downloads.sourceforge.net/project/cdpr/cdpr/2.4/cdpr-2.4.tgz"
	cdpr_cmd_tar := exec.Command("curl", "-L", cdpr_tar_url, "-o", "package.tar.gz")
	err := cdpr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cdpr_zip_url := "https://downloads.sourceforge.net/project/cdpr/cdpr/2.4/cdpr-2.4.tgz"
	cdpr_cmd_zip := exec.Command("curl", "-L", cdpr_zip_url, "-o", "package.zip")
	err = cdpr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cdpr_bin_url := "https://downloads.sourceforge.net/project/cdpr/cdpr/2.4/cdpr-2.4.tgz"
	cdpr_cmd_bin := exec.Command("curl", "-L", cdpr_bin_url, "-o", "binary.bin")
	err = cdpr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cdpr_src_url := "https://downloads.sourceforge.net/project/cdpr/cdpr/2.4/cdpr-2.4.tgz"
	cdpr_cmd_src := exec.Command("curl", "-L", cdpr_src_url, "-o", "source.tar.gz")
	err = cdpr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cdpr_cmd_direct := exec.Command("./binary")
	err = cdpr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
