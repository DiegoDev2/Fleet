package main

// Chmlib - Library for dealing with Microsoft ITSS/CHM files
// Homepage: http://www.jedrea.com/chmlib/

import (
	"fmt"
	
	"os/exec"
)

func installChmlib() {
	// Método 1: Descargar y extraer .tar.gz
	chmlib_tar_url := "http://www.jedrea.com/chmlib/chmlib-0.40.tar.gz"
	chmlib_cmd_tar := exec.Command("curl", "-L", chmlib_tar_url, "-o", "package.tar.gz")
	err := chmlib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	chmlib_zip_url := "http://www.jedrea.com/chmlib/chmlib-0.40.zip"
	chmlib_cmd_zip := exec.Command("curl", "-L", chmlib_zip_url, "-o", "package.zip")
	err = chmlib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	chmlib_bin_url := "http://www.jedrea.com/chmlib/chmlib-0.40.bin"
	chmlib_cmd_bin := exec.Command("curl", "-L", chmlib_bin_url, "-o", "binary.bin")
	err = chmlib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	chmlib_src_url := "http://www.jedrea.com/chmlib/chmlib-0.40.src.tar.gz"
	chmlib_cmd_src := exec.Command("curl", "-L", chmlib_src_url, "-o", "source.tar.gz")
	err = chmlib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	chmlib_cmd_direct := exec.Command("./binary")
	err = chmlib_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
