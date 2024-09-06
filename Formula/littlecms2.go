package main

// LittleCms2 - Color management engine supporting ICC profiles
// Homepage: https://www.littlecms.com/

import (
	"fmt"
	
	"os/exec"
)

func installLittleCms2() {
	// Método 1: Descargar y extraer .tar.gz
	littlecms2_tar_url := "https://downloads.sourceforge.net/project/lcms/lcms/2.16/lcms2-2.16.tar.gz"
	littlecms2_cmd_tar := exec.Command("curl", "-L", littlecms2_tar_url, "-o", "package.tar.gz")
	err := littlecms2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	littlecms2_zip_url := "https://downloads.sourceforge.net/project/lcms/lcms/2.16/lcms2-2.16.zip"
	littlecms2_cmd_zip := exec.Command("curl", "-L", littlecms2_zip_url, "-o", "package.zip")
	err = littlecms2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	littlecms2_bin_url := "https://downloads.sourceforge.net/project/lcms/lcms/2.16/lcms2-2.16.bin"
	littlecms2_cmd_bin := exec.Command("curl", "-L", littlecms2_bin_url, "-o", "binary.bin")
	err = littlecms2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	littlecms2_src_url := "https://downloads.sourceforge.net/project/lcms/lcms/2.16/lcms2-2.16.src.tar.gz"
	littlecms2_cmd_src := exec.Command("curl", "-L", littlecms2_src_url, "-o", "source.tar.gz")
	err = littlecms2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	littlecms2_cmd_direct := exec.Command("./binary")
	err = littlecms2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: libtiff")
	exec.Command("latte", "install", "libtiff").Run()
}
