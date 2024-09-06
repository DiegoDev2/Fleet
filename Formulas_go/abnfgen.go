package main

// Abnfgen - Quickly generate random documents that match an ABFN grammar
// Homepage: https://www.quut.com/abnfgen/

import (
	"fmt"
	
	"os/exec"
)

func installAbnfgen() {
	// Método 1: Descargar y extraer .tar.gz
	abnfgen_tar_url := "https://www.quut.com/abnfgen/abnfgen-0.21.tar.gz"
	abnfgen_cmd_tar := exec.Command("curl", "-L", abnfgen_tar_url, "-o", "package.tar.gz")
	err := abnfgen_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	abnfgen_zip_url := "https://www.quut.com/abnfgen/abnfgen-0.21.zip"
	abnfgen_cmd_zip := exec.Command("curl", "-L", abnfgen_zip_url, "-o", "package.zip")
	err = abnfgen_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	abnfgen_bin_url := "https://www.quut.com/abnfgen/abnfgen-0.21.bin"
	abnfgen_cmd_bin := exec.Command("curl", "-L", abnfgen_bin_url, "-o", "binary.bin")
	err = abnfgen_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	abnfgen_src_url := "https://www.quut.com/abnfgen/abnfgen-0.21.src.tar.gz"
	abnfgen_cmd_src := exec.Command("curl", "-L", abnfgen_src_url, "-o", "source.tar.gz")
	err = abnfgen_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	abnfgen_cmd_direct := exec.Command("./binary")
	err = abnfgen_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
