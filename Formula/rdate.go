package main

// Rdate - Set the system's date from a remote host
// Homepage: https://www.aelius.com/njh/rdate/

import (
	"fmt"
	
	"os/exec"
)

func installRdate() {
	// Método 1: Descargar y extraer .tar.gz
	rdate_tar_url := "https://www.aelius.com/njh/rdate/rdate-1.5.tar.gz"
	rdate_cmd_tar := exec.Command("curl", "-L", rdate_tar_url, "-o", "package.tar.gz")
	err := rdate_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rdate_zip_url := "https://www.aelius.com/njh/rdate/rdate-1.5.zip"
	rdate_cmd_zip := exec.Command("curl", "-L", rdate_zip_url, "-o", "package.zip")
	err = rdate_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rdate_bin_url := "https://www.aelius.com/njh/rdate/rdate-1.5.bin"
	rdate_cmd_bin := exec.Command("curl", "-L", rdate_bin_url, "-o", "binary.bin")
	err = rdate_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rdate_src_url := "https://www.aelius.com/njh/rdate/rdate-1.5.src.tar.gz"
	rdate_cmd_src := exec.Command("curl", "-L", rdate_src_url, "-o", "source.tar.gz")
	err = rdate_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rdate_cmd_direct := exec.Command("./binary")
	err = rdate_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
