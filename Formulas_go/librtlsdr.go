package main

// Librtlsdr - Use Realtek DVB-T dongles as a cheap SDR
// Homepage: https://osmocom.org/projects/rtl-sdr/wiki

import (
	"fmt"
	
	"os/exec"
)

func installLibrtlsdr() {
	// Método 1: Descargar y extraer .tar.gz
	librtlsdr_tar_url := "https://github.com/steve-m/librtlsdr/archive/refs/tags/v2.0.2.tar.gz"
	librtlsdr_cmd_tar := exec.Command("curl", "-L", librtlsdr_tar_url, "-o", "package.tar.gz")
	err := librtlsdr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	librtlsdr_zip_url := "https://github.com/steve-m/librtlsdr/archive/refs/tags/v2.0.2.zip"
	librtlsdr_cmd_zip := exec.Command("curl", "-L", librtlsdr_zip_url, "-o", "package.zip")
	err = librtlsdr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	librtlsdr_bin_url := "https://github.com/steve-m/librtlsdr/archive/refs/tags/v2.0.2.bin"
	librtlsdr_cmd_bin := exec.Command("curl", "-L", librtlsdr_bin_url, "-o", "binary.bin")
	err = librtlsdr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	librtlsdr_src_url := "https://github.com/steve-m/librtlsdr/archive/refs/tags/v2.0.2.src.tar.gz"
	librtlsdr_cmd_src := exec.Command("curl", "-L", librtlsdr_src_url, "-o", "source.tar.gz")
	err = librtlsdr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	librtlsdr_cmd_direct := exec.Command("./binary")
	err = librtlsdr_cmd_direct.Run()
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
