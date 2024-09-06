package main

// Sec - Event correlation tool for event processing of various kinds
// Homepage: https://simple-evcorr.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installSec() {
	// Método 1: Descargar y extraer .tar.gz
	sec_tar_url := "https://github.com/simple-evcorr/sec/releases/download/2.9.2/sec-2.9.2.tar.gz"
	sec_cmd_tar := exec.Command("curl", "-L", sec_tar_url, "-o", "package.tar.gz")
	err := sec_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sec_zip_url := "https://github.com/simple-evcorr/sec/releases/download/2.9.2/sec-2.9.2.zip"
	sec_cmd_zip := exec.Command("curl", "-L", sec_zip_url, "-o", "package.zip")
	err = sec_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sec_bin_url := "https://github.com/simple-evcorr/sec/releases/download/2.9.2/sec-2.9.2.bin"
	sec_cmd_bin := exec.Command("curl", "-L", sec_bin_url, "-o", "binary.bin")
	err = sec_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sec_src_url := "https://github.com/simple-evcorr/sec/releases/download/2.9.2/sec-2.9.2.src.tar.gz"
	sec_cmd_src := exec.Command("curl", "-L", sec_src_url, "-o", "source.tar.gz")
	err = sec_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sec_cmd_direct := exec.Command("./binary")
	err = sec_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
