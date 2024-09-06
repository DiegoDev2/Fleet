package main

// Veclibfort - GNU Fortran compatibility for Apple's vecLib
// Homepage: https://github.com/mcg1969/vecLibFort

import (
	"fmt"
	
	"os/exec"
)

func installVeclibfort() {
	// Método 1: Descargar y extraer .tar.gz
	veclibfort_tar_url := "https://github.com/mcg1969/vecLibFort/archive/refs/tags/0.4.3.tar.gz"
	veclibfort_cmd_tar := exec.Command("curl", "-L", veclibfort_tar_url, "-o", "package.tar.gz")
	err := veclibfort_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	veclibfort_zip_url := "https://github.com/mcg1969/vecLibFort/archive/refs/tags/0.4.3.zip"
	veclibfort_cmd_zip := exec.Command("curl", "-L", veclibfort_zip_url, "-o", "package.zip")
	err = veclibfort_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	veclibfort_bin_url := "https://github.com/mcg1969/vecLibFort/archive/refs/tags/0.4.3.bin"
	veclibfort_cmd_bin := exec.Command("curl", "-L", veclibfort_bin_url, "-o", "binary.bin")
	err = veclibfort_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	veclibfort_src_url := "https://github.com/mcg1969/vecLibFort/archive/refs/tags/0.4.3.src.tar.gz"
	veclibfort_cmd_src := exec.Command("curl", "-L", veclibfort_src_url, "-o", "source.tar.gz")
	err = veclibfort_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	veclibfort_cmd_direct := exec.Command("./binary")
	err = veclibfort_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gcc")
	exec.Command("latte", "install", "gcc").Run()
}
