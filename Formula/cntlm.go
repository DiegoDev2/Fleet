package main

// Cntlm - NTLM authentication proxy with tunneling
// Homepage: https://cntlm.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installCntlm() {
	// Método 1: Descargar y extraer .tar.gz
	cntlm_tar_url := "https://downloads.sourceforge.net/project/cntlm/cntlm/cntlm%200.92.3/cntlm-0.92.3.tar.bz2"
	cntlm_cmd_tar := exec.Command("curl", "-L", cntlm_tar_url, "-o", "package.tar.gz")
	err := cntlm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cntlm_zip_url := "https://downloads.sourceforge.net/project/cntlm/cntlm/cntlm%200.92.3/cntlm-0.92.3.tar.bz2"
	cntlm_cmd_zip := exec.Command("curl", "-L", cntlm_zip_url, "-o", "package.zip")
	err = cntlm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cntlm_bin_url := "https://downloads.sourceforge.net/project/cntlm/cntlm/cntlm%200.92.3/cntlm-0.92.3.tar.bz2"
	cntlm_cmd_bin := exec.Command("curl", "-L", cntlm_bin_url, "-o", "binary.bin")
	err = cntlm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cntlm_src_url := "https://downloads.sourceforge.net/project/cntlm/cntlm/cntlm%200.92.3/cntlm-0.92.3.tar.bz2"
	cntlm_cmd_src := exec.Command("curl", "-L", cntlm_src_url, "-o", "source.tar.gz")
	err = cntlm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cntlm_cmd_direct := exec.Command("./binary")
	err = cntlm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
