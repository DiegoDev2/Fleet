package main

// Udunits - Unidata unit conversion library
// Homepage: https://www.unidata.ucar.edu/software/udunits/

import (
	"fmt"
	
	"os/exec"
)

func installUdunits() {
	// Método 1: Descargar y extraer .tar.gz
	udunits_tar_url := "https://artifacts.unidata.ucar.edu/repository/downloads-udunits/2.2.28/udunits-2.2.28.tar.gz"
	udunits_cmd_tar := exec.Command("curl", "-L", udunits_tar_url, "-o", "package.tar.gz")
	err := udunits_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	udunits_zip_url := "https://artifacts.unidata.ucar.edu/repository/downloads-udunits/2.2.28/udunits-2.2.28.zip"
	udunits_cmd_zip := exec.Command("curl", "-L", udunits_zip_url, "-o", "package.zip")
	err = udunits_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	udunits_bin_url := "https://artifacts.unidata.ucar.edu/repository/downloads-udunits/2.2.28/udunits-2.2.28.bin"
	udunits_cmd_bin := exec.Command("curl", "-L", udunits_bin_url, "-o", "binary.bin")
	err = udunits_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	udunits_src_url := "https://artifacts.unidata.ucar.edu/repository/downloads-udunits/2.2.28/udunits-2.2.28.src.tar.gz"
	udunits_cmd_src := exec.Command("curl", "-L", udunits_src_url, "-o", "source.tar.gz")
	err = udunits_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	udunits_cmd_direct := exec.Command("./binary")
	err = udunits_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: texinfo")
exec.Command("latte", "install", "texinfo")
}
