package main

// Avanor - Quick-growing roguelike game with easy ADOM-like UI
// Homepage: https://avanor.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installAvanor() {
	// Método 1: Descargar y extraer .tar.gz
	avanor_tar_url := "https://downloads.sourceforge.net/project/avanor/avanor/0.5.8/avanor-0.5.8-src.tar.bz2"
	avanor_cmd_tar := exec.Command("curl", "-L", avanor_tar_url, "-o", "package.tar.gz")
	err := avanor_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	avanor_zip_url := "https://downloads.sourceforge.net/project/avanor/avanor/0.5.8/avanor-0.5.8-src.tar.bz2"
	avanor_cmd_zip := exec.Command("curl", "-L", avanor_zip_url, "-o", "package.zip")
	err = avanor_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	avanor_bin_url := "https://downloads.sourceforge.net/project/avanor/avanor/0.5.8/avanor-0.5.8-src.tar.bz2"
	avanor_cmd_bin := exec.Command("curl", "-L", avanor_bin_url, "-o", "binary.bin")
	err = avanor_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	avanor_src_url := "https://downloads.sourceforge.net/project/avanor/avanor/0.5.8/avanor-0.5.8-src.tar.bz2"
	avanor_cmd_src := exec.Command("curl", "-L", avanor_src_url, "-o", "source.tar.gz")
	err = avanor_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	avanor_cmd_direct := exec.Command("./binary")
	err = avanor_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
