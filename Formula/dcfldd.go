package main

// Dcfldd - Enhanced version of dd for forensics and security
// Homepage: https://dcfldd.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installDcfldd() {
	// Método 1: Descargar y extraer .tar.gz
	dcfldd_tar_url := "https://downloads.sourceforge.net/project/dcfldd/dcfldd/1.3.4-1/dcfldd-1.3.4-1.tar.gz"
	dcfldd_cmd_tar := exec.Command("curl", "-L", dcfldd_tar_url, "-o", "package.tar.gz")
	err := dcfldd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dcfldd_zip_url := "https://downloads.sourceforge.net/project/dcfldd/dcfldd/1.3.4-1/dcfldd-1.3.4-1.zip"
	dcfldd_cmd_zip := exec.Command("curl", "-L", dcfldd_zip_url, "-o", "package.zip")
	err = dcfldd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dcfldd_bin_url := "https://downloads.sourceforge.net/project/dcfldd/dcfldd/1.3.4-1/dcfldd-1.3.4-1.bin"
	dcfldd_cmd_bin := exec.Command("curl", "-L", dcfldd_bin_url, "-o", "binary.bin")
	err = dcfldd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dcfldd_src_url := "https://downloads.sourceforge.net/project/dcfldd/dcfldd/1.3.4-1/dcfldd-1.3.4-1.src.tar.gz"
	dcfldd_cmd_src := exec.Command("curl", "-L", dcfldd_src_url, "-o", "source.tar.gz")
	err = dcfldd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dcfldd_cmd_direct := exec.Command("./binary")
	err = dcfldd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
