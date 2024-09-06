package main

// Befunge93 - Esoteric programming language
// Homepage: https://catseye.tc/article/Languages.md#befunge-93

import (
	"fmt"
	
	"os/exec"
)

func installBefunge93() {
	// Método 1: Descargar y extraer .tar.gz
	befunge93_tar_url := "https://catseye.tc/distfiles/befunge-93-2.25.zip"
	befunge93_cmd_tar := exec.Command("curl", "-L", befunge93_tar_url, "-o", "package.tar.gz")
	err := befunge93_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	befunge93_zip_url := "https://catseye.tc/distfiles/befunge-93-2.25.zip"
	befunge93_cmd_zip := exec.Command("curl", "-L", befunge93_zip_url, "-o", "package.zip")
	err = befunge93_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	befunge93_bin_url := "https://catseye.tc/distfiles/befunge-93-2.25.zip"
	befunge93_cmd_bin := exec.Command("curl", "-L", befunge93_bin_url, "-o", "binary.bin")
	err = befunge93_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	befunge93_src_url := "https://catseye.tc/distfiles/befunge-93-2.25.zip"
	befunge93_cmd_src := exec.Command("curl", "-L", befunge93_src_url, "-o", "source.tar.gz")
	err = befunge93_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	befunge93_cmd_direct := exec.Command("./binary")
	err = befunge93_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
