package main

// Openslp - Implementation of Service Location Protocol
// Homepage: http://www.openslp.org

import (
	"fmt"
	
	"os/exec"
)

func installOpenslp() {
	// Método 1: Descargar y extraer .tar.gz
	openslp_tar_url := "https://downloads.sourceforge.net/project/openslp/2.0.0/2.0.0%20Release/openslp-2.0.0.tar.gz"
	openslp_cmd_tar := exec.Command("curl", "-L", openslp_tar_url, "-o", "package.tar.gz")
	err := openslp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openslp_zip_url := "https://downloads.sourceforge.net/project/openslp/2.0.0/2.0.0%20Release/openslp-2.0.0.zip"
	openslp_cmd_zip := exec.Command("curl", "-L", openslp_zip_url, "-o", "package.zip")
	err = openslp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openslp_bin_url := "https://downloads.sourceforge.net/project/openslp/2.0.0/2.0.0%20Release/openslp-2.0.0.bin"
	openslp_cmd_bin := exec.Command("curl", "-L", openslp_bin_url, "-o", "binary.bin")
	err = openslp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openslp_src_url := "https://downloads.sourceforge.net/project/openslp/2.0.0/2.0.0%20Release/openslp-2.0.0.src.tar.gz"
	openslp_cmd_src := exec.Command("curl", "-L", openslp_src_url, "-o", "source.tar.gz")
	err = openslp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openslp_cmd_direct := exec.Command("./binary")
	err = openslp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
