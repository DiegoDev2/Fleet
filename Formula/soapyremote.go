package main

// Soapyremote - Use any Soapy SDR remotely
// Homepage: https://github.com/pothosware/SoapyRemote/wiki

import (
	"fmt"
	
	"os/exec"
)

func installSoapyremote() {
	// Método 1: Descargar y extraer .tar.gz
	soapyremote_tar_url := "https://github.com/pothosware/SoapyRemote/archive/refs/tags/soapy-remote-0.5.2.tar.gz"
	soapyremote_cmd_tar := exec.Command("curl", "-L", soapyremote_tar_url, "-o", "package.tar.gz")
	err := soapyremote_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	soapyremote_zip_url := "https://github.com/pothosware/SoapyRemote/archive/refs/tags/soapy-remote-0.5.2.zip"
	soapyremote_cmd_zip := exec.Command("curl", "-L", soapyremote_zip_url, "-o", "package.zip")
	err = soapyremote_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	soapyremote_bin_url := "https://github.com/pothosware/SoapyRemote/archive/refs/tags/soapy-remote-0.5.2.bin"
	soapyremote_cmd_bin := exec.Command("curl", "-L", soapyremote_bin_url, "-o", "binary.bin")
	err = soapyremote_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	soapyremote_src_url := "https://github.com/pothosware/SoapyRemote/archive/refs/tags/soapy-remote-0.5.2.src.tar.gz"
	soapyremote_cmd_src := exec.Command("curl", "-L", soapyremote_src_url, "-o", "source.tar.gz")
	err = soapyremote_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	soapyremote_cmd_direct := exec.Command("./binary")
	err = soapyremote_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: soapysdr")
	exec.Command("latte", "install", "soapysdr").Run()
}
