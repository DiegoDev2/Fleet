package main

// Sng - Enable lossless editing of PNGs via a textual representation
// Homepage: https://sng.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installSng() {
	// Método 1: Descargar y extraer .tar.gz
	sng_tar_url := "https://downloads.sourceforge.net/project/sng/sng-1.1.1.tar.xz"
	sng_cmd_tar := exec.Command("curl", "-L", sng_tar_url, "-o", "package.tar.gz")
	err := sng_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sng_zip_url := "https://downloads.sourceforge.net/project/sng/sng-1.1.1.tar.xz"
	sng_cmd_zip := exec.Command("curl", "-L", sng_zip_url, "-o", "package.zip")
	err = sng_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sng_bin_url := "https://downloads.sourceforge.net/project/sng/sng-1.1.1.tar.xz"
	sng_cmd_bin := exec.Command("curl", "-L", sng_bin_url, "-o", "binary.bin")
	err = sng_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sng_src_url := "https://downloads.sourceforge.net/project/sng/sng-1.1.1.tar.xz"
	sng_cmd_src := exec.Command("curl", "-L", sng_src_url, "-o", "source.tar.gz")
	err = sng_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sng_cmd_direct := exec.Command("./binary")
	err = sng_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: xorgrgb")
	exec.Command("latte", "install", "xorgrgb").Run()
}
