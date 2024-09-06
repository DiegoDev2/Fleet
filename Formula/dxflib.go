package main

// Dxflib - C++ library for parsing DXF files
// Homepage: https://www.ribbonsoft.com/en/what-is-dxflib

import (
	"fmt"
	
	"os/exec"
)

func installDxflib() {
	// Método 1: Descargar y extraer .tar.gz
	dxflib_tar_url := "https://www.ribbonsoft.com/archives/dxflib/dxflib-3.26.4-src.tar.gz"
	dxflib_cmd_tar := exec.Command("curl", "-L", dxflib_tar_url, "-o", "package.tar.gz")
	err := dxflib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dxflib_zip_url := "https://www.ribbonsoft.com/archives/dxflib/dxflib-3.26.4-src.zip"
	dxflib_cmd_zip := exec.Command("curl", "-L", dxflib_zip_url, "-o", "package.zip")
	err = dxflib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dxflib_bin_url := "https://www.ribbonsoft.com/archives/dxflib/dxflib-3.26.4-src.bin"
	dxflib_cmd_bin := exec.Command("curl", "-L", dxflib_bin_url, "-o", "binary.bin")
	err = dxflib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dxflib_src_url := "https://www.ribbonsoft.com/archives/dxflib/dxflib-3.26.4-src.src.tar.gz"
	dxflib_cmd_src := exec.Command("curl", "-L", dxflib_src_url, "-o", "source.tar.gz")
	err = dxflib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dxflib_cmd_direct := exec.Command("./binary")
	err = dxflib_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: qt")
	exec.Command("latte", "install", "qt").Run()
}
