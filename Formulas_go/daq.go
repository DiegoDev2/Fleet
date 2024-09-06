package main

// Daq - Network intrusion prevention and detection system
// Homepage: https://www.snort.org/

import (
	"fmt"
	
	"os/exec"
)

func installDaq() {
	// Método 1: Descargar y extraer .tar.gz
	daq_tar_url := "https://github.com/snort3/libdaq/archive/refs/tags/v3.0.16.tar.gz"
	daq_cmd_tar := exec.Command("curl", "-L", daq_tar_url, "-o", "package.tar.gz")
	err := daq_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	daq_zip_url := "https://github.com/snort3/libdaq/archive/refs/tags/v3.0.16.zip"
	daq_cmd_zip := exec.Command("curl", "-L", daq_zip_url, "-o", "package.zip")
	err = daq_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	daq_bin_url := "https://github.com/snort3/libdaq/archive/refs/tags/v3.0.16.bin"
	daq_cmd_bin := exec.Command("curl", "-L", daq_bin_url, "-o", "binary.bin")
	err = daq_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	daq_src_url := "https://github.com/snort3/libdaq/archive/refs/tags/v3.0.16.src.tar.gz"
	daq_cmd_src := exec.Command("curl", "-L", daq_src_url, "-o", "source.tar.gz")
	err = daq_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	daq_cmd_direct := exec.Command("./binary")
	err = daq_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
