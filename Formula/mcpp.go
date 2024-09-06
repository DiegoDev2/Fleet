package main

// Mcpp - Alternative C/C++ preprocessor
// Homepage: https://mcpp.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installMcpp() {
	// Método 1: Descargar y extraer .tar.gz
	mcpp_tar_url := "https://downloads.sourceforge.net/project/mcpp/mcpp/V.2.7.2/mcpp-2.7.2.tar.gz"
	mcpp_cmd_tar := exec.Command("curl", "-L", mcpp_tar_url, "-o", "package.tar.gz")
	err := mcpp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mcpp_zip_url := "https://downloads.sourceforge.net/project/mcpp/mcpp/V.2.7.2/mcpp-2.7.2.zip"
	mcpp_cmd_zip := exec.Command("curl", "-L", mcpp_zip_url, "-o", "package.zip")
	err = mcpp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mcpp_bin_url := "https://downloads.sourceforge.net/project/mcpp/mcpp/V.2.7.2/mcpp-2.7.2.bin"
	mcpp_cmd_bin := exec.Command("curl", "-L", mcpp_bin_url, "-o", "binary.bin")
	err = mcpp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mcpp_src_url := "https://downloads.sourceforge.net/project/mcpp/mcpp/V.2.7.2/mcpp-2.7.2.src.tar.gz"
	mcpp_cmd_src := exec.Command("curl", "-L", mcpp_src_url, "-o", "source.tar.gz")
	err = mcpp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mcpp_cmd_direct := exec.Command("./binary")
	err = mcpp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
