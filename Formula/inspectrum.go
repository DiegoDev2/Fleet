package main

// Inspectrum - Offline radio signal analyser
// Homepage: https://github.com/miek/inspectrum

import (
	"fmt"
	
	"os/exec"
)

func installInspectrum() {
	// Método 1: Descargar y extraer .tar.gz
	inspectrum_tar_url := "https://github.com/miek/inspectrum/archive/refs/tags/v0.3.1.tar.gz"
	inspectrum_cmd_tar := exec.Command("curl", "-L", inspectrum_tar_url, "-o", "package.tar.gz")
	err := inspectrum_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	inspectrum_zip_url := "https://github.com/miek/inspectrum/archive/refs/tags/v0.3.1.zip"
	inspectrum_cmd_zip := exec.Command("curl", "-L", inspectrum_zip_url, "-o", "package.zip")
	err = inspectrum_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	inspectrum_bin_url := "https://github.com/miek/inspectrum/archive/refs/tags/v0.3.1.bin"
	inspectrum_cmd_bin := exec.Command("curl", "-L", inspectrum_bin_url, "-o", "binary.bin")
	err = inspectrum_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	inspectrum_src_url := "https://github.com/miek/inspectrum/archive/refs/tags/v0.3.1.src.tar.gz"
	inspectrum_cmd_src := exec.Command("curl", "-L", inspectrum_src_url, "-o", "source.tar.gz")
	err = inspectrum_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	inspectrum_cmd_direct := exec.Command("./binary")
	err = inspectrum_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: fftw")
	exec.Command("latte", "install", "fftw").Run()
	fmt.Println("Instalando dependencia: liquid-dsp")
	exec.Command("latte", "install", "liquid-dsp").Run()
	fmt.Println("Instalando dependencia: qt@5")
	exec.Command("latte", "install", "qt@5").Run()
}
