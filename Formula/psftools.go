package main

// Psftools - Tools for fixed-width bitmap fonts
// Homepage: https://www.seasip.info/Unix/PSF/

import (
	"fmt"
	
	"os/exec"
)

func installPsftools() {
	// Método 1: Descargar y extraer .tar.gz
	psftools_tar_url := "https://www.seasip.info/Unix/PSF/psftools-1.0.14.tar.gz"
	psftools_cmd_tar := exec.Command("curl", "-L", psftools_tar_url, "-o", "package.tar.gz")
	err := psftools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	psftools_zip_url := "https://www.seasip.info/Unix/PSF/psftools-1.0.14.zip"
	psftools_cmd_zip := exec.Command("curl", "-L", psftools_zip_url, "-o", "package.zip")
	err = psftools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	psftools_bin_url := "https://www.seasip.info/Unix/PSF/psftools-1.0.14.bin"
	psftools_cmd_bin := exec.Command("curl", "-L", psftools_bin_url, "-o", "binary.bin")
	err = psftools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	psftools_src_url := "https://www.seasip.info/Unix/PSF/psftools-1.0.14.src.tar.gz"
	psftools_cmd_src := exec.Command("curl", "-L", psftools_src_url, "-o", "source.tar.gz")
	err = psftools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	psftools_cmd_direct := exec.Command("./binary")
	err = psftools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
}
