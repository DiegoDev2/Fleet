package main

// Iir1 - DSP IIR realtime filter library written in C++
// Homepage: https://berndporr.github.io/iir1/

import (
	"fmt"
	
	"os/exec"
)

func installIir1() {
	// Método 1: Descargar y extraer .tar.gz
	iir1_tar_url := "https://github.com/berndporr/iir1/archive/refs/tags/1.9.5.tar.gz"
	iir1_cmd_tar := exec.Command("curl", "-L", iir1_tar_url, "-o", "package.tar.gz")
	err := iir1_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	iir1_zip_url := "https://github.com/berndporr/iir1/archive/refs/tags/1.9.5.zip"
	iir1_cmd_zip := exec.Command("curl", "-L", iir1_zip_url, "-o", "package.zip")
	err = iir1_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	iir1_bin_url := "https://github.com/berndporr/iir1/archive/refs/tags/1.9.5.bin"
	iir1_cmd_bin := exec.Command("curl", "-L", iir1_bin_url, "-o", "binary.bin")
	err = iir1_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	iir1_src_url := "https://github.com/berndporr/iir1/archive/refs/tags/1.9.5.src.tar.gz"
	iir1_cmd_src := exec.Command("curl", "-L", iir1_src_url, "-o", "source.tar.gz")
	err = iir1_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	iir1_cmd_direct := exec.Command("./binary")
	err = iir1_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
