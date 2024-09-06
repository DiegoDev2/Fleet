package main

// Cfitsio - C access to FITS data files with optional Fortran wrappers
// Homepage: https://heasarc.gsfc.nasa.gov/docs/software/fitsio/fitsio.html

import (
	"fmt"
	
	"os/exec"
)

func installCfitsio() {
	// Método 1: Descargar y extraer .tar.gz
	cfitsio_tar_url := "https://heasarc.gsfc.nasa.gov/FTP/software/fitsio/c/cfitsio-4.4.1.tar.gz"
	cfitsio_cmd_tar := exec.Command("curl", "-L", cfitsio_tar_url, "-o", "package.tar.gz")
	err := cfitsio_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cfitsio_zip_url := "https://heasarc.gsfc.nasa.gov/FTP/software/fitsio/c/cfitsio-4.4.1.zip"
	cfitsio_cmd_zip := exec.Command("curl", "-L", cfitsio_zip_url, "-o", "package.zip")
	err = cfitsio_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cfitsio_bin_url := "https://heasarc.gsfc.nasa.gov/FTP/software/fitsio/c/cfitsio-4.4.1.bin"
	cfitsio_cmd_bin := exec.Command("curl", "-L", cfitsio_bin_url, "-o", "binary.bin")
	err = cfitsio_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cfitsio_src_url := "https://heasarc.gsfc.nasa.gov/FTP/software/fitsio/c/cfitsio-4.4.1.src.tar.gz"
	cfitsio_cmd_src := exec.Command("curl", "-L", cfitsio_src_url, "-o", "source.tar.gz")
	err = cfitsio_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cfitsio_cmd_direct := exec.Command("./binary")
	err = cfitsio_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
