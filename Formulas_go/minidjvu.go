package main

// Minidjvu - DjVu multipage encoder, single page encoder/decoder
// Homepage: https://minidjvu.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installMinidjvu() {
	// Método 1: Descargar y extraer .tar.gz
	minidjvu_tar_url := "https://downloads.sourceforge.net/project/minidjvu/minidjvu/0.8/minidjvu-0.8.tar.gz"
	minidjvu_cmd_tar := exec.Command("curl", "-L", minidjvu_tar_url, "-o", "package.tar.gz")
	err := minidjvu_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	minidjvu_zip_url := "https://downloads.sourceforge.net/project/minidjvu/minidjvu/0.8/minidjvu-0.8.zip"
	minidjvu_cmd_zip := exec.Command("curl", "-L", minidjvu_zip_url, "-o", "package.zip")
	err = minidjvu_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	minidjvu_bin_url := "https://downloads.sourceforge.net/project/minidjvu/minidjvu/0.8/minidjvu-0.8.bin"
	minidjvu_cmd_bin := exec.Command("curl", "-L", minidjvu_bin_url, "-o", "binary.bin")
	err = minidjvu_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	minidjvu_src_url := "https://downloads.sourceforge.net/project/minidjvu/minidjvu/0.8/minidjvu-0.8.src.tar.gz"
	minidjvu_cmd_src := exec.Command("curl", "-L", minidjvu_src_url, "-o", "source.tar.gz")
	err = minidjvu_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	minidjvu_cmd_direct := exec.Command("./binary")
	err = minidjvu_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: djvulibre")
exec.Command("latte", "install", "djvulibre")
	fmt.Println("Instalando dependencia: libtiff")
exec.Command("latte", "install", "libtiff")
	fmt.Println("Instalando dependencia: gzip")
exec.Command("latte", "install", "gzip")
}
