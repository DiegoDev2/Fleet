package main

// Skymaker - Generates fake astronomical images
// Homepage: https://www.astromatic.net/software/skymaker

import (
	"fmt"
	
	"os/exec"
)

func installSkymaker() {
	// Método 1: Descargar y extraer .tar.gz
	skymaker_tar_url := "https://www.astromatic.net/download/skymaker/skymaker-3.10.5.tar.gz"
	skymaker_cmd_tar := exec.Command("curl", "-L", skymaker_tar_url, "-o", "package.tar.gz")
	err := skymaker_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	skymaker_zip_url := "https://www.astromatic.net/download/skymaker/skymaker-3.10.5.zip"
	skymaker_cmd_zip := exec.Command("curl", "-L", skymaker_zip_url, "-o", "package.zip")
	err = skymaker_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	skymaker_bin_url := "https://www.astromatic.net/download/skymaker/skymaker-3.10.5.bin"
	skymaker_cmd_bin := exec.Command("curl", "-L", skymaker_bin_url, "-o", "binary.bin")
	err = skymaker_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	skymaker_src_url := "https://www.astromatic.net/download/skymaker/skymaker-3.10.5.src.tar.gz"
	skymaker_cmd_src := exec.Command("curl", "-L", skymaker_src_url, "-o", "source.tar.gz")
	err = skymaker_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	skymaker_cmd_direct := exec.Command("./binary")
	err = skymaker_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: fftw")
	exec.Command("latte", "install", "fftw").Run()
}
