package main

// Npush - Logic game similar to Sokoban and Boulder Dash
// Homepage: https://npush.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installNpush() {
	// Método 1: Descargar y extraer .tar.gz
	npush_tar_url := "https://downloads.sourceforge.net/project/npush/npush/0.7/npush-0.7.tgz"
	npush_cmd_tar := exec.Command("curl", "-L", npush_tar_url, "-o", "package.tar.gz")
	err := npush_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	npush_zip_url := "https://downloads.sourceforge.net/project/npush/npush/0.7/npush-0.7.tgz"
	npush_cmd_zip := exec.Command("curl", "-L", npush_zip_url, "-o", "package.zip")
	err = npush_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	npush_bin_url := "https://downloads.sourceforge.net/project/npush/npush/0.7/npush-0.7.tgz"
	npush_cmd_bin := exec.Command("curl", "-L", npush_bin_url, "-o", "binary.bin")
	err = npush_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	npush_src_url := "https://downloads.sourceforge.net/project/npush/npush/0.7/npush-0.7.tgz"
	npush_cmd_src := exec.Command("curl", "-L", npush_src_url, "-o", "source.tar.gz")
	err = npush_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	npush_cmd_direct := exec.Command("./binary")
	err = npush_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
