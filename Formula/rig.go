package main

// Rig - Provides fake name and address data
// Homepage: https://rig.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installRig() {
	// Método 1: Descargar y extraer .tar.gz
	rig_tar_url := "https://downloads.sourceforge.net/project/rig/rig/1.11/rig-1.11.tar.gz"
	rig_cmd_tar := exec.Command("curl", "-L", rig_tar_url, "-o", "package.tar.gz")
	err := rig_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rig_zip_url := "https://downloads.sourceforge.net/project/rig/rig/1.11/rig-1.11.zip"
	rig_cmd_zip := exec.Command("curl", "-L", rig_zip_url, "-o", "package.zip")
	err = rig_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rig_bin_url := "https://downloads.sourceforge.net/project/rig/rig/1.11/rig-1.11.bin"
	rig_cmd_bin := exec.Command("curl", "-L", rig_bin_url, "-o", "binary.bin")
	err = rig_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rig_src_url := "https://downloads.sourceforge.net/project/rig/rig/1.11/rig-1.11.src.tar.gz"
	rig_cmd_src := exec.Command("curl", "-L", rig_src_url, "-o", "source.tar.gz")
	err = rig_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rig_cmd_direct := exec.Command("./binary")
	err = rig_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
