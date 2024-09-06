package main

// Pwgen - Password generator
// Homepage: https://pwgen.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installPwgen() {
	// Método 1: Descargar y extraer .tar.gz
	pwgen_tar_url := "https://downloads.sourceforge.net/project/pwgen/pwgen/2.08/pwgen-2.08.tar.gz"
	pwgen_cmd_tar := exec.Command("curl", "-L", pwgen_tar_url, "-o", "package.tar.gz")
	err := pwgen_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pwgen_zip_url := "https://downloads.sourceforge.net/project/pwgen/pwgen/2.08/pwgen-2.08.zip"
	pwgen_cmd_zip := exec.Command("curl", "-L", pwgen_zip_url, "-o", "package.zip")
	err = pwgen_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pwgen_bin_url := "https://downloads.sourceforge.net/project/pwgen/pwgen/2.08/pwgen-2.08.bin"
	pwgen_cmd_bin := exec.Command("curl", "-L", pwgen_bin_url, "-o", "binary.bin")
	err = pwgen_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pwgen_src_url := "https://downloads.sourceforge.net/project/pwgen/pwgen/2.08/pwgen-2.08.src.tar.gz"
	pwgen_cmd_src := exec.Command("curl", "-L", pwgen_src_url, "-o", "source.tar.gz")
	err = pwgen_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pwgen_cmd_direct := exec.Command("./binary")
	err = pwgen_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
