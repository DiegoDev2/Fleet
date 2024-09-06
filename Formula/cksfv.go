package main

// Cksfv - File verification utility
// Homepage: https://zakalwe.fi/~shd/foss/cksfv/

import (
	"fmt"
	
	"os/exec"
)

func installCksfv() {
	// Método 1: Descargar y extraer .tar.gz
	cksfv_tar_url := "https://zakalwe.fi/~shd/foss/cksfv/files/cksfv-1.3.15.tar.bz2"
	cksfv_cmd_tar := exec.Command("curl", "-L", cksfv_tar_url, "-o", "package.tar.gz")
	err := cksfv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cksfv_zip_url := "https://zakalwe.fi/~shd/foss/cksfv/files/cksfv-1.3.15.tar.bz2"
	cksfv_cmd_zip := exec.Command("curl", "-L", cksfv_zip_url, "-o", "package.zip")
	err = cksfv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cksfv_bin_url := "https://zakalwe.fi/~shd/foss/cksfv/files/cksfv-1.3.15.tar.bz2"
	cksfv_cmd_bin := exec.Command("curl", "-L", cksfv_bin_url, "-o", "binary.bin")
	err = cksfv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cksfv_src_url := "https://zakalwe.fi/~shd/foss/cksfv/files/cksfv-1.3.15.tar.bz2"
	cksfv_cmd_src := exec.Command("curl", "-L", cksfv_src_url, "-o", "source.tar.gz")
	err = cksfv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cksfv_cmd_direct := exec.Command("./binary")
	err = cksfv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
