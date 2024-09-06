package main

// Cspice - Observation geometry system for robotic space science missions
// Homepage: https://naif.jpl.nasa.gov/naif/toolkit.html

import (
	"fmt"
	
	"os/exec"
)

func installCspice() {
	// Método 1: Descargar y extraer .tar.gz
	cspice_tar_url := "https://naif.jpl.nasa.gov/pub/naif/toolkit/C/MacIntel_OSX_AppleC_64bit/packages/cspice.tar.Z"
	cspice_cmd_tar := exec.Command("curl", "-L", cspice_tar_url, "-o", "package.tar.gz")
	err := cspice_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cspice_zip_url := "https://naif.jpl.nasa.gov/pub/naif/toolkit/C/MacIntel_OSX_AppleC_64bit/packages/cspice.tar.Z"
	cspice_cmd_zip := exec.Command("curl", "-L", cspice_zip_url, "-o", "package.zip")
	err = cspice_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cspice_bin_url := "https://naif.jpl.nasa.gov/pub/naif/toolkit/C/MacIntel_OSX_AppleC_64bit/packages/cspice.tar.Z"
	cspice_cmd_bin := exec.Command("curl", "-L", cspice_bin_url, "-o", "binary.bin")
	err = cspice_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cspice_src_url := "https://naif.jpl.nasa.gov/pub/naif/toolkit/C/MacIntel_OSX_AppleC_64bit/packages/cspice.tar.Z"
	cspice_cmd_src := exec.Command("curl", "-L", cspice_src_url, "-o", "source.tar.gz")
	err = cspice_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cspice_cmd_direct := exec.Command("./binary")
	err = cspice_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: tcsh")
exec.Command("latte", "install", "tcsh")
}
