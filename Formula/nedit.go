package main

// Nedit - Fast, compact Motif/X11 plain text editor
// Homepage: https://sourceforge.net/projects/nedit/

import (
	"fmt"
	
	"os/exec"
)

func installNedit() {
	// Método 1: Descargar y extraer .tar.gz
	nedit_tar_url := "https://downloads.sourceforge.net/project/nedit/nedit-source/nedit-5.7-src.tar.gz"
	nedit_cmd_tar := exec.Command("curl", "-L", nedit_tar_url, "-o", "package.tar.gz")
	err := nedit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nedit_zip_url := "https://downloads.sourceforge.net/project/nedit/nedit-source/nedit-5.7-src.zip"
	nedit_cmd_zip := exec.Command("curl", "-L", nedit_zip_url, "-o", "package.zip")
	err = nedit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nedit_bin_url := "https://downloads.sourceforge.net/project/nedit/nedit-source/nedit-5.7-src.bin"
	nedit_cmd_bin := exec.Command("curl", "-L", nedit_bin_url, "-o", "binary.bin")
	err = nedit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nedit_src_url := "https://downloads.sourceforge.net/project/nedit/nedit-source/nedit-5.7-src.src.tar.gz"
	nedit_cmd_src := exec.Command("curl", "-L", nedit_src_url, "-o", "source.tar.gz")
	err = nedit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nedit_cmd_direct := exec.Command("./binary")
	err = nedit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libice")
	exec.Command("latte", "install", "libice").Run()
	fmt.Println("Instalando dependencia: libsm")
	exec.Command("latte", "install", "libsm").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: libxext")
	exec.Command("latte", "install", "libxext").Run()
	fmt.Println("Instalando dependencia: libxp")
	exec.Command("latte", "install", "libxp").Run()
	fmt.Println("Instalando dependencia: libxpm")
	exec.Command("latte", "install", "libxpm").Run()
	fmt.Println("Instalando dependencia: libxt")
	exec.Command("latte", "install", "libxt").Run()
	fmt.Println("Instalando dependencia: openmotif")
	exec.Command("latte", "install", "openmotif").Run()
}
