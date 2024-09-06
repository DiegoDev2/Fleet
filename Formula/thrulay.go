package main

// Thrulay - Measure performance of a network
// Homepage: https://sourceforge.net/projects/thrulay/

import (
	"fmt"
	
	"os/exec"
)

func installThrulay() {
	// Método 1: Descargar y extraer .tar.gz
	thrulay_tar_url := "https://downloads.sourceforge.net/project/thrulay/thrulay/0.9/thrulay-0.9.tar.gz"
	thrulay_cmd_tar := exec.Command("curl", "-L", thrulay_tar_url, "-o", "package.tar.gz")
	err := thrulay_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	thrulay_zip_url := "https://downloads.sourceforge.net/project/thrulay/thrulay/0.9/thrulay-0.9.zip"
	thrulay_cmd_zip := exec.Command("curl", "-L", thrulay_zip_url, "-o", "package.zip")
	err = thrulay_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	thrulay_bin_url := "https://downloads.sourceforge.net/project/thrulay/thrulay/0.9/thrulay-0.9.bin"
	thrulay_cmd_bin := exec.Command("curl", "-L", thrulay_bin_url, "-o", "binary.bin")
	err = thrulay_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	thrulay_src_url := "https://downloads.sourceforge.net/project/thrulay/thrulay/0.9/thrulay-0.9.src.tar.gz"
	thrulay_cmd_src := exec.Command("curl", "-L", thrulay_src_url, "-o", "source.tar.gz")
	err = thrulay_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	thrulay_cmd_direct := exec.Command("./binary")
	err = thrulay_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
