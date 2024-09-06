package main

// Vpcs - Virtual PC simulator for testing IP routing
// Homepage: https://vpcs.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installVpcs() {
	// Método 1: Descargar y extraer .tar.gz
	vpcs_tar_url := "https://downloads.sourceforge.net/project/vpcs/0.8/vpcs-0.8-src.tbz"
	vpcs_cmd_tar := exec.Command("curl", "-L", vpcs_tar_url, "-o", "package.tar.gz")
	err := vpcs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vpcs_zip_url := "https://downloads.sourceforge.net/project/vpcs/0.8/vpcs-0.8-src.tbz"
	vpcs_cmd_zip := exec.Command("curl", "-L", vpcs_zip_url, "-o", "package.zip")
	err = vpcs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vpcs_bin_url := "https://downloads.sourceforge.net/project/vpcs/0.8/vpcs-0.8-src.tbz"
	vpcs_cmd_bin := exec.Command("curl", "-L", vpcs_bin_url, "-o", "binary.bin")
	err = vpcs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vpcs_src_url := "https://downloads.sourceforge.net/project/vpcs/0.8/vpcs-0.8-src.tbz"
	vpcs_cmd_src := exec.Command("curl", "-L", vpcs_src_url, "-o", "source.tar.gz")
	err = vpcs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vpcs_cmd_direct := exec.Command("./binary")
	err = vpcs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
