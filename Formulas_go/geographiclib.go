package main

// Geographiclib - C++ geography library
// Homepage: https://geographiclib.sourceforge.io/

import (
	"fmt"
	
	"os/exec"
)

func installGeographiclib() {
	// Método 1: Descargar y extraer .tar.gz
	geographiclib_tar_url := "https://github.com/geographiclib/geographiclib/archive/refs/tags/r2.4.tar.gz"
	geographiclib_cmd_tar := exec.Command("curl", "-L", geographiclib_tar_url, "-o", "package.tar.gz")
	err := geographiclib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	geographiclib_zip_url := "https://github.com/geographiclib/geographiclib/archive/refs/tags/r2.4.zip"
	geographiclib_cmd_zip := exec.Command("curl", "-L", geographiclib_zip_url, "-o", "package.zip")
	err = geographiclib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	geographiclib_bin_url := "https://github.com/geographiclib/geographiclib/archive/refs/tags/r2.4.bin"
	geographiclib_cmd_bin := exec.Command("curl", "-L", geographiclib_bin_url, "-o", "binary.bin")
	err = geographiclib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	geographiclib_src_url := "https://github.com/geographiclib/geographiclib/archive/refs/tags/r2.4.src.tar.gz"
	geographiclib_cmd_src := exec.Command("curl", "-L", geographiclib_src_url, "-o", "source.tar.gz")
	err = geographiclib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	geographiclib_cmd_direct := exec.Command("./binary")
	err = geographiclib_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
