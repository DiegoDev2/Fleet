package main

// AmplMp - Open-source library for mathematical programming
// Homepage: https://www.ampl.com/

import (
	"fmt"
	
	"os/exec"
)

func installAmplMp() {
	// Método 1: Descargar y extraer .tar.gz
	amplmp_tar_url := "https://github.com/ampl/mp/archive/refs/tags/3.1.0.tar.gz"
	amplmp_cmd_tar := exec.Command("curl", "-L", amplmp_tar_url, "-o", "package.tar.gz")
	err := amplmp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	amplmp_zip_url := "https://github.com/ampl/mp/archive/refs/tags/3.1.0.zip"
	amplmp_cmd_zip := exec.Command("curl", "-L", amplmp_zip_url, "-o", "package.zip")
	err = amplmp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	amplmp_bin_url := "https://github.com/ampl/mp/archive/refs/tags/3.1.0.bin"
	amplmp_cmd_bin := exec.Command("curl", "-L", amplmp_bin_url, "-o", "binary.bin")
	err = amplmp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	amplmp_src_url := "https://github.com/ampl/mp/archive/refs/tags/3.1.0.src.tar.gz"
	amplmp_cmd_src := exec.Command("curl", "-L", amplmp_src_url, "-o", "source.tar.gz")
	err = amplmp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	amplmp_cmd_direct := exec.Command("./binary")
	err = amplmp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
