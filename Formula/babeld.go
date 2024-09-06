package main

// Babeld - Loop-avoiding distance-vector routing protocol
// Homepage: https://www.irif.fr/~jch/software/babel/

import (
	"fmt"
	
	"os/exec"
)

func installBabeld() {
	// Método 1: Descargar y extraer .tar.gz
	babeld_tar_url := "https://www.irif.fr/~jch/software/files/babeld-1.13.1.tar.gz"
	babeld_cmd_tar := exec.Command("curl", "-L", babeld_tar_url, "-o", "package.tar.gz")
	err := babeld_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	babeld_zip_url := "https://www.irif.fr/~jch/software/files/babeld-1.13.1.zip"
	babeld_cmd_zip := exec.Command("curl", "-L", babeld_zip_url, "-o", "package.zip")
	err = babeld_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	babeld_bin_url := "https://www.irif.fr/~jch/software/files/babeld-1.13.1.bin"
	babeld_cmd_bin := exec.Command("curl", "-L", babeld_bin_url, "-o", "binary.bin")
	err = babeld_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	babeld_src_url := "https://www.irif.fr/~jch/software/files/babeld-1.13.1.src.tar.gz"
	babeld_cmd_src := exec.Command("curl", "-L", babeld_src_url, "-o", "source.tar.gz")
	err = babeld_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	babeld_cmd_direct := exec.Command("./binary")
	err = babeld_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
