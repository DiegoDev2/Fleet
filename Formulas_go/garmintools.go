package main

// Garmintools - Interface to the Garmin Forerunner GPS units
// Homepage: https://code.google.com/archive/p/garmintools/

import (
	"fmt"
	
	"os/exec"
)

func installGarmintools() {
	// Método 1: Descargar y extraer .tar.gz
	garmintools_tar_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/garmintools/garmintools-0.10.tar.gz"
	garmintools_cmd_tar := exec.Command("curl", "-L", garmintools_tar_url, "-o", "package.tar.gz")
	err := garmintools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	garmintools_zip_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/garmintools/garmintools-0.10.zip"
	garmintools_cmd_zip := exec.Command("curl", "-L", garmintools_zip_url, "-o", "package.zip")
	err = garmintools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	garmintools_bin_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/garmintools/garmintools-0.10.bin"
	garmintools_cmd_bin := exec.Command("curl", "-L", garmintools_bin_url, "-o", "binary.bin")
	err = garmintools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	garmintools_src_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/garmintools/garmintools-0.10.src.tar.gz"
	garmintools_cmd_src := exec.Command("curl", "-L", garmintools_src_url, "-o", "source.tar.gz")
	err = garmintools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	garmintools_cmd_direct := exec.Command("./binary")
	err = garmintools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libusb-compat")
exec.Command("latte", "install", "libusb-compat")
}
