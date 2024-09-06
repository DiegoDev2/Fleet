package main

// Asymptote - Powerful descriptive vector graphics language
// Homepage: https://asymptote.sourceforge.io

import (
	"fmt"
	
	"os/exec"
)

func installAsymptote() {
	// Método 1: Descargar y extraer .tar.gz
	asymptote_tar_url := "https://downloads.sourceforge.net/project/asymptote/2.91/asymptote-2.91.src.tgz"
	asymptote_cmd_tar := exec.Command("curl", "-L", asymptote_tar_url, "-o", "package.tar.gz")
	err := asymptote_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	asymptote_zip_url := "https://downloads.sourceforge.net/project/asymptote/2.91/asymptote-2.91.src.tgz"
	asymptote_cmd_zip := exec.Command("curl", "-L", asymptote_zip_url, "-o", "package.zip")
	err = asymptote_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	asymptote_bin_url := "https://downloads.sourceforge.net/project/asymptote/2.91/asymptote-2.91.src.tgz"
	asymptote_cmd_bin := exec.Command("curl", "-L", asymptote_bin_url, "-o", "binary.bin")
	err = asymptote_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	asymptote_src_url := "https://downloads.sourceforge.net/project/asymptote/2.91/asymptote-2.91.src.tgz"
	asymptote_cmd_src := exec.Command("curl", "-L", asymptote_src_url, "-o", "source.tar.gz")
	err = asymptote_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	asymptote_cmd_direct := exec.Command("./binary")
	err = asymptote_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: glm")
exec.Command("latte", "install", "glm")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: fftw")
exec.Command("latte", "install", "fftw")
	fmt.Println("Instalando dependencia: ghostscript")
exec.Command("latte", "install", "ghostscript")
	fmt.Println("Instalando dependencia: gsl")
exec.Command("latte", "install", "gsl")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
	fmt.Println("Instalando dependencia: freeglut")
exec.Command("latte", "install", "freeglut")
	fmt.Println("Instalando dependencia: libtirpc")
exec.Command("latte", "install", "libtirpc")
	fmt.Println("Instalando dependencia: mesa")
exec.Command("latte", "install", "mesa")
}
