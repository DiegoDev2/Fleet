package main

// Soapyrtlsdr - SoapySDR RTL-SDR Support Module
// Homepage: https://github.com/pothosware/SoapyRTLSDR/wiki

import (
	"fmt"
	
	"os/exec"
)

func installSoapyrtlsdr() {
	// Método 1: Descargar y extraer .tar.gz
	soapyrtlsdr_tar_url := "https://github.com/pothosware/SoapyRTLSDR/archive/refs/tags/soapy-rtl-sdr-0.3.3.tar.gz"
	soapyrtlsdr_cmd_tar := exec.Command("curl", "-L", soapyrtlsdr_tar_url, "-o", "package.tar.gz")
	err := soapyrtlsdr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	soapyrtlsdr_zip_url := "https://github.com/pothosware/SoapyRTLSDR/archive/refs/tags/soapy-rtl-sdr-0.3.3.zip"
	soapyrtlsdr_cmd_zip := exec.Command("curl", "-L", soapyrtlsdr_zip_url, "-o", "package.zip")
	err = soapyrtlsdr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	soapyrtlsdr_bin_url := "https://github.com/pothosware/SoapyRTLSDR/archive/refs/tags/soapy-rtl-sdr-0.3.3.bin"
	soapyrtlsdr_cmd_bin := exec.Command("curl", "-L", soapyrtlsdr_bin_url, "-o", "binary.bin")
	err = soapyrtlsdr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	soapyrtlsdr_src_url := "https://github.com/pothosware/SoapyRTLSDR/archive/refs/tags/soapy-rtl-sdr-0.3.3.src.tar.gz"
	soapyrtlsdr_cmd_src := exec.Command("curl", "-L", soapyrtlsdr_src_url, "-o", "source.tar.gz")
	err = soapyrtlsdr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	soapyrtlsdr_cmd_direct := exec.Command("./binary")
	err = soapyrtlsdr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: librtlsdr")
	exec.Command("latte", "install", "librtlsdr").Run()
	fmt.Println("Instalando dependencia: soapysdr")
	exec.Command("latte", "install", "soapysdr").Run()
}
