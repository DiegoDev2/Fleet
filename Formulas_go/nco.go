package main

// Nco - Command-line operators for netCDF and HDF files
// Homepage: https://nco.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installNco() {
	// Método 1: Descargar y extraer .tar.gz
	nco_tar_url := "https://github.com/nco/nco/archive/refs/tags/5.2.8.tar.gz"
	nco_cmd_tar := exec.Command("curl", "-L", nco_tar_url, "-o", "package.tar.gz")
	err := nco_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nco_zip_url := "https://github.com/nco/nco/archive/refs/tags/5.2.8.zip"
	nco_cmd_zip := exec.Command("curl", "-L", nco_zip_url, "-o", "package.zip")
	err = nco_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nco_bin_url := "https://github.com/nco/nco/archive/refs/tags/5.2.8.bin"
	nco_cmd_bin := exec.Command("curl", "-L", nco_bin_url, "-o", "binary.bin")
	err = nco_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nco_src_url := "https://github.com/nco/nco/archive/refs/tags/5.2.8.src.tar.gz"
	nco_cmd_src := exec.Command("curl", "-L", nco_src_url, "-o", "source.tar.gz")
	err = nco_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nco_cmd_direct := exec.Command("./binary")
	err = nco_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: gsl")
exec.Command("latte", "install", "gsl")
	fmt.Println("Instalando dependencia: netcdf")
exec.Command("latte", "install", "netcdf")
	fmt.Println("Instalando dependencia: texinfo")
exec.Command("latte", "install", "texinfo")
	fmt.Println("Instalando dependencia: udunits")
exec.Command("latte", "install", "udunits")
}
