package main

// Primer3 - Program for designing PCR primers
// Homepage: https://primer3.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installPrimer3() {
	// Método 1: Descargar y extraer .tar.gz
	primer3_tar_url := "https://downloads.sourceforge.net/project/primer3/primer3/2.4.0/primer3-2.4.0.tar.gz"
	primer3_cmd_tar := exec.Command("curl", "-L", primer3_tar_url, "-o", "package.tar.gz")
	err := primer3_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	primer3_zip_url := "https://downloads.sourceforge.net/project/primer3/primer3/2.4.0/primer3-2.4.0.zip"
	primer3_cmd_zip := exec.Command("curl", "-L", primer3_zip_url, "-o", "package.zip")
	err = primer3_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	primer3_bin_url := "https://downloads.sourceforge.net/project/primer3/primer3/2.4.0/primer3-2.4.0.bin"
	primer3_cmd_bin := exec.Command("curl", "-L", primer3_bin_url, "-o", "binary.bin")
	err = primer3_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	primer3_src_url := "https://downloads.sourceforge.net/project/primer3/primer3/2.4.0/primer3-2.4.0.src.tar.gz"
	primer3_cmd_src := exec.Command("curl", "-L", primer3_src_url, "-o", "source.tar.gz")
	err = primer3_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	primer3_cmd_direct := exec.Command("./binary")
	err = primer3_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
