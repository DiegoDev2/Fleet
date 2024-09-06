package main

// Bbe - Sed-like editor for binary files
// Homepage: https://sourceforge.net/projects/bbe-/

import (
	"fmt"
	
	"os/exec"
)

func installBbe() {
	// Método 1: Descargar y extraer .tar.gz
	bbe_tar_url := "https://downloads.sourceforge.net/project/bbe-/bbe/0.2.2/bbe-0.2.2.tar.gz"
	bbe_cmd_tar := exec.Command("curl", "-L", bbe_tar_url, "-o", "package.tar.gz")
	err := bbe_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bbe_zip_url := "https://downloads.sourceforge.net/project/bbe-/bbe/0.2.2/bbe-0.2.2.zip"
	bbe_cmd_zip := exec.Command("curl", "-L", bbe_zip_url, "-o", "package.zip")
	err = bbe_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bbe_bin_url := "https://downloads.sourceforge.net/project/bbe-/bbe/0.2.2/bbe-0.2.2.bin"
	bbe_cmd_bin := exec.Command("curl", "-L", bbe_bin_url, "-o", "binary.bin")
	err = bbe_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bbe_src_url := "https://downloads.sourceforge.net/project/bbe-/bbe/0.2.2/bbe-0.2.2.src.tar.gz"
	bbe_cmd_src := exec.Command("curl", "-L", bbe_src_url, "-o", "source.tar.gz")
	err = bbe_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bbe_cmd_direct := exec.Command("./binary")
	err = bbe_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
