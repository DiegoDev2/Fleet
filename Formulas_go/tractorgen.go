package main

// Tractorgen - Generates ASCII tractor art
// Homepage: https://vergenet.net/~conrad/software/tractorgen/

import (
	"fmt"
	
	"os/exec"
)

func installTractorgen() {
	// Método 1: Descargar y extraer .tar.gz
	tractorgen_tar_url := "https://vergenet.net/~conrad/software/tractorgen/dl/tractorgen-0.31.7.tar.gz"
	tractorgen_cmd_tar := exec.Command("curl", "-L", tractorgen_tar_url, "-o", "package.tar.gz")
	err := tractorgen_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tractorgen_zip_url := "https://vergenet.net/~conrad/software/tractorgen/dl/tractorgen-0.31.7.zip"
	tractorgen_cmd_zip := exec.Command("curl", "-L", tractorgen_zip_url, "-o", "package.zip")
	err = tractorgen_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tractorgen_bin_url := "https://vergenet.net/~conrad/software/tractorgen/dl/tractorgen-0.31.7.bin"
	tractorgen_cmd_bin := exec.Command("curl", "-L", tractorgen_bin_url, "-o", "binary.bin")
	err = tractorgen_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tractorgen_src_url := "https://vergenet.net/~conrad/software/tractorgen/dl/tractorgen-0.31.7.src.tar.gz"
	tractorgen_cmd_src := exec.Command("curl", "-L", tractorgen_src_url, "-o", "source.tar.gz")
	err = tractorgen_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tractorgen_cmd_direct := exec.Command("./binary")
	err = tractorgen_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
