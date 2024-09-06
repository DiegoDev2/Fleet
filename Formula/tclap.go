package main

// Tclap - Templatized C++ command-line parser library
// Homepage: https://tclap.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installTclap() {
	// Método 1: Descargar y extraer .tar.gz
	tclap_tar_url := "https://downloads.sourceforge.net/project/tclap/tclap-1.2.5.tar.gz"
	tclap_cmd_tar := exec.Command("curl", "-L", tclap_tar_url, "-o", "package.tar.gz")
	err := tclap_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tclap_zip_url := "https://downloads.sourceforge.net/project/tclap/tclap-1.2.5.zip"
	tclap_cmd_zip := exec.Command("curl", "-L", tclap_zip_url, "-o", "package.zip")
	err = tclap_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tclap_bin_url := "https://downloads.sourceforge.net/project/tclap/tclap-1.2.5.bin"
	tclap_cmd_bin := exec.Command("curl", "-L", tclap_bin_url, "-o", "binary.bin")
	err = tclap_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tclap_src_url := "https://downloads.sourceforge.net/project/tclap/tclap-1.2.5.src.tar.gz"
	tclap_cmd_src := exec.Command("curl", "-L", tclap_src_url, "-o", "source.tar.gz")
	err = tclap_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tclap_cmd_direct := exec.Command("./binary")
	err = tclap_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
