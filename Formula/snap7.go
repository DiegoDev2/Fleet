package main

// Snap7 - Ethernet communication suite that works natively with Siemens S7 PLCs
// Homepage: https://snap7.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installSnap7() {
	// Método 1: Descargar y extraer .tar.gz
	snap7_tar_url := "https://downloads.sourceforge.net/project/snap7/1.4.2/snap7-full-1.4.2.7z"
	snap7_cmd_tar := exec.Command("curl", "-L", snap7_tar_url, "-o", "package.tar.gz")
	err := snap7_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	snap7_zip_url := "https://downloads.sourceforge.net/project/snap7/1.4.2/snap7-full-1.4.2.7z"
	snap7_cmd_zip := exec.Command("curl", "-L", snap7_zip_url, "-o", "package.zip")
	err = snap7_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	snap7_bin_url := "https://downloads.sourceforge.net/project/snap7/1.4.2/snap7-full-1.4.2.7z"
	snap7_cmd_bin := exec.Command("curl", "-L", snap7_bin_url, "-o", "binary.bin")
	err = snap7_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	snap7_src_url := "https://downloads.sourceforge.net/project/snap7/1.4.2/snap7-full-1.4.2.7z"
	snap7_cmd_src := exec.Command("curl", "-L", snap7_src_url, "-o", "source.tar.gz")
	err = snap7_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	snap7_cmd_direct := exec.Command("./binary")
	err = snap7_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
