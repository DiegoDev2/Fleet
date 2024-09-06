package main

// Runcocoa - Tools to run Cocoa/Objective-C and C code from the command-line
// Homepage: https://github.com/michaeltyson/Commandline-Cocoa

import (
	"fmt"
	
	"os/exec"
)

func installRuncocoa() {
	// Método 1: Descargar y extraer .tar.gz
	runcocoa_tar_url := "https://github.com/michaeltyson/Commandline-Cocoa/archive/834f73b4b5d0d2be0d336c9869973f5f0db55949.tar.gz"
	runcocoa_cmd_tar := exec.Command("curl", "-L", runcocoa_tar_url, "-o", "package.tar.gz")
	err := runcocoa_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	runcocoa_zip_url := "https://github.com/michaeltyson/Commandline-Cocoa/archive/834f73b4b5d0d2be0d336c9869973f5f0db55949.zip"
	runcocoa_cmd_zip := exec.Command("curl", "-L", runcocoa_zip_url, "-o", "package.zip")
	err = runcocoa_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	runcocoa_bin_url := "https://github.com/michaeltyson/Commandline-Cocoa/archive/834f73b4b5d0d2be0d336c9869973f5f0db55949.bin"
	runcocoa_cmd_bin := exec.Command("curl", "-L", runcocoa_bin_url, "-o", "binary.bin")
	err = runcocoa_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	runcocoa_src_url := "https://github.com/michaeltyson/Commandline-Cocoa/archive/834f73b4b5d0d2be0d336c9869973f5f0db55949.src.tar.gz"
	runcocoa_cmd_src := exec.Command("curl", "-L", runcocoa_src_url, "-o", "source.tar.gz")
	err = runcocoa_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	runcocoa_cmd_direct := exec.Command("./binary")
	err = runcocoa_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
