package main

// Ogmtools - OGG media streams manipulation tools
// Homepage: https://www.bunkus.org/videotools/ogmtools/

import (
	"fmt"
	
	"os/exec"
)

func installOgmtools() {
	// Método 1: Descargar y extraer .tar.gz
	ogmtools_tar_url := "https://www.bunkus.org/videotools/ogmtools/ogmtools-1.5.tar.bz2"
	ogmtools_cmd_tar := exec.Command("curl", "-L", ogmtools_tar_url, "-o", "package.tar.gz")
	err := ogmtools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ogmtools_zip_url := "https://www.bunkus.org/videotools/ogmtools/ogmtools-1.5.tar.bz2"
	ogmtools_cmd_zip := exec.Command("curl", "-L", ogmtools_zip_url, "-o", "package.zip")
	err = ogmtools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ogmtools_bin_url := "https://www.bunkus.org/videotools/ogmtools/ogmtools-1.5.tar.bz2"
	ogmtools_cmd_bin := exec.Command("curl", "-L", ogmtools_bin_url, "-o", "binary.bin")
	err = ogmtools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ogmtools_src_url := "https://www.bunkus.org/videotools/ogmtools/ogmtools-1.5.tar.bz2"
	ogmtools_cmd_src := exec.Command("curl", "-L", ogmtools_src_url, "-o", "source.tar.gz")
	err = ogmtools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ogmtools_cmd_direct := exec.Command("./binary")
	err = ogmtools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libogg")
	exec.Command("latte", "install", "libogg").Run()
	fmt.Println("Instalando dependencia: libvorbis")
	exec.Command("latte", "install", "libvorbis").Run()
}
