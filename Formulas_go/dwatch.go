package main

// Dwatch - Watch programs and perform actions based on a configuration file
// Homepage: https://siag.nu/dwatch/

import (
	"fmt"
	
	"os/exec"
)

func installDwatch() {
	// Método 1: Descargar y extraer .tar.gz
	dwatch_tar_url := "https://siag.nu/pub/dwatch/dwatch-0.1.1.tar.gz"
	dwatch_cmd_tar := exec.Command("curl", "-L", dwatch_tar_url, "-o", "package.tar.gz")
	err := dwatch_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dwatch_zip_url := "https://siag.nu/pub/dwatch/dwatch-0.1.1.zip"
	dwatch_cmd_zip := exec.Command("curl", "-L", dwatch_zip_url, "-o", "package.zip")
	err = dwatch_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dwatch_bin_url := "https://siag.nu/pub/dwatch/dwatch-0.1.1.bin"
	dwatch_cmd_bin := exec.Command("curl", "-L", dwatch_bin_url, "-o", "binary.bin")
	err = dwatch_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dwatch_src_url := "https://siag.nu/pub/dwatch/dwatch-0.1.1.src.tar.gz"
	dwatch_cmd_src := exec.Command("curl", "-L", dwatch_src_url, "-o", "source.tar.gz")
	err = dwatch_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dwatch_cmd_direct := exec.Command("./binary")
	err = dwatch_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
