package main

// Pmccabe - Calculate McCabe-style cyclomatic complexity for C/C++ code
// Homepage: https://gitlab.com/pmccabe/pmccabe

import (
	"fmt"
	
	"os/exec"
)

func installPmccabe() {
	// Método 1: Descargar y extraer .tar.gz
	pmccabe_tar_url := "https://gitlab.com/pmccabe/pmccabe/-/archive/v2.8/pmccabe-v2.8.tar.bz2"
	pmccabe_cmd_tar := exec.Command("curl", "-L", pmccabe_tar_url, "-o", "package.tar.gz")
	err := pmccabe_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pmccabe_zip_url := "https://gitlab.com/pmccabe/pmccabe/-/archive/v2.8/pmccabe-v2.8.tar.bz2"
	pmccabe_cmd_zip := exec.Command("curl", "-L", pmccabe_zip_url, "-o", "package.zip")
	err = pmccabe_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pmccabe_bin_url := "https://gitlab.com/pmccabe/pmccabe/-/archive/v2.8/pmccabe-v2.8.tar.bz2"
	pmccabe_cmd_bin := exec.Command("curl", "-L", pmccabe_bin_url, "-o", "binary.bin")
	err = pmccabe_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pmccabe_src_url := "https://gitlab.com/pmccabe/pmccabe/-/archive/v2.8/pmccabe-v2.8.tar.bz2"
	pmccabe_cmd_src := exec.Command("curl", "-L", pmccabe_src_url, "-o", "source.tar.gz")
	err = pmccabe_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pmccabe_cmd_direct := exec.Command("./binary")
	err = pmccabe_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
