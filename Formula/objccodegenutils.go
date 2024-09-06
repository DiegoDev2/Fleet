package main

// ObjcCodegenutils - Three small tools to help work with XCode
// Homepage: https://github.com/puls/objc-codegenutils

import (
	"fmt"
	
	"os/exec"
)

func installObjcCodegenutils() {
	// Método 1: Descargar y extraer .tar.gz
	objccodegenutils_tar_url := "https://github.com/puls/objc-codegenutils/archive/refs/tags/v1.0.tar.gz"
	objccodegenutils_cmd_tar := exec.Command("curl", "-L", objccodegenutils_tar_url, "-o", "package.tar.gz")
	err := objccodegenutils_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	objccodegenutils_zip_url := "https://github.com/puls/objc-codegenutils/archive/refs/tags/v1.0.zip"
	objccodegenutils_cmd_zip := exec.Command("curl", "-L", objccodegenutils_zip_url, "-o", "package.zip")
	err = objccodegenutils_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	objccodegenutils_bin_url := "https://github.com/puls/objc-codegenutils/archive/refs/tags/v1.0.bin"
	objccodegenutils_cmd_bin := exec.Command("curl", "-L", objccodegenutils_bin_url, "-o", "binary.bin")
	err = objccodegenutils_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	objccodegenutils_src_url := "https://github.com/puls/objc-codegenutils/archive/refs/tags/v1.0.src.tar.gz"
	objccodegenutils_cmd_src := exec.Command("curl", "-L", objccodegenutils_src_url, "-o", "source.tar.gz")
	err = objccodegenutils_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	objccodegenutils_cmd_direct := exec.Command("./binary")
	err = objccodegenutils_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
