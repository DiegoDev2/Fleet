package main

// Limesuite - Device drivers utilities, and interface layers for LimeSDR
// Homepage: https://myriadrf.org/projects/software/lime-suite/

import (
	"fmt"
	
	"os/exec"
)

func installLimesuite() {
	// Método 1: Descargar y extraer .tar.gz
	limesuite_tar_url := "https://github.com/myriadrf/LimeSuite/archive/refs/tags/v23.11.0.tar.gz"
	limesuite_cmd_tar := exec.Command("curl", "-L", limesuite_tar_url, "-o", "package.tar.gz")
	err := limesuite_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	limesuite_zip_url := "https://github.com/myriadrf/LimeSuite/archive/refs/tags/v23.11.0.zip"
	limesuite_cmd_zip := exec.Command("curl", "-L", limesuite_zip_url, "-o", "package.zip")
	err = limesuite_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	limesuite_bin_url := "https://github.com/myriadrf/LimeSuite/archive/refs/tags/v23.11.0.bin"
	limesuite_cmd_bin := exec.Command("curl", "-L", limesuite_bin_url, "-o", "binary.bin")
	err = limesuite_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	limesuite_src_url := "https://github.com/myriadrf/LimeSuite/archive/refs/tags/v23.11.0.src.tar.gz"
	limesuite_cmd_src := exec.Command("curl", "-L", limesuite_src_url, "-o", "source.tar.gz")
	err = limesuite_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	limesuite_cmd_direct := exec.Command("./binary")
	err = limesuite_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: fltk")
	exec.Command("latte", "install", "fltk").Run()
	fmt.Println("Instalando dependencia: gnuplot")
	exec.Command("latte", "install", "gnuplot").Run()
	fmt.Println("Instalando dependencia: libusb")
	exec.Command("latte", "install", "libusb").Run()
	fmt.Println("Instalando dependencia: soapysdr")
	exec.Command("latte", "install", "soapysdr").Run()
}
