package main

// Maxima - Computer algebra system
// Homepage: https://maxima.sourceforge.io/

import (
	"fmt"
	
	"os/exec"
)

func installMaxima() {
	// Método 1: Descargar y extraer .tar.gz
	maxima_tar_url := "https://downloads.sourceforge.net/project/maxima/Maxima-source/5.47.0-source/maxima-5.47.0.tar.gz"
	maxima_cmd_tar := exec.Command("curl", "-L", maxima_tar_url, "-o", "package.tar.gz")
	err := maxima_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	maxima_zip_url := "https://downloads.sourceforge.net/project/maxima/Maxima-source/5.47.0-source/maxima-5.47.0.zip"
	maxima_cmd_zip := exec.Command("curl", "-L", maxima_zip_url, "-o", "package.zip")
	err = maxima_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	maxima_bin_url := "https://downloads.sourceforge.net/project/maxima/Maxima-source/5.47.0-source/maxima-5.47.0.bin"
	maxima_cmd_bin := exec.Command("curl", "-L", maxima_bin_url, "-o", "binary.bin")
	err = maxima_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	maxima_src_url := "https://downloads.sourceforge.net/project/maxima/Maxima-source/5.47.0-source/maxima-5.47.0.src.tar.gz"
	maxima_cmd_src := exec.Command("curl", "-L", maxima_src_url, "-o", "source.tar.gz")
	err = maxima_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	maxima_cmd_direct := exec.Command("./binary")
	err = maxima_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gawk")
exec.Command("latte", "install", "gawk")
	fmt.Println("Instalando dependencia: gnu-sed")
exec.Command("latte", "install", "gnu-sed")
	fmt.Println("Instalando dependencia: perl")
exec.Command("latte", "install", "perl")
	fmt.Println("Instalando dependencia: texinfo")
exec.Command("latte", "install", "texinfo")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: gnuplot")
exec.Command("latte", "install", "gnuplot")
	fmt.Println("Instalando dependencia: rlwrap")
exec.Command("latte", "install", "rlwrap")
	fmt.Println("Instalando dependencia: sbcl")
exec.Command("latte", "install", "sbcl")
}
