package main

// Lsdvd - Read the content info of a DVD
// Homepage: https://sourceforge.net/projects/lsdvd/

import (
	"fmt"
	
	"os/exec"
)

func installLsdvd() {
	// Método 1: Descargar y extraer .tar.gz
	lsdvd_tar_url := "https://downloads.sourceforge.net/project/lsdvd/lsdvd/lsdvd-0.17.tar.gz"
	lsdvd_cmd_tar := exec.Command("curl", "-L", lsdvd_tar_url, "-o", "package.tar.gz")
	err := lsdvd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lsdvd_zip_url := "https://downloads.sourceforge.net/project/lsdvd/lsdvd/lsdvd-0.17.zip"
	lsdvd_cmd_zip := exec.Command("curl", "-L", lsdvd_zip_url, "-o", "package.zip")
	err = lsdvd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lsdvd_bin_url := "https://downloads.sourceforge.net/project/lsdvd/lsdvd/lsdvd-0.17.bin"
	lsdvd_cmd_bin := exec.Command("curl", "-L", lsdvd_bin_url, "-o", "binary.bin")
	err = lsdvd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lsdvd_src_url := "https://downloads.sourceforge.net/project/lsdvd/lsdvd/lsdvd-0.17.src.tar.gz"
	lsdvd_cmd_src := exec.Command("curl", "-L", lsdvd_src_url, "-o", "source.tar.gz")
	err = lsdvd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lsdvd_cmd_direct := exec.Command("./binary")
	err = lsdvd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libdvdcss")
exec.Command("latte", "install", "libdvdcss")
	fmt.Println("Instalando dependencia: libdvdread")
exec.Command("latte", "install", "libdvdread")
}
