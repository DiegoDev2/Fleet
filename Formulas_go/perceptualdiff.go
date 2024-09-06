package main

// Perceptualdiff - Perceptual image comparison tool
// Homepage: https://pdiff.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installPerceptualdiff() {
	// Método 1: Descargar y extraer .tar.gz
	perceptualdiff_tar_url := "https://downloads.sourceforge.net/project/pdiff/pdiff/perceptualdiff-1.1.1/perceptualdiff-1.1.1-src.tar.gz"
	perceptualdiff_cmd_tar := exec.Command("curl", "-L", perceptualdiff_tar_url, "-o", "package.tar.gz")
	err := perceptualdiff_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	perceptualdiff_zip_url := "https://downloads.sourceforge.net/project/pdiff/pdiff/perceptualdiff-1.1.1/perceptualdiff-1.1.1-src.zip"
	perceptualdiff_cmd_zip := exec.Command("curl", "-L", perceptualdiff_zip_url, "-o", "package.zip")
	err = perceptualdiff_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	perceptualdiff_bin_url := "https://downloads.sourceforge.net/project/pdiff/pdiff/perceptualdiff-1.1.1/perceptualdiff-1.1.1-src.bin"
	perceptualdiff_cmd_bin := exec.Command("curl", "-L", perceptualdiff_bin_url, "-o", "binary.bin")
	err = perceptualdiff_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	perceptualdiff_src_url := "https://downloads.sourceforge.net/project/pdiff/pdiff/perceptualdiff-1.1.1/perceptualdiff-1.1.1-src.src.tar.gz"
	perceptualdiff_cmd_src := exec.Command("curl", "-L", perceptualdiff_src_url, "-o", "source.tar.gz")
	err = perceptualdiff_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	perceptualdiff_cmd_direct := exec.Command("./binary")
	err = perceptualdiff_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: freeimage")
exec.Command("latte", "install", "freeimage")
}
