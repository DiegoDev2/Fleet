package main

// Soapysdr - Vendor and platform neutral SDR support library
// Homepage: https://github.com/pothosware/SoapySDR/wiki

import (
	"fmt"
	
	"os/exec"
)

func installSoapysdr() {
	// Método 1: Descargar y extraer .tar.gz
	soapysdr_tar_url := "https://github.com/pothosware/SoapySDR/archive/refs/tags/soapy-sdr-0.8.1.tar.gz"
	soapysdr_cmd_tar := exec.Command("curl", "-L", soapysdr_tar_url, "-o", "package.tar.gz")
	err := soapysdr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	soapysdr_zip_url := "https://github.com/pothosware/SoapySDR/archive/refs/tags/soapy-sdr-0.8.1.zip"
	soapysdr_cmd_zip := exec.Command("curl", "-L", soapysdr_zip_url, "-o", "package.zip")
	err = soapysdr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	soapysdr_bin_url := "https://github.com/pothosware/SoapySDR/archive/refs/tags/soapy-sdr-0.8.1.bin"
	soapysdr_cmd_bin := exec.Command("curl", "-L", soapysdr_bin_url, "-o", "binary.bin")
	err = soapysdr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	soapysdr_src_url := "https://github.com/pothosware/SoapySDR/archive/refs/tags/soapy-sdr-0.8.1.src.tar.gz"
	soapysdr_cmd_src := exec.Command("curl", "-L", soapysdr_src_url, "-o", "source.tar.gz")
	err = soapysdr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	soapysdr_cmd_direct := exec.Command("./binary")
	err = soapysdr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: python-setuptools")
exec.Command("latte", "install", "python-setuptools")
	fmt.Println("Instalando dependencia: swig")
exec.Command("latte", "install", "swig")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
