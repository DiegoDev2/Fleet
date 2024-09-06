package main

// Imageworsener - Utility and library for image scaling and processing
// Homepage: https://entropymine.com/imageworsener/

import (
	"fmt"
	
	"os/exec"
)

func installImageworsener() {
	// Método 1: Descargar y extraer .tar.gz
	imageworsener_tar_url := "https://entropymine.com/imageworsener/imageworsener-1.3.5.tar.gz"
	imageworsener_cmd_tar := exec.Command("curl", "-L", imageworsener_tar_url, "-o", "package.tar.gz")
	err := imageworsener_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	imageworsener_zip_url := "https://entropymine.com/imageworsener/imageworsener-1.3.5.zip"
	imageworsener_cmd_zip := exec.Command("curl", "-L", imageworsener_zip_url, "-o", "package.zip")
	err = imageworsener_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	imageworsener_bin_url := "https://entropymine.com/imageworsener/imageworsener-1.3.5.bin"
	imageworsener_cmd_bin := exec.Command("curl", "-L", imageworsener_bin_url, "-o", "binary.bin")
	err = imageworsener_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	imageworsener_src_url := "https://entropymine.com/imageworsener/imageworsener-1.3.5.src.tar.gz"
	imageworsener_cmd_src := exec.Command("curl", "-L", imageworsener_src_url, "-o", "source.tar.gz")
	err = imageworsener_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	imageworsener_cmd_direct := exec.Command("./binary")
	err = imageworsener_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
}
