package main

// AstrometryNet - Automatic identification of astronomical images
// Homepage: https://github.com/dstndstn/astrometry.net

import (
	"fmt"
	
	"os/exec"
)

func installAstrometryNet() {
	// Método 1: Descargar y extraer .tar.gz
	astrometrynet_tar_url := "https://github.com/dstndstn/astrometry.net/releases/download/0.95/astrometry.net-0.95.tar.gz"
	astrometrynet_cmd_tar := exec.Command("curl", "-L", astrometrynet_tar_url, "-o", "package.tar.gz")
	err := astrometrynet_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	astrometrynet_zip_url := "https://github.com/dstndstn/astrometry.net/releases/download/0.95/astrometry.net-0.95.zip"
	astrometrynet_cmd_zip := exec.Command("curl", "-L", astrometrynet_zip_url, "-o", "package.zip")
	err = astrometrynet_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	astrometrynet_bin_url := "https://github.com/dstndstn/astrometry.net/releases/download/0.95/astrometry.net-0.95.bin"
	astrometrynet_cmd_bin := exec.Command("curl", "-L", astrometrynet_bin_url, "-o", "binary.bin")
	err = astrometrynet_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	astrometrynet_src_url := "https://github.com/dstndstn/astrometry.net/releases/download/0.95/astrometry.net-0.95.src.tar.gz"
	astrometrynet_cmd_src := exec.Command("curl", "-L", astrometrynet_src_url, "-o", "source.tar.gz")
	err = astrometrynet_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	astrometrynet_cmd_direct := exec.Command("./binary")
	err = astrometrynet_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: python-setuptools")
exec.Command("latte", "install", "python-setuptools")
	fmt.Println("Instalando dependencia: swig")
exec.Command("latte", "install", "swig")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: cfitsio")
exec.Command("latte", "install", "cfitsio")
	fmt.Println("Instalando dependencia: gsl")
exec.Command("latte", "install", "gsl")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: netpbm")
exec.Command("latte", "install", "netpbm")
	fmt.Println("Instalando dependencia: numpy")
exec.Command("latte", "install", "numpy")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: wcslib")
exec.Command("latte", "install", "wcslib")
}
