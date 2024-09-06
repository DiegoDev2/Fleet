package main

// Djvulibre - DjVu viewer
// Homepage: https://djvu.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installDjvulibre() {
	// Método 1: Descargar y extraer .tar.gz
	djvulibre_tar_url := "https://downloads.sourceforge.net/djvu/djvulibre-3.5.28.tar.gz"
	djvulibre_cmd_tar := exec.Command("curl", "-L", djvulibre_tar_url, "-o", "package.tar.gz")
	err := djvulibre_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	djvulibre_zip_url := "https://downloads.sourceforge.net/djvu/djvulibre-3.5.28.zip"
	djvulibre_cmd_zip := exec.Command("curl", "-L", djvulibre_zip_url, "-o", "package.zip")
	err = djvulibre_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	djvulibre_bin_url := "https://downloads.sourceforge.net/djvu/djvulibre-3.5.28.bin"
	djvulibre_cmd_bin := exec.Command("curl", "-L", djvulibre_bin_url, "-o", "binary.bin")
	err = djvulibre_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	djvulibre_src_url := "https://downloads.sourceforge.net/djvu/djvulibre-3.5.28.src.tar.gz"
	djvulibre_cmd_src := exec.Command("curl", "-L", djvulibre_src_url, "-o", "source.tar.gz")
	err = djvulibre_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	djvulibre_cmd_direct := exec.Command("./binary")
	err = djvulibre_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: libtiff")
	exec.Command("latte", "install", "libtiff").Run()
}
