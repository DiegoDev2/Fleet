package main

// Mpg321 - Command-line MP3 player
// Homepage: https://mpg321.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installMpg321() {
	// Método 1: Descargar y extraer .tar.gz
	mpg321_tar_url := "https://downloads.sourceforge.net/project/mpg321/mpg321/0.3.2/mpg321_0.3.2.orig.tar.gz"
	mpg321_cmd_tar := exec.Command("curl", "-L", mpg321_tar_url, "-o", "package.tar.gz")
	err := mpg321_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mpg321_zip_url := "https://downloads.sourceforge.net/project/mpg321/mpg321/0.3.2/mpg321_0.3.2.orig.zip"
	mpg321_cmd_zip := exec.Command("curl", "-L", mpg321_zip_url, "-o", "package.zip")
	err = mpg321_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mpg321_bin_url := "https://downloads.sourceforge.net/project/mpg321/mpg321/0.3.2/mpg321_0.3.2.orig.bin"
	mpg321_cmd_bin := exec.Command("curl", "-L", mpg321_bin_url, "-o", "binary.bin")
	err = mpg321_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mpg321_src_url := "https://downloads.sourceforge.net/project/mpg321/mpg321/0.3.2/mpg321_0.3.2.orig.src.tar.gz"
	mpg321_cmd_src := exec.Command("curl", "-L", mpg321_src_url, "-o", "source.tar.gz")
	err = mpg321_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mpg321_cmd_direct := exec.Command("./binary")
	err = mpg321_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libao")
	exec.Command("latte", "install", "libao").Run()
	fmt.Println("Instalando dependencia: libid3tag")
	exec.Command("latte", "install", "libid3tag").Run()
	fmt.Println("Instalando dependencia: mad")
	exec.Command("latte", "install", "mad").Run()
}
