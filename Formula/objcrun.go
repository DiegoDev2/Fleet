package main

// ObjcRun - Use Objective-C files for shell script-like tasks
// Homepage: https://github.com/iljaiwas/objc-run

import (
	"fmt"
	
	"os/exec"
)

func installObjcRun() {
	// Método 1: Descargar y extraer .tar.gz
	objcrun_tar_url := "https://github.com/iljaiwas/objc-run/archive/refs/tags/1.4.tar.gz"
	objcrun_cmd_tar := exec.Command("curl", "-L", objcrun_tar_url, "-o", "package.tar.gz")
	err := objcrun_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	objcrun_zip_url := "https://github.com/iljaiwas/objc-run/archive/refs/tags/1.4.zip"
	objcrun_cmd_zip := exec.Command("curl", "-L", objcrun_zip_url, "-o", "package.zip")
	err = objcrun_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	objcrun_bin_url := "https://github.com/iljaiwas/objc-run/archive/refs/tags/1.4.bin"
	objcrun_cmd_bin := exec.Command("curl", "-L", objcrun_bin_url, "-o", "binary.bin")
	err = objcrun_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	objcrun_src_url := "https://github.com/iljaiwas/objc-run/archive/refs/tags/1.4.src.tar.gz"
	objcrun_cmd_src := exec.Command("curl", "-L", objcrun_src_url, "-o", "source.tar.gz")
	err = objcrun_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	objcrun_cmd_direct := exec.Command("./binary")
	err = objcrun_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
