package main

// GnuUnits - GNU unit conversion tool
// Homepage: https://www.gnu.org/software/units/

import (
	"fmt"
	
	"os/exec"
)

func installGnuUnits() {
	// Método 1: Descargar y extraer .tar.gz
	gnuunits_tar_url := "https://ftp.gnu.org/gnu/units/units-2.23.tar.gz"
	gnuunits_cmd_tar := exec.Command("curl", "-L", gnuunits_tar_url, "-o", "package.tar.gz")
	err := gnuunits_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gnuunits_zip_url := "https://ftp.gnu.org/gnu/units/units-2.23.zip"
	gnuunits_cmd_zip := exec.Command("curl", "-L", gnuunits_zip_url, "-o", "package.zip")
	err = gnuunits_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gnuunits_bin_url := "https://ftp.gnu.org/gnu/units/units-2.23.bin"
	gnuunits_cmd_bin := exec.Command("curl", "-L", gnuunits_bin_url, "-o", "binary.bin")
	err = gnuunits_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gnuunits_src_url := "https://ftp.gnu.org/gnu/units/units-2.23.src.tar.gz"
	gnuunits_cmd_src := exec.Command("curl", "-L", gnuunits_src_url, "-o", "source.tar.gz")
	err = gnuunits_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gnuunits_cmd_direct := exec.Command("./binary")
	err = gnuunits_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
}
