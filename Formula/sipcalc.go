package main

// Sipcalc - Advanced console-based IP subnet calculator
// Homepage: https://www.routemeister.net/projects/sipcalc/

import (
	"fmt"
	
	"os/exec"
)

func installSipcalc() {
	// Método 1: Descargar y extraer .tar.gz
	sipcalc_tar_url := "https://www.routemeister.net/projects/sipcalc/files/sipcalc-1.1.6.tar.gz"
	sipcalc_cmd_tar := exec.Command("curl", "-L", sipcalc_tar_url, "-o", "package.tar.gz")
	err := sipcalc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sipcalc_zip_url := "https://www.routemeister.net/projects/sipcalc/files/sipcalc-1.1.6.zip"
	sipcalc_cmd_zip := exec.Command("curl", "-L", sipcalc_zip_url, "-o", "package.zip")
	err = sipcalc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sipcalc_bin_url := "https://www.routemeister.net/projects/sipcalc/files/sipcalc-1.1.6.bin"
	sipcalc_cmd_bin := exec.Command("curl", "-L", sipcalc_bin_url, "-o", "binary.bin")
	err = sipcalc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sipcalc_src_url := "https://www.routemeister.net/projects/sipcalc/files/sipcalc-1.1.6.src.tar.gz"
	sipcalc_cmd_src := exec.Command("curl", "-L", sipcalc_src_url, "-o", "source.tar.gz")
	err = sipcalc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sipcalc_cmd_direct := exec.Command("./binary")
	err = sipcalc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
