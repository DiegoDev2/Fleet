package main

// Smake - Portable make program with automake features
// Homepage: https://codeberg.org/schilytools/schilytools

import (
	"fmt"
	
	"os/exec"
)

func installSmake() {
	// Método 1: Descargar y extraer .tar.gz
	smake_tar_url := "https://codeberg.org/schilytools/schilytools/archive/2024-03-21.tar.gz"
	smake_cmd_tar := exec.Command("curl", "-L", smake_tar_url, "-o", "package.tar.gz")
	err := smake_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	smake_zip_url := "https://codeberg.org/schilytools/schilytools/archive/2024-03-21.zip"
	smake_cmd_zip := exec.Command("curl", "-L", smake_zip_url, "-o", "package.zip")
	err = smake_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	smake_bin_url := "https://codeberg.org/schilytools/schilytools/archive/2024-03-21.bin"
	smake_cmd_bin := exec.Command("curl", "-L", smake_bin_url, "-o", "binary.bin")
	err = smake_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	smake_src_url := "https://codeberg.org/schilytools/schilytools/archive/2024-03-21.src.tar.gz"
	smake_cmd_src := exec.Command("curl", "-L", smake_src_url, "-o", "source.tar.gz")
	err = smake_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	smake_cmd_direct := exec.Command("./binary")
	err = smake_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
