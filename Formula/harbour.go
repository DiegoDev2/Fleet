package main

// Harbour - Portable, xBase-compatible programming language and environment
// Homepage: https://harbour.github.io

import (
	"fmt"
	
	"os/exec"
)

func installHarbour() {
	// Método 1: Descargar y extraer .tar.gz
	harbour_tar_url := "https://downloads.sourceforge.net/project/harbour-project/source/3.0.0/harbour-3.0.0.tar.bz2"
	harbour_cmd_tar := exec.Command("curl", "-L", harbour_tar_url, "-o", "package.tar.gz")
	err := harbour_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	harbour_zip_url := "https://downloads.sourceforge.net/project/harbour-project/source/3.0.0/harbour-3.0.0.tar.bz2"
	harbour_cmd_zip := exec.Command("curl", "-L", harbour_zip_url, "-o", "package.zip")
	err = harbour_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	harbour_bin_url := "https://downloads.sourceforge.net/project/harbour-project/source/3.0.0/harbour-3.0.0.tar.bz2"
	harbour_cmd_bin := exec.Command("curl", "-L", harbour_bin_url, "-o", "binary.bin")
	err = harbour_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	harbour_src_url := "https://downloads.sourceforge.net/project/harbour-project/source/3.0.0/harbour-3.0.0.tar.bz2"
	harbour_cmd_src := exec.Command("curl", "-L", harbour_src_url, "-o", "source.tar.gz")
	err = harbour_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	harbour_cmd_direct := exec.Command("./binary")
	err = harbour_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: libharu")
	exec.Command("latte", "install", "libharu").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: libxdiff")
	exec.Command("latte", "install", "libxdiff").Run()
	fmt.Println("Instalando dependencia: minizip")
	exec.Command("latte", "install", "minizip").Run()
	fmt.Println("Instalando dependencia: pcre")
	exec.Command("latte", "install", "pcre").Run()
}
