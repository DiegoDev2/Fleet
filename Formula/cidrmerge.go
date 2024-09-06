package main

// Cidrmerge - CIDR merging with network exclusion
// Homepage: https://cidrmerge.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installCidrmerge() {
	// Método 1: Descargar y extraer .tar.gz
	cidrmerge_tar_url := "https://downloads.sourceforge.net/project/cidrmerge/cidrmerge/cidrmerge-1.5.3/cidrmerge-1.5.3.tar.gz"
	cidrmerge_cmd_tar := exec.Command("curl", "-L", cidrmerge_tar_url, "-o", "package.tar.gz")
	err := cidrmerge_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cidrmerge_zip_url := "https://downloads.sourceforge.net/project/cidrmerge/cidrmerge/cidrmerge-1.5.3/cidrmerge-1.5.3.zip"
	cidrmerge_cmd_zip := exec.Command("curl", "-L", cidrmerge_zip_url, "-o", "package.zip")
	err = cidrmerge_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cidrmerge_bin_url := "https://downloads.sourceforge.net/project/cidrmerge/cidrmerge/cidrmerge-1.5.3/cidrmerge-1.5.3.bin"
	cidrmerge_cmd_bin := exec.Command("curl", "-L", cidrmerge_bin_url, "-o", "binary.bin")
	err = cidrmerge_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cidrmerge_src_url := "https://downloads.sourceforge.net/project/cidrmerge/cidrmerge/cidrmerge-1.5.3/cidrmerge-1.5.3.src.tar.gz"
	cidrmerge_cmd_src := exec.Command("curl", "-L", cidrmerge_src_url, "-o", "source.tar.gz")
	err = cidrmerge_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cidrmerge_cmd_direct := exec.Command("./binary")
	err = cidrmerge_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
