package main

// TaskSpooler - Batch system to run tasks one after another
// Homepage: https://viric.name/soft/ts/

import (
	"fmt"
	
	"os/exec"
)

func installTaskSpooler() {
	// Método 1: Descargar y extraer .tar.gz
	taskspooler_tar_url := "https://viric.name/soft/ts/ts-1.0.3.tar.gz"
	taskspooler_cmd_tar := exec.Command("curl", "-L", taskspooler_tar_url, "-o", "package.tar.gz")
	err := taskspooler_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	taskspooler_zip_url := "https://viric.name/soft/ts/ts-1.0.3.zip"
	taskspooler_cmd_zip := exec.Command("curl", "-L", taskspooler_zip_url, "-o", "package.zip")
	err = taskspooler_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	taskspooler_bin_url := "https://viric.name/soft/ts/ts-1.0.3.bin"
	taskspooler_cmd_bin := exec.Command("curl", "-L", taskspooler_bin_url, "-o", "binary.bin")
	err = taskspooler_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	taskspooler_src_url := "https://viric.name/soft/ts/ts-1.0.3.src.tar.gz"
	taskspooler_cmd_src := exec.Command("curl", "-L", taskspooler_src_url, "-o", "source.tar.gz")
	err = taskspooler_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	taskspooler_cmd_direct := exec.Command("./binary")
	err = taskspooler_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
