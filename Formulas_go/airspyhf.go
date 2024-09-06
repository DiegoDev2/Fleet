package main

// Airspyhf - Driver and tools for a software-defined radio
// Homepage: https://airspy.com/

import (
	"fmt"
	
	"os/exec"
)

func installAirspyhf() {
	// Método 1: Descargar y extraer .tar.gz
	airspyhf_tar_url := "https://github.com/airspy/airspyhf/archive/refs/tags/1.6.8.tar.gz"
	airspyhf_cmd_tar := exec.Command("curl", "-L", airspyhf_tar_url, "-o", "package.tar.gz")
	err := airspyhf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	airspyhf_zip_url := "https://github.com/airspy/airspyhf/archive/refs/tags/1.6.8.zip"
	airspyhf_cmd_zip := exec.Command("curl", "-L", airspyhf_zip_url, "-o", "package.zip")
	err = airspyhf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	airspyhf_bin_url := "https://github.com/airspy/airspyhf/archive/refs/tags/1.6.8.bin"
	airspyhf_cmd_bin := exec.Command("curl", "-L", airspyhf_bin_url, "-o", "binary.bin")
	err = airspyhf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	airspyhf_src_url := "https://github.com/airspy/airspyhf/archive/refs/tags/1.6.8.src.tar.gz"
	airspyhf_cmd_src := exec.Command("curl", "-L", airspyhf_src_url, "-o", "source.tar.gz")
	err = airspyhf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	airspyhf_cmd_direct := exec.Command("./binary")
	err = airspyhf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libusb")
exec.Command("latte", "install", "libusb")
}
