package main

// Sextractor - Extract catalogs of sources from astronomical images
// Homepage: https://github.com/astromatic/sextractor

import (
	"fmt"
	
	"os/exec"
)

func installSextractor() {
	// Método 1: Descargar y extraer .tar.gz
	sextractor_tar_url := "https://github.com/astromatic/sextractor/archive/refs/tags/2.28.0.tar.gz"
	sextractor_cmd_tar := exec.Command("curl", "-L", sextractor_tar_url, "-o", "package.tar.gz")
	err := sextractor_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sextractor_zip_url := "https://github.com/astromatic/sextractor/archive/refs/tags/2.28.0.zip"
	sextractor_cmd_zip := exec.Command("curl", "-L", sextractor_zip_url, "-o", "package.zip")
	err = sextractor_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sextractor_bin_url := "https://github.com/astromatic/sextractor/archive/refs/tags/2.28.0.bin"
	sextractor_cmd_bin := exec.Command("curl", "-L", sextractor_bin_url, "-o", "binary.bin")
	err = sextractor_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sextractor_src_url := "https://github.com/astromatic/sextractor/archive/refs/tags/2.28.0.src.tar.gz"
	sextractor_cmd_src := exec.Command("curl", "-L", sextractor_src_url, "-o", "source.tar.gz")
	err = sextractor_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sextractor_cmd_direct := exec.Command("./binary")
	err = sextractor_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: cfitsio")
exec.Command("latte", "install", "cfitsio")
	fmt.Println("Instalando dependencia: fftw")
exec.Command("latte", "install", "fftw")
	fmt.Println("Instalando dependencia: openblas")
exec.Command("latte", "install", "openblas")
}
