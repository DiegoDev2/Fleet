package main

// Austin - Python frame stack sampler for CPython
// Homepage: https://github.com/P403n1x87/austin

import (
	"fmt"
	
	"os/exec"
)

func installAustin() {
	// Método 1: Descargar y extraer .tar.gz
	austin_tar_url := "https://github.com/P403n1x87/austin/archive/refs/tags/v3.6.0.tar.gz"
	austin_cmd_tar := exec.Command("curl", "-L", austin_tar_url, "-o", "package.tar.gz")
	err := austin_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	austin_zip_url := "https://github.com/P403n1x87/austin/archive/refs/tags/v3.6.0.zip"
	austin_cmd_zip := exec.Command("curl", "-L", austin_zip_url, "-o", "package.zip")
	err = austin_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	austin_bin_url := "https://github.com/P403n1x87/austin/archive/refs/tags/v3.6.0.bin"
	austin_cmd_bin := exec.Command("curl", "-L", austin_bin_url, "-o", "binary.bin")
	err = austin_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	austin_src_url := "https://github.com/P403n1x87/austin/archive/refs/tags/v3.6.0.src.tar.gz"
	austin_cmd_src := exec.Command("curl", "-L", austin_src_url, "-o", "source.tar.gz")
	err = austin_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	austin_cmd_direct := exec.Command("./binary")
	err = austin_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
