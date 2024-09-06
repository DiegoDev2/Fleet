package main

// Gpsim - Simulator for Microchip's PIC microcontrollers
// Homepage: https://gpsim.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installGpsim() {
	// Método 1: Descargar y extraer .tar.gz
	gpsim_tar_url := "https://downloads.sourceforge.net/project/gpsim/gpsim/0.32.0/gpsim-0.32.1.tar.gz"
	gpsim_cmd_tar := exec.Command("curl", "-L", gpsim_tar_url, "-o", "package.tar.gz")
	err := gpsim_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gpsim_zip_url := "https://downloads.sourceforge.net/project/gpsim/gpsim/0.32.0/gpsim-0.32.1.zip"
	gpsim_cmd_zip := exec.Command("curl", "-L", gpsim_zip_url, "-o", "package.zip")
	err = gpsim_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gpsim_bin_url := "https://downloads.sourceforge.net/project/gpsim/gpsim/0.32.0/gpsim-0.32.1.bin"
	gpsim_cmd_bin := exec.Command("curl", "-L", gpsim_bin_url, "-o", "binary.bin")
	err = gpsim_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gpsim_src_url := "https://downloads.sourceforge.net/project/gpsim/gpsim/0.32.0/gpsim-0.32.1.src.tar.gz"
	gpsim_cmd_src := exec.Command("curl", "-L", gpsim_src_url, "-o", "source.tar.gz")
	err = gpsim_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gpsim_cmd_direct := exec.Command("./binary")
	err = gpsim_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gputils")
exec.Command("latte", "install", "gputils")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: popt")
exec.Command("latte", "install", "popt")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
