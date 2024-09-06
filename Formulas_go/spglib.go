package main

// Spglib - C library for finding and handling crystal symmetries
// Homepage: https://spglib.readthedocs.io/

import (
	"fmt"
	
	"os/exec"
)

func installSpglib() {
	// Método 1: Descargar y extraer .tar.gz
	spglib_tar_url := "https://github.com/spglib/spglib/archive/refs/tags/v2.5.0.tar.gz"
	spglib_cmd_tar := exec.Command("curl", "-L", spglib_tar_url, "-o", "package.tar.gz")
	err := spglib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	spglib_zip_url := "https://github.com/spglib/spglib/archive/refs/tags/v2.5.0.zip"
	spglib_cmd_zip := exec.Command("curl", "-L", spglib_zip_url, "-o", "package.zip")
	err = spglib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	spglib_bin_url := "https://github.com/spglib/spglib/archive/refs/tags/v2.5.0.bin"
	spglib_cmd_bin := exec.Command("curl", "-L", spglib_bin_url, "-o", "binary.bin")
	err = spglib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	spglib_src_url := "https://github.com/spglib/spglib/archive/refs/tags/v2.5.0.src.tar.gz"
	spglib_cmd_src := exec.Command("curl", "-L", spglib_src_url, "-o", "source.tar.gz")
	err = spglib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	spglib_cmd_direct := exec.Command("./binary")
	err = spglib_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: gcc")
exec.Command("latte", "install", "gcc")
}
