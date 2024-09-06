package main

// Grap - Language for typesetting graphs
// Homepage: https://www.lunabase.org/~faber/Vault/software/grap/

import (
	"fmt"
	
	"os/exec"
)

func installGrap() {
	// Método 1: Descargar y extraer .tar.gz
	grap_tar_url := "https://www.lunabase.org/~faber/Vault/software/grap/grap-1.48.tar.gz"
	grap_cmd_tar := exec.Command("curl", "-L", grap_tar_url, "-o", "package.tar.gz")
	err := grap_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	grap_zip_url := "https://www.lunabase.org/~faber/Vault/software/grap/grap-1.48.zip"
	grap_cmd_zip := exec.Command("curl", "-L", grap_zip_url, "-o", "package.zip")
	err = grap_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	grap_bin_url := "https://www.lunabase.org/~faber/Vault/software/grap/grap-1.48.bin"
	grap_cmd_bin := exec.Command("curl", "-L", grap_bin_url, "-o", "binary.bin")
	err = grap_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	grap_src_url := "https://www.lunabase.org/~faber/Vault/software/grap/grap-1.48.src.tar.gz"
	grap_cmd_src := exec.Command("curl", "-L", grap_src_url, "-o", "source.tar.gz")
	err = grap_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	grap_cmd_direct := exec.Command("./binary")
	err = grap_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
