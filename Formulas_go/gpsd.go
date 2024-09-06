package main

// Gpsd - Global Positioning System (GPS) daemon
// Homepage: https://gpsd.gitlab.io/gpsd/

import (
	"fmt"
	
	"os/exec"
)

func installGpsd() {
	// Método 1: Descargar y extraer .tar.gz
	gpsd_tar_url := "https://download.savannah.gnu.org/releases/gpsd/gpsd-3.25.tar.xz"
	gpsd_cmd_tar := exec.Command("curl", "-L", gpsd_tar_url, "-o", "package.tar.gz")
	err := gpsd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gpsd_zip_url := "https://download.savannah.gnu.org/releases/gpsd/gpsd-3.25.tar.xz"
	gpsd_cmd_zip := exec.Command("curl", "-L", gpsd_zip_url, "-o", "package.zip")
	err = gpsd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gpsd_bin_url := "https://download.savannah.gnu.org/releases/gpsd/gpsd-3.25.tar.xz"
	gpsd_cmd_bin := exec.Command("curl", "-L", gpsd_bin_url, "-o", "binary.bin")
	err = gpsd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gpsd_src_url := "https://download.savannah.gnu.org/releases/gpsd/gpsd-3.25.tar.xz"
	gpsd_cmd_src := exec.Command("curl", "-L", gpsd_src_url, "-o", "source.tar.gz")
	err = gpsd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gpsd_cmd_direct := exec.Command("./binary")
	err = gpsd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: asciidoctor")
exec.Command("latte", "install", "asciidoctor")
	fmt.Println("Instalando dependencia: python-setuptools")
exec.Command("latte", "install", "python-setuptools")
	fmt.Println("Instalando dependencia: scons")
exec.Command("latte", "install", "scons")
}
