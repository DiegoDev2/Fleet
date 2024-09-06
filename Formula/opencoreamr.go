package main

// OpencoreAmr - Audio codecs extracted from Android open source project
// Homepage: https://opencore-amr.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installOpencoreAmr() {
	// Método 1: Descargar y extraer .tar.gz
	opencoreamr_tar_url := "https://downloads.sourceforge.net/project/opencore-amr/opencore-amr/opencore-amr-0.1.6.tar.gz"
	opencoreamr_cmd_tar := exec.Command("curl", "-L", opencoreamr_tar_url, "-o", "package.tar.gz")
	err := opencoreamr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	opencoreamr_zip_url := "https://downloads.sourceforge.net/project/opencore-amr/opencore-amr/opencore-amr-0.1.6.zip"
	opencoreamr_cmd_zip := exec.Command("curl", "-L", opencoreamr_zip_url, "-o", "package.zip")
	err = opencoreamr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	opencoreamr_bin_url := "https://downloads.sourceforge.net/project/opencore-amr/opencore-amr/opencore-amr-0.1.6.bin"
	opencoreamr_cmd_bin := exec.Command("curl", "-L", opencoreamr_bin_url, "-o", "binary.bin")
	err = opencoreamr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	opencoreamr_src_url := "https://downloads.sourceforge.net/project/opencore-amr/opencore-amr/opencore-amr-0.1.6.src.tar.gz"
	opencoreamr_cmd_src := exec.Command("curl", "-L", opencoreamr_src_url, "-o", "source.tar.gz")
	err = opencoreamr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	opencoreamr_cmd_direct := exec.Command("./binary")
	err = opencoreamr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
