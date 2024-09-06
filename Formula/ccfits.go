package main

// Ccfits - Object oriented interface to the cfitsio library
// Homepage: https://heasarc.gsfc.nasa.gov/fitsio/CCfits/

import (
	"fmt"
	
	"os/exec"
)

func installCcfits() {
	// Método 1: Descargar y extraer .tar.gz
	ccfits_tar_url := "https://heasarc.gsfc.nasa.gov/fitsio/CCfits/CCfits-2.6.tar.gz"
	ccfits_cmd_tar := exec.Command("curl", "-L", ccfits_tar_url, "-o", "package.tar.gz")
	err := ccfits_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ccfits_zip_url := "https://heasarc.gsfc.nasa.gov/fitsio/CCfits/CCfits-2.6.zip"
	ccfits_cmd_zip := exec.Command("curl", "-L", ccfits_zip_url, "-o", "package.zip")
	err = ccfits_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ccfits_bin_url := "https://heasarc.gsfc.nasa.gov/fitsio/CCfits/CCfits-2.6.bin"
	ccfits_cmd_bin := exec.Command("curl", "-L", ccfits_bin_url, "-o", "binary.bin")
	err = ccfits_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ccfits_src_url := "https://heasarc.gsfc.nasa.gov/fitsio/CCfits/CCfits-2.6.src.tar.gz"
	ccfits_cmd_src := exec.Command("curl", "-L", ccfits_src_url, "-o", "source.tar.gz")
	err = ccfits_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ccfits_cmd_direct := exec.Command("./binary")
	err = ccfits_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cfitsio")
	exec.Command("latte", "install", "cfitsio").Run()
}
