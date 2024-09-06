package main

// LiquidDsp - Digital signal processing library for software-defined radios
// Homepage: https://liquidsdr.org/

import (
	"fmt"
	
	"os/exec"
)

func installLiquidDsp() {
	// Método 1: Descargar y extraer .tar.gz
	liquiddsp_tar_url := "https://github.com/jgaeddert/liquid-dsp/archive/refs/tags/v1.6.0.tar.gz"
	liquiddsp_cmd_tar := exec.Command("curl", "-L", liquiddsp_tar_url, "-o", "package.tar.gz")
	err := liquiddsp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	liquiddsp_zip_url := "https://github.com/jgaeddert/liquid-dsp/archive/refs/tags/v1.6.0.zip"
	liquiddsp_cmd_zip := exec.Command("curl", "-L", liquiddsp_zip_url, "-o", "package.zip")
	err = liquiddsp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	liquiddsp_bin_url := "https://github.com/jgaeddert/liquid-dsp/archive/refs/tags/v1.6.0.bin"
	liquiddsp_cmd_bin := exec.Command("curl", "-L", liquiddsp_bin_url, "-o", "binary.bin")
	err = liquiddsp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	liquiddsp_src_url := "https://github.com/jgaeddert/liquid-dsp/archive/refs/tags/v1.6.0.src.tar.gz"
	liquiddsp_cmd_src := exec.Command("curl", "-L", liquiddsp_src_url, "-o", "source.tar.gz")
	err = liquiddsp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	liquiddsp_cmd_direct := exec.Command("./binary")
	err = liquiddsp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: fftw")
	exec.Command("latte", "install", "fftw").Run()
}
