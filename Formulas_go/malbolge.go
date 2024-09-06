package main

// Malbolge - Deliberately difficult to program esoteric programming language
// Homepage: https://esoteric.sange.fi/orphaned/malbolge/README.txt

import (
	"fmt"
	
	"os/exec"
)

func installMalbolge() {
	// Método 1: Descargar y extraer .tar.gz
	malbolge_tar_url := "https://esoteric.sange.fi/orphaned/malbolge/malbolge.c"
	malbolge_cmd_tar := exec.Command("curl", "-L", malbolge_tar_url, "-o", "package.tar.gz")
	err := malbolge_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	malbolge_zip_url := "https://esoteric.sange.fi/orphaned/malbolge/malbolge.c"
	malbolge_cmd_zip := exec.Command("curl", "-L", malbolge_zip_url, "-o", "package.zip")
	err = malbolge_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	malbolge_bin_url := "https://esoteric.sange.fi/orphaned/malbolge/malbolge.c"
	malbolge_cmd_bin := exec.Command("curl", "-L", malbolge_bin_url, "-o", "binary.bin")
	err = malbolge_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	malbolge_src_url := "https://esoteric.sange.fi/orphaned/malbolge/malbolge.c"
	malbolge_cmd_src := exec.Command("curl", "-L", malbolge_src_url, "-o", "source.tar.gz")
	err = malbolge_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	malbolge_cmd_direct := exec.Command("./binary")
	err = malbolge_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
