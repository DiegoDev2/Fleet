package main

// Cdlabelgen - CD/DVD inserts and envelopes
// Homepage: https://www.aczoom.com/tools/cdinsert/

import (
	"fmt"
	
	"os/exec"
)

func installCdlabelgen() {
	// Método 1: Descargar y extraer .tar.gz
	cdlabelgen_tar_url := "https://www.aczoom.com/pub/tools/cdlabelgen-4.3.0.tgz"
	cdlabelgen_cmd_tar := exec.Command("curl", "-L", cdlabelgen_tar_url, "-o", "package.tar.gz")
	err := cdlabelgen_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cdlabelgen_zip_url := "https://www.aczoom.com/pub/tools/cdlabelgen-4.3.0.tgz"
	cdlabelgen_cmd_zip := exec.Command("curl", "-L", cdlabelgen_zip_url, "-o", "package.zip")
	err = cdlabelgen_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cdlabelgen_bin_url := "https://www.aczoom.com/pub/tools/cdlabelgen-4.3.0.tgz"
	cdlabelgen_cmd_bin := exec.Command("curl", "-L", cdlabelgen_bin_url, "-o", "binary.bin")
	err = cdlabelgen_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cdlabelgen_src_url := "https://www.aczoom.com/pub/tools/cdlabelgen-4.3.0.tgz"
	cdlabelgen_cmd_src := exec.Command("curl", "-L", cdlabelgen_src_url, "-o", "source.tar.gz")
	err = cdlabelgen_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cdlabelgen_cmd_direct := exec.Command("./binary")
	err = cdlabelgen_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
