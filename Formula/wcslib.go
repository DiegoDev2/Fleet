package main

// Wcslib - Library and utilities for the FITS World Coordinate System
// Homepage: https://www.atnf.csiro.au/people/mcalabre/WCS/

import (
	"fmt"
	
	"os/exec"
)

func installWcslib() {
	// Método 1: Descargar y extraer .tar.gz
	wcslib_tar_url := "https://www.atnf.csiro.au/people/mcalabre/WCS/wcslib-8.2.2.tar.bz2"
	wcslib_cmd_tar := exec.Command("curl", "-L", wcslib_tar_url, "-o", "package.tar.gz")
	err := wcslib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wcslib_zip_url := "https://www.atnf.csiro.au/people/mcalabre/WCS/wcslib-8.2.2.tar.bz2"
	wcslib_cmd_zip := exec.Command("curl", "-L", wcslib_zip_url, "-o", "package.zip")
	err = wcslib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wcslib_bin_url := "https://www.atnf.csiro.au/people/mcalabre/WCS/wcslib-8.2.2.tar.bz2"
	wcslib_cmd_bin := exec.Command("curl", "-L", wcslib_bin_url, "-o", "binary.bin")
	err = wcslib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wcslib_src_url := "https://www.atnf.csiro.au/people/mcalabre/WCS/wcslib-8.2.2.tar.bz2"
	wcslib_cmd_src := exec.Command("curl", "-L", wcslib_src_url, "-o", "source.tar.gz")
	err = wcslib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wcslib_cmd_direct := exec.Command("./binary")
	err = wcslib_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cfitsio")
	exec.Command("latte", "install", "cfitsio").Run()
}
