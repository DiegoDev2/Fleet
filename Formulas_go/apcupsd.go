package main

// Apcupsd - Daemon for controlling APC UPSes
// Homepage: http://www.apcupsd.org

import (
	"fmt"
	
	"os/exec"
)

func installApcupsd() {
	// Método 1: Descargar y extraer .tar.gz
	apcupsd_tar_url := "https://downloads.sourceforge.net/project/apcupsd/apcupsd%20-%20Stable/3.14.14/apcupsd-3.14.14.tar.gz"
	apcupsd_cmd_tar := exec.Command("curl", "-L", apcupsd_tar_url, "-o", "package.tar.gz")
	err := apcupsd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	apcupsd_zip_url := "https://downloads.sourceforge.net/project/apcupsd/apcupsd%20-%20Stable/3.14.14/apcupsd-3.14.14.zip"
	apcupsd_cmd_zip := exec.Command("curl", "-L", apcupsd_zip_url, "-o", "package.zip")
	err = apcupsd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	apcupsd_bin_url := "https://downloads.sourceforge.net/project/apcupsd/apcupsd%20-%20Stable/3.14.14/apcupsd-3.14.14.bin"
	apcupsd_cmd_bin := exec.Command("curl", "-L", apcupsd_bin_url, "-o", "binary.bin")
	err = apcupsd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	apcupsd_src_url := "https://downloads.sourceforge.net/project/apcupsd/apcupsd%20-%20Stable/3.14.14/apcupsd-3.14.14.src.tar.gz"
	apcupsd_cmd_src := exec.Command("curl", "-L", apcupsd_src_url, "-o", "source.tar.gz")
	err = apcupsd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	apcupsd_cmd_direct := exec.Command("./binary")
	err = apcupsd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gd")
exec.Command("latte", "install", "gd")
	fmt.Println("Instalando dependencia: libusb-compat")
exec.Command("latte", "install", "libusb-compat")
}
