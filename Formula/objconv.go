package main

// Objconv - Object file converter
// Homepage: https://www.agner.org/optimize/#objconv

import (
	"fmt"
	
	"os/exec"
)

func installObjconv() {
	// Método 1: Descargar y extraer .tar.gz
	objconv_tar_url := "https://www.agner.org/optimize/objconv.zip"
	objconv_cmd_tar := exec.Command("curl", "-L", objconv_tar_url, "-o", "package.tar.gz")
	err := objconv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	objconv_zip_url := "https://www.agner.org/optimize/objconv.zip"
	objconv_cmd_zip := exec.Command("curl", "-L", objconv_zip_url, "-o", "package.zip")
	err = objconv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	objconv_bin_url := "https://www.agner.org/optimize/objconv.zip"
	objconv_cmd_bin := exec.Command("curl", "-L", objconv_bin_url, "-o", "binary.bin")
	err = objconv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	objconv_src_url := "https://www.agner.org/optimize/objconv.zip"
	objconv_cmd_src := exec.Command("curl", "-L", objconv_src_url, "-o", "source.tar.gz")
	err = objconv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	objconv_cmd_direct := exec.Command("./binary")
	err = objconv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
